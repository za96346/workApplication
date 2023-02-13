package socket

import (
	"backend/handler"
	"backend/logger"
	"backend/methods"
	"backend/mysql"
	panichandler "backend/panicHandler"
	"backend/redis"
	"backend/response"
	"backend/table"
	"fmt"
	"sync"

	"github.com/goinggo/mapstructure"
	"github.com/gorilla/websocket"
)

var lock = new(sync.Mutex)
var instance *Manager
var Redis = redis.Singleton()
var Log = logger.Logger()

type Manager struct {
	Conn chan struct {
		BanchId int64
		User table.UserTable // 當前連線的使用者資料
		Value response.Member
		Company table.CompanyTable // 當前連線的 公司資料
	}
	SendMsg chan Msg
	ConnLine map[int64]*websocket.Conn
}
type Msg struct {
	BanchId int64
	OnlineUser []response.Member
	EditUser []response.User
	ShiftData []response.Shift
	BanchStyle []table.BanchStyle
	WeekendSetting [] table.WeekendSetting
	Status int
	StartDay string
	EndDay string
	State map[string]any
}

func Singleton () *Manager {
	defer panichandler.Recover()
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Manager{
				Conn: make(chan struct{
					BanchId int64
					User table.UserTable
					Value response.Member
					Company table.CompanyTable

				}),
				SendMsg: make(chan Msg),
				ConnLine: make(map[int64]*websocket.Conn),
			}
			(*instance).Worker()
		}
	}
	return instance
}

func (mg *Manager) send (banchId int64, user table.UserTable, company table.CompanyTable) {
	defer panichandler.Recover()
	str, end, year, month := methods.GetNextMonthSE()
	// 發送訊息
	onlineUsers := (*redis.Singleton()).GetShiftRoomUser(banchId)
	EditUsers := (*mysql.Singleton()).SelectUser(4, banchId, user.CompanyCode)
	ShiftData := (*redis.Singleton()).GetShiftData(banchId, year, month)
	BanchStyle := (*mysql.Singleton()).SelectBanchStyle(2, banchId)
	fmt.Print("開始結束", year, month)
	currentStep := methods.CheckWhichStep()

	// 自己的編輯狀態
	disabledTable := false
	if currentStep == 2 && user.Permession == 2 {disabledTable = true}
	if currentStep == 3 {disabledTable = true}

	// 整理 回傳的編輯使用者資料
	editUserData := []response.User{}
	for _, v := range *EditUsers {
		editUserData = append(editUserData, response.User{
			UserName: v.UserName,
			UserId: v.UserId,
			CompanyCode: v.CompanyCode,
			EmployeeNumber: v.EmployeeNumber,
			OnWorkDay: v.OnWorkDay,
			Banch: v.Banch,
			Permession: v.Permession,
			WorkState: "on",
		})
	}
	(*mg).SendMsg <- Msg{
		BanchId: banchId,
		EditUser: editUserData,
		OnlineUser: *onlineUsers,
		ShiftData: *ShiftData,
		BanchStyle: *BanchStyle,
		Status: currentStep, // 1 開放編輯、 2 主管審核 3 確認發布
		StartDay: str,
		EndDay: end,
		State: map[string]any{
			"disabledTable": disabledTable,
		},
	}
}

func (mg *Manager) Worker() {
	defer panichandler.Recover()
	for i := 0; i < 5; i++ {
		go (*mg).enterRoom()
		go (*mg).sendMsg()
	}
}

func (mg *Manager) enterRoom () {
	defer panichandler.Recover()
	for v := range (*mg).Conn {
		Log.Printf("\n使用者編號 %d 進入部門房間 %d\n", v.Value.UserId, v.BanchId)
		(*redis.Singleton()).EnterShiftRoom(v.BanchId, v.Value)
		(*mg).send(v.BanchId, v.User, v.Company)
	}
}

func (mg *Manager) sendMsg () {
	defer panichandler.Recover()
	for v := range (*mg).SendMsg {
		userAll := (redis.Singleton().GetShiftRoomUser(v.BanchId))
		for _, user := range *userAll {
			if (*mg).ConnLine[user.UserId] != nil {
				go (*mg).ConnLine[user.UserId].WriteJSON(v)
			}
		}
	}
}

func (mg *Manager) TokenPrase (tokenParams string) (table.UserTable, table.CompanyTable, bool) {
	defer panichandler.Recover()
	// 判斷 token 是否過期
	if !(*Redis).IsTokenExited(tokenParams) {
		return table.UserTable{}, table.CompanyTable{}, false
	}
		
	// 解析 token
	userInfo, err := handler.ParseToken(tokenParams)
	if err != nil {
		return table.UserTable{}, table.CompanyTable{}, false
	}
	
	user := new(table.UserTable)
	company := new(table.CompanyTable)

	mapstructure.Decode(userInfo["User"], user)
	mapstructure.Decode(userInfo["Company"], company)

	(*Redis).ResetExpireTime(tokenParams)

	return *user, *company, true
}