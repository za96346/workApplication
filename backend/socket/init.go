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
		CompanyId int64
		Value response.Member
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
					CompanyId int64
					Value response.Member
				}),
				SendMsg: make(chan Msg),
				ConnLine: make(map[int64]*websocket.Conn),
			}
			(*instance).Worker()
		}
	}
	return instance
}

func (mg *Manager) send (banchId int64, companyId int64) {
	defer panichandler.Recover()
	// 發送訊息
	onlineUsers := (*redis.Singleton()).GetShiftRoomUser(banchId)
	EditUsers := (*mysql.Singleton()).SelectUser(4, banchId)
	ShiftData := (*redis.Singleton()).GetShiftData(banchId)
	BanchStyle := (*mysql.Singleton()).SelectBanchStyle(2, banchId)
	WeekendSetting := (*mysql.Singleton()).SelectWeekendSetting(2, companyId)
	str, end := methods.GetNextMonthSE()
	// fmt.Print("開始結束", str, end)

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
		WeekendSetting: *WeekendSetting,
		Status: methods.CheckWhichStep(), // 1 開放編輯、 2 主管審核 3 確認發布
		StartDay: str,
		EndDay: end,
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
		(*mg).send(v.BanchId, v.CompanyId)
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