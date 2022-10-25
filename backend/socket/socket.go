package socket

import (
	"backend/methods"
	panichandler "backend/panicHandler"
	"backend/redis"
	"backend/response"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "time"

	"github.com/gorilla/websocket"
)
 
var upgrader = websocket.Upgrader{} // use default options
 
func socketHandler(w http.ResponseWriter, r *http.Request) {
	defer panichandler.Recover()
    // Upgrade our raw HTTP connection to a websocket based one
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("Error during connection upgradation:", err)
        return
    }
    defer conn.Close()

	// header
	token := r.Header.Get("token")
	banchId := r.Header.Get("banchId")
	conBanchId , perr := methods.AnyToInt64(banchId)

	// 解析token
	user, sts := Singleton().TokenPrase(token)
	if !sts || perr != nil {
		return
	}

	// 添加成員
	v := response.Member{
		UserName: user.UserName,
		UserId: user.UserId,
		Pic: "",
		Color: "#rgb(222,222,222)",
		Online: 1,
		Position: -1,
	}
	Singleton().Conn <- struct {
			BanchId int64
			Value response.Member
		}{
			BanchId: conBanchId,
			Value: v,
	}

	// 加入全域連線
	Singleton().ConnLine[user.UserId] = conn

    // The event loop
    for {
		// 重設 token 過期時間
		(*redis.Singleton()).ResetExpireTime(token)

		// 接收訊息
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error during message reading:", err)
            break
        }

		// 插入 資料
		shift := response.Shift{}
		json.Unmarshal(msg, &shift)
		(*redis.Singleton()).InsertShiftData(conBanchId, shift)


		// 發送訊息
		users := (*redis.Singleton()).GetShiftRoomUser(conBanchId)
		data := (*redis.Singleton()).GetShiftData(conBanchId)
		(*Singleton()).SendMsg <- Msg{
			BanchId: conBanchId,
			User: *users,
			Data: *data,
		}
    }

	// 離開房間
	fmt.Printf("\n使用者 %d 離開房間 %d\n", user.UserId, conBanchId)
	(*redis.Singleton()).LeaveShiftRoom(conBanchId, user.UserId)
}
 
func Conn() {
	// rabbitMQ.Conn()
    http.HandleFunc("/workAppSocket/so", socketHandler)
    log.Fatal(http.ListenAndServe("localhost:4001", nil))
}