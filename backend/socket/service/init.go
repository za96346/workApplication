package service

import (
	"backend/handler"
	"backend/methods"
	"backend/mysql"
	panichandler "backend/panicHandler"
	"backend/redis"
	"backend/response"
	"backend/mysql/table"
	"fmt"

	"github.com/goinggo/mapstructure"
	"github.com/gorilla/websocket"
)

var instance *Manager

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

type MsgState map[string]any

type Msg struct {
	BanchId int64
	OnlineUser []response.Member // 上線的人
	EditUser []response.User // 此部門的使用者
	ShiftData []response.Shift // 班表的資料
	BanchStyle []table.BanchStyle // 部門圖標
	WeekendSetting [] table.WeekendSetting // 假日設定
	Status int // 此編輯的進度
	StartDay string // 此編輯開始日
	EndDay string // 此編輯結束日
	State MsgState // 每個人的前端各種 狀態
	NewEntering string // 剛進入的資料
	NewLeaving string // 剛出去的資料
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
/*
	step = 1,2,3,4
	permission = 100, 1, 2
*/
// 確認編輯狀態
func (mg * Manager) CheckState (step int, permission int) *MsgState {
	// 自己的編輯狀態
	disabledTable := false
	if step == 1 {disabledTable = true} // 尚未開放編輯
	if step == 2 {disabledTable = false}
	if step == 3 && permission != 1 {disabledTable = true}

	// 是否顯示 提交按鈕
	submitAble := false
	if step == 3 && permission == 1 {submitAble = true}

	state := MsgState{
		"disabledTable": disabledTable,
		"submitAble": submitAble,
	}
	return &state
}

// state["newEntering"] boolean
// state["newLeaving"] boolean
func (mg *Manager) send (
	banchId int64,
	user table.UserTable,
	company table.CompanyTable,
	state map[string]any,
) {

	defer panichandler.Recover()
	str, end, year, month := methods.GetNextMonthSE()
	// 發送訊息
	onlineUsers := (*redis.Singleton()).GetShiftRoomUser(banchId)
	EditUsers := (*mysql.Singleton()).SelectUser(4, banchId, user.CompanyCode)
	ShiftData := (*redis.Singleton()).GetShiftData(banchId, year, month)
	BanchStyle := (*mysql.Singleton()).SelectBanchStyle(2, banchId)
	fmt.Print("開始結束", year, month)
	currentStep := methods.CheckWhichStep()

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

	// 是否是剛進房間
	newEntering := ""
	if state["newEntering"] == true {
		newEntering = user.UserName
	}

	// 是否是剛出房間
	newLeaving := ""
	if state["newLeaving"] == true {
		newLeaving = user.UserName
	}

	// 傳入 隊列
	(*mg).SendMsg <- Msg{
		BanchId: banchId,
		EditUser: editUserData,
		OnlineUser: *onlineUsers,
		ShiftData: *ShiftData,
		BanchStyle: *BanchStyle,
		Status: currentStep, // 1 開放編輯、 2 主管審核 3 確認發布
		StartDay: str,
		EndDay: end,
		NewEntering: newEntering,
		NewLeaving: newLeaving,
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
		(*mg).send(v.BanchId, v.User, v.Company, map[string]any {
			"newEntering": true,
		})
	}
}

func (mg *Manager) sendMsg () {
	defer panichandler.Recover()
	for v := range (*mg).SendMsg {
		userAll := (redis.Singleton().GetShiftRoomUser(v.BanchId))
		for _, user := range *userAll {
			if (*mg).ConnLine[user.UserId] != nil {
				v.State = *(*mg).CheckState(v.Status, user.Permission)
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