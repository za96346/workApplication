package service

import (
	// "backend/handler"
	// "backend/methods"
	// "backend/mysql"
	// "backend/mysql/table"
	"backend/mysql/table"
	panichandler "backend/panicHandler"
	"backend/redis"
	"fmt"
	"strconv"

	// "backend/redis"
	"backend/response"
	// "backend/socket/method"
	// "encoding/json"
	// "fmt"
	// "time"
	"backend/socket/abstract"

	"net/http"
)

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
}

// 實例化 shift socket
var shiftSocket = abstract.Socket[
	ConnType,
	MessageType,
] {
	EnterRoomCallBack: func(v abstract.RealConnT[ConnType]) {
		(*redis.Singleton()).EnterShiftRoom(v.OtherProps.BanchId, v.OtherProps.Value)
	},
	SendMessageCallBack: func(v abstract.RealMsgT[MessageType], sendMsg func(ConnId string)) {
		userAll := (redis.Singleton().GetShiftRoomUser(v.OtherProps.BanchId)) // 獲取 該聊天室 成員
		fmt.Print("users => ", len(*(userAll)))
		fmt.Print("roomId => ", v.OtherProps.BanchId)
		for _, user := range *userAll {
			sendMsg(strconv.FormatInt(user.UserId, 10))
		}
	},

}

func init()  {
	shiftSocket.Contruct(5) // 開啟工作者
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
	
	// // header
	// token := r.URL.Query()["token"]
	// banchId := r.URL.Query()["banchId"]
	// if len(token) <= 0 || len(banchId) <= 0 {
	// 	return
	// }
	// conBanchId , perr := methods.AnyToInt64(banchId[0])

	// // 解析token
	// user, company, sts := Singleton().TokenPrase(token[0])
	// if !sts || perr != nil {
	// 	return
	// }

	// // 檢查公司部門
	// companyBanch := (*mysql.Singleton()).SelectCompanyBanch(1, company.CompanyId)
	// if methods.IsNotExited(companyBanch) {
	// 	Log.Print("公司查無部門")
	// 	return
	// }
	// count := 0
	// for _, v := range *companyBanch {
	// 	if v.Id != conBanchId {
	// 		count += 1
	// 	} else {
	// 		break
	// 	}
	// }
	// if count == len(*companyBanch) {
	// 	Log.Print("該公司未有此部門")
	// 	return
	// }

	// // 添加成員
	// v := response.Member{
	// 	UserName: user.UserName,
	// 	UserId: user.UserId,
	// 	Permission: user.Permession,
	// 	Pic: "",
	// 	Color: fmt.Sprintf("rgba(%d,%d,%d, 0.3)", handler.Rand(0, 255), handler.Rand(0, 255), handler.Rand(0, 255)),
	// 	Online: 1,
	// 	Position: -1,
	// }
	// Singleton().Conn <- struct {
	// 		BanchId int64
	// 		User table.UserTable
	// 		Value response.Member
	// 		Company table.CompanyTable
	// 	}{
	// 		BanchId: conBanchId,
	// 		User: user,
	// 		Value: v,
	// 		Company: company,
	// }
	otherProps := ConnType{
		BanchId: 1,
		Company: *new(table.CompanyTable),
	}
	shiftSocket.EnterRoom(strconv.FormatInt(otherProps.User.UserId, 10), conn, otherProps)


    // The event loop
    for {
		// // 重設 token 過期時間
		// (*redis.Singleton()).ResetExpireTime(token[0])

		// // 接收訊息
        _, _, err := conn.ReadMessage()
        if err != nil {
            Log.Println("Error during message reading:", err)
            break
        }
		// else if user.Banch != conBanchId && user.Permession != 100 {
		// 	Log.Print("該權限不是管理員 因此無法編輯 此部門")
		// 	continue
		// }
		msgOtherProps := MessageType{
			StartDay: "hi",
			BanchId: 1,
		}
		msg := abstract.RealMsgT[MessageType]{
			RoomKey: "q",
			OtherProps: msgOtherProps,
		}
		shiftSocket.SendMessage(msg)

		// data := struct {
		// 	Types string // postion shift
		// 	Data struct {
		// 		response.Shift
		// 		MyPosition int
		// 	}
		// }{}
		// if json.Unmarshal(msg, &data) != nil {
		// 	Log.Println(json.Unmarshal(msg, &data))
		// 	continue
		// }

		// // type position => 位置
		// // type shift => 班表的資料
		// // type done => 完成編輯
		// switch data.Types {
		// case "position":
		// 	// 我的位置
		// 	v.Position = data.Data.MyPosition
		// 	(*redis.Singleton()).EnterShiftRoom(conBanchId, v)
		// 	// send
		// 	Singleton().send(conBanchId, user, company, map[string]any{})
		// 	break
		// case "shift":
		// 	// 插入 班表資料
		// 	shift := response.Shift {
		// 		UserId: data.Data.UserId,
		// 		Date: data.Data.Date,
		// 		Icon: data.Data.Icon,
		// 		BanchStyleId: data.Data.BanchStyleId,
		// 		RestTime: data.Data.RestTime,
		// 		OnShiftTime: data.Data.OnShiftTime,
		// 		OffShiftTime: data.Data.OffShiftTime,
		// 	}
		// 	(*redis.Singleton()).InsertShiftData(conBanchId, shift)
		// 	// send
		// 	Singleton().send(conBanchId, user, company, map[string]any{})
		// 	break
		// case "done":
		// 	_, _, year, month := method.GetNextMonthSE()
		// 	if method.CheckWhichStep() == 3 {
		// 		shiftArr := (*redis.Singleton()).GetShiftData(conBanchId, year, month)
		// 		for _, v := range *shiftArr {
		// 			// 格式 時間
		// 			onDate, _ :=  time.Parse("2006-01-02 15:04:05", v.Date + " " + v.OnShiftTime)
		// 			offDate, _ :=  time.Parse("2006-01-02 15:04:05", v.Date + " " + v.OffShiftTime)
		// 			now := time.Now()

		// 			// 建立資料
		// 			shift := table.ShiftTable {
		// 				UserId: v.UserId,
		// 				BanchStyleId: v.BanchStyleId,
		// 				BanchId: conBanchId,
		// 				Icon: v.Icon,
		// 				Year: year,
		// 				Month: month,
		// 				OnShiftTime: onDate,
		// 				OffShiftTime: offDate,
		// 				RestTime: v.RestTime,
		// 				CreateTime: now,
		// 				LastModify: now,
		// 			}
		// 			(*mysql.Singleton()).InsertShift(&shift)
		// 		}
		// 	}
		// 	// send
		// 	Singleton().send(conBanchId, user, company, map[string]any{
		// 		"finished": true,
		// 	})
		// 	break
		// default:
		// 	continue
		// }

    }
	shiftSocket.LeaveRoom("1")

	// // 離開房間
	// Log.Printf("\n使用者 %d 離開房間 %d\n", user.UserId, conBanchId)
	// (*redis.Singleton()).LeaveShiftRoom(conBanchId, user.UserId)

	// // send
	// Singleton().send(conBanchId, user, company, map[string]any{
	// 	"newLeaving": true,
	// })
}