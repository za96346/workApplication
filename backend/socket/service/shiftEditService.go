package service

import (
	"backend/handler"
	"backend/handler/shiftEdit"
	"backend/methods"
	"backend/mysql/table"
	panichandler "backend/panicHandler"
	"backend/response"
	"fmt"
	"strconv"
	"time"

	// "backend/socket/method"
	"encoding/json"
	// "fmt"
	// "time"
	"backend/socket/abstract"
	"backend/socket/method"

	"net/http"

	"github.com/goinggo/mapstructure"
)
type MsgState map[string]any
// 連線資訊
type ConnType struct {
	BanchId int64
	User table.UserTable // 當前連線的使用者資料
	Value response.Member
	Company table.CompanyTable // 當前連線的 公司資料
}

// 要傳送socket 的訊息
type MessageType struct {
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
	LauchPerson table.UserTable // 此訊息的發起人
	RowsShiftTotal map[int64]float64 //列的總計
	ColumnsShiftTotal map[string]float64 // 欄的總計
}

// 接收到的資料
type ReceivedMsgType struct {
	Types string // postion shift
	Data struct {
		response.Shift
		MyPosition int
	}
}

// 聊天室 的狀態
type RoomStatus struct {
	LastFinishedYearMonth string
}

// 實例化 shift socket
var shiftSocket = abstract.Instance[ConnType, MessageType](5)

func init()  {
	shiftSocket.EnterRoomCallBack = func(v ConnType, roomKey string) {
		(*Redis).EnterShiftRoom(v.BanchId, v.Value)
		sendMsgHandler(v.BanchId, v.User, v.Company, map[string]any{
			"newEntering": true,
		})
	};
	shiftSocket.SendMessageCallBack = func(
		v MessageType,
		roomKey string,
		sendMsg func(ConnId string, Msg any),
	) {
		userAll := (Redis.GetShiftRoomUser(v.BanchId)) // 獲取 該聊天室 成員
		// fmt.Print("users => ", len(*(userAll)))
		// fmt.Print("roomId => ", v.BanchId)
		for _, user := range *userAll {

			// 這是只有 自己是發起人才要傳送 錯誤訊息
			if user.UserId != v.LauchPerson.UserId {
				v.State["errorMsg"] = ""
			}

			// 根據 權限 獲取 前端 操作 狀態
			getCheckState := CheckState(v.Status, user.Permission, user.BanchId, v.BanchId)
			v.State["disabledTable"] = getCheckState["disabledTable"]
			v.State["submitAble"] = getCheckState["submitAble"]

			sendMsg(strconv.FormatInt(user.UserId, 10), v)
		}
	};
}

func ShiftSocketHandler(w http.ResponseWriter, r *http.Request) {
	defer panichandler.Recover()
    // Upgrade our raw HTTP connection to a websocket based one
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        Log.Print("Error during connection upgradation:", err)
        return
    }
    defer conn.Close()
	
	// header
	token := r.URL.Query()["token"]
	banchId := r.URL.Query()["banchId"]
	if len(token) <= 0 || len(banchId) <= 0 {
		return
	}
	conBanchId , perr := methods.AnyToInt64(banchId[0])

	// 解析token
	user, company, sts := TokenPrase(token[0])
	if !sts || perr != nil {
		return
	}

	// 檢查公司部門
	companyBanch := (*Mysql).SelectCompanyBanch(1, company.CompanyId)
	if methods.IsNotExited(companyBanch) {
		Log.Print("公司查無部門")
		return
	}
	count := 0
	for _, v := range *companyBanch {
		if v.Id != conBanchId {
			count += 1
		} else {
			break
		}
	}
	if count == len(*companyBanch) {
		Log.Print("該公司未有此部門")
		return
	}

	// 添加成員
	newRoomMember := response.Member{
		UserName: user.UserName,
		UserId: user.UserId,
		BanchId: user.Banch,
		Permission: user.Permession,
		Pic: "",
		Color: fmt.Sprintf("rgba(%d,%d,%d, 0.3)", handler.Rand(0, 255), handler.Rand(0, 255), handler.Rand(0, 255)),
		Online: 1,
		Position: -1,
	}
	connProps := ConnType{
		BanchId: conBanchId,
		User: user,
		Value: newRoomMember,
		Company: company,
	}
	shiftSocket.EnterRoom(
		strconv.FormatInt(connProps.User.UserId, 10),
		conn,
		connProps,
		"any",
	)

	logMsg := ""

    // The event loop
    for {
		// // 重設 token 過期時間
		// (*Redis).ResetExpireTime(token[0])
		_, _, year, month := method.GetNextMonthSE()
		// // 接收訊息
        _, receivedMsg, err := conn.ReadMessage()
        if err != nil {
            Log.Println("Error during message reading:", err)
            break
        } else if user.Banch != conBanchId && user.Permession != 100 {
			Log.Print("該權限不是管理員 因此無法編輯 此部門")
			continue
		}

		data := ReceivedMsgType{}
		if json.Unmarshal(receivedMsg, &data) != nil {
			Log.Println(json.Unmarshal(receivedMsg, &data))
			continue
		}
		fmt.Println("type", data.Types)
		

		// type position => 位置
		// type shift => 班表的資料
		// type done => 完成編輯
		switch data.Types {
		case "position":
			// 我的位置
			newRoomMember.Position = data.Data.MyPosition
			(*Redis).EnterShiftRoom(conBanchId, newRoomMember)
			// send
			sendMsgHandler(conBanchId, user, company, map[string]any{})
			break
		case "shift":
			// 插入 班表資料
			shift := response.Shift {
				UserId: data.Data.UserId,
				Date: data.Data.Date,
				Icon: data.Data.Icon,
				BanchStyleId: data.Data.BanchStyleId,
				RestTime: data.Data.RestTime,
				OnShiftTime: data.Data.OnShiftTime,
				OffShiftTime: data.Data.OffShiftTime,
			}

			(*Redis).InsertShiftData(conBanchId, shift)
			// 這邊要記錄log
			go func ()  {
				findUser := (*Mysql).SelectUser(1, shift.UserId)
				if !methods.IsNotExited(findUser) {
					logMsg = fmt.Sprintln(user.UserName, "新增了 ", (*findUser)[0].UserName, "  ", shift.Date, "  ", shift.Icon)
				}
			}()
			// send
			sendMsgHandler(conBanchId, user, company, map[string]any{})
			break
		case "done":
			if method.CheckWhichStep(conBanchId) == 3 {
				shiftArr := (*Redis).GetShiftData(conBanchId, year, month)
				insertResult := true
				transaction, _ := (*Mysql).MysqlDB.Begin()
				for _, v := range *shiftArr {
					// 格式 時間
					onDate, _ :=  time.Parse("2006-01-02 15:04:05", v.Date + " " + v.OnShiftTime)
					offDate, _ :=  time.Parse("2006-01-02 15:04:05", v.Date + " " + v.OffShiftTime)
					now := time.Now()

					// 建立資料
					shift := table.ShiftTable {
						UserId: v.UserId,
						BanchStyleId: v.BanchStyleId,
						BanchId: conBanchId,
						Icon: v.Icon,
						Year: year,
						Month: month,
						OnShiftTime: onDate,
						OffShiftTime: offDate,
						RestTime: v.RestTime,
						CreateTime: now,
						LastModify: now,
					}
					insertResult, _ = (*Mysql).InsertShift(&shift)
					if !insertResult {
						break
					}
				}

				// 如果新增失敗
				if !insertResult {
					fmt.Println("rollBack")
					transaction.Rollback()

					// 紀錄log
					logMsg = fmt.Sprintln(user.UserName, "  ",year, "-", month, "班表資料 提交 失敗")

					// send failed
					sendMsgHandler(conBanchId, user, company, map[string]any{
						"finished": true,
						"errorMsg": "執行失敗",
					})
					continue
				} else {
					// 紀錄log
					logMsg = fmt.Sprintln(user.UserName, "  ",year, "-", month, "班表資料 提交 成功")
					transaction.Commit()
					// 快取資料庫操作
					(*Redis).DeleteShiftData(conBanchId)
					(*Redis).InsertShiftRoomStatus(conBanchId, RoomStatus{
						LastFinishedYearMonth: fmt.Sprintln(year, "/", month),
					})
					// send finished
					sendMsgHandler(conBanchId, user, company, map[string]any{
						"finished": true,
					})
				}
			}
			break
		default:
			continue
		}
		if (data.Types != "position") {
			go (*Mysql).InsertShiftEditLog(&table.ShiftEditLog{
				BanchId: conBanchId,
				Year: year,
				Month: month,
				Msg: logMsg,
			})	
		}
    }
	shiftSocket.LeaveRoom("1")

	// // 離開房間
	(*Redis).LeaveShiftRoom(conBanchId, user.UserId)

	// // send
	sendMsgHandler(conBanchId, user, company, map[string]any{
		"newLeaving": true,
	})
}

// state["newEntering"] boolean
// state["newLeaving"] boolean
// state["finished"] boolean
func sendMsgHandler(
	banchId int64,
	user table.UserTable,
	company table.CompanyTable,
	state map[string]any,
) {
	defer panichandler.Recover()
	str, end, year, month := method.GetNextMonthSE()
	// 發送訊息
	onlineUsers := (*Redis).GetShiftRoomUser(banchId) // 線上使用者資料
	EditUsers := (*Mysql).SelectUser(4, banchId, user.CompanyCode) // 被編輯的使用者
	ShiftData := (*Redis).GetShiftData(banchId, year, month) // 當前的班表資料
	BanchStyle := (*Mysql).SelectBanchStyle(2, banchId) // 部門圖標
	currentStep := method.CheckWhichStep(banchId) // 當前的編輯狀態

	rowsShiftTotal, columnsShiftTotal := shiftEdit.ShiftTotal(ShiftData)
	// fmt.Println(*rowsShiftTotal)

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
		(*Mysql).InsertShiftEditLog(&table.ShiftEditLog{
			BanchId: banchId,
			Year: year,
			Month: month,
			Msg: user.UserName + " " + "進入房間",
		})	
	}

	// 是否是剛出房間
	newLeaving := ""
	if state["newLeaving"] == true {
		newLeaving = user.UserName
		(*Mysql).InsertShiftEditLog(&table.ShiftEditLog{
			BanchId: banchId,
			Year: year,
			Month: month,
			Msg: user.UserName + " " + "離開房間",
		})	
	}

	msg := MessageType{
		BanchId: banchId,
		EditUser: editUserData,
		OnlineUser: *onlineUsers,
		ShiftData: *ShiftData,
		BanchStyle: *BanchStyle,
		Status: currentStep, // 1 開放編輯、 2 主管審核 3 確認發布 4. 等待 編輯
		StartDay: str,
		EndDay: end,
		State: state,
		NewEntering: newEntering,
		NewLeaving: newLeaving,
		LauchPerson: user,
		RowsShiftTotal: *rowsShiftTotal,
		ColumnsShiftTotal: *columnsShiftTotal,
	}
	shiftSocket.SendMessage(msg, "any")
}


/*
	step = 1,2,3,4
	permission = 100, 1, 2
*/
// 確認編輯狀態
func CheckState (step int, permission int, selfBanchId int64, roomBanchId int64) MsgState {
	// 自己的編輯狀態
	disabledTable := false
	if step == 1 {disabledTable = true} // 尚未開放編輯
	if step == 2 {disabledTable = false}
	if step == 3 && permission != 1 {disabledTable = true} // 只有 主管可以編輯
	if step == 4 {disabledTable = true}

	// 是否顯示 提交按鈕
	submitAble := false
	if step == 3 && permission == 1 && (roomBanchId == selfBanchId) {
		submitAble = true
	}
	state := MsgState{
		"disabledTable": disabledTable,
		"submitAble": submitAble,
	}
	// fmt.Println("permission=> ", permission, state)
	return state
}

func TokenPrase (tokenParams string) (table.UserTable, table.CompanyTable, bool) {
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