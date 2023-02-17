package service

import (
	"backend/handler"
	"backend/methods"
	"backend/mysql"
	panichandler "backend/panicHandler"
	"backend/redis"
	"backend/response"
	"backend/mysql/table"
	"encoding/json"
	"fmt"

	"net/http"
)



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
	user, company, sts := Singleton().TokenPrase(token[0])
	if !sts || perr != nil {
		return
	}

	// 檢查公司部門
	companyBanch := (*mysql.Singleton()).SelectCompanyBanch(1, company.CompanyId)
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
	v := response.Member{
		UserName: user.UserName,
		UserId: user.UserId,
		Permission: user.Permession,
		Pic: "",
		Color: fmt.Sprintf("rgba(%d,%d,%d, 0.3)", handler.Rand(0, 255), handler.Rand(0, 255), handler.Rand(0, 255)),
		Online: 1,
		Position: -1,
	}
	Singleton().Conn <- struct {
			BanchId int64
			User table.UserTable
			Value response.Member
			Company table.CompanyTable
		}{
			BanchId: conBanchId,
			User: user,
			Value: v,
			Company: company,
	}
	// 加入全域連線
	Singleton().ConnLine[user.UserId] = conn
	defer delete(Singleton().ConnLine, user.UserId)

    // The event loop
    for {
		Log.Println("socket 傳送")
		Log.Println("使用者id", user.UserId)
		Log.Print("使用者姓名", user.UserName)
		Log.Println("使用者公司碼", user.CompanyCode)
		Log.Println("使用者公司id", company.CompanyId)
		// 重設 token 過期時間
		(*redis.Singleton()).ResetExpireTime(token[0])

		// 接收訊息
        _, msg, err := conn.ReadMessage()
        if err != nil {
            Log.Println("Error during message reading:", err)
            break
        } else if user.Banch != conBanchId && user.Permession != 100 {
			Log.Print("該權限不是管理員 因此無法編輯 此部門")
			continue
		}

		data := struct {
			Types string // postion shift
			Data struct {
				response.Shift
				MyPosition int
			}
		}{}
		if json.Unmarshal(msg, &data) != nil {
			Log.Println(json.Unmarshal(msg, &data))
			continue
		}

		// type position => 位置
		// type shift => 班表的資料
		// type done => 完成編輯
		switch data.Types {
		case "position":
			// 我的位置
			v.Position = data.Data.MyPosition
			(*redis.Singleton()).EnterShiftRoom(conBanchId, v)
			Log.Println("收到 position =>",  v.Position)
			break
		case "shift":
			// 插入 班表資料
			shift := response.Shift {
				UserId: data.Data.UserId,
				Date: data.Data.Date,
				BanchStyleId: data.Data.BanchStyleId,
				RestTime: data.Data.RestTime,
				OnShiftTime: data.Data.OnShiftTime,
				OffShiftTime: data.Data.OffShiftTime,
			}
			(*redis.Singleton()).InsertShiftData(conBanchId, shift)
			Log.Println("收到 shift =>", shift)
			break
		case "done":
			
			// (*redis.Singleton()).GetShiftData(conBanchId, )
		default:
			continue
		}

		// send
		Singleton().send(conBanchId, user, company, map[string]any{})

    }

	// 離開房間
	Log.Printf("\n使用者 %d 離開房間 %d\n", user.UserId, conBanchId)
	(*redis.Singleton()).LeaveShiftRoom(conBanchId, user.UserId)

	// send
	Singleton().send(conBanchId, user, company, map[string]any{
		"newLeaving": true,
	})
}