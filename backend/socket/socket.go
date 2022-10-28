package socket

import (
	"backend/handler"
	"backend/methods"
	panichandler "backend/panicHandler"
	"backend/redis"
	"backend/response"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// "time"

	"github.com/gorilla/websocket"
)
 
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options
 
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
	token := r.URL.Query()["token"]
	banchId := r.URL.Query()["banchId"]
	if len(token) <= 0 || len(banchId) <= 0 {
		return
	}
	conBanchId , perr := methods.AnyToInt64(banchId[0])

	// 解析token
	user, sts := Singleton().TokenPrase(token[0])
	if !sts || perr != nil {
		return
	}

	// 添加成員
	v := response.Member{
		UserName: user.UserName,
		UserId: user.UserId,
		Pic: "",
		Color: fmt.Sprintf("rgb(%d,%d,%d)", handler.Rand(0, 255), handler.Rand(0, 255), handler.Rand(0, 255)),
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
		(*redis.Singleton()).ResetExpireTime(token[0])

		// 接收訊息
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error during message reading:", err)
            break
        }

		data := struct {
			Types int
			Data struct {
				response.Shift
				MyPosition int
			}
		}{}
		if json.Unmarshal(msg, &data) != nil {
			continue
		}

		switch data.Types {
		case 1:
			// 我的位置
			v.Position = data.Data.MyPosition
			(*redis.Singleton()).EnterShiftRoom(conBanchId, v)
			fmt.Println("收到 position =>",  v.Position)
			break
		case 2:
			// 插入 班表資料
			shift := response.Shift {
				UserId: data.Data.UserId,
				Position: data.Data.Position,
				OnShiftTime: data.Data.OnShiftTime,
				OffShiftTime: data.Data.OffShiftTime,
			}
			(*redis.Singleton()).InsertShiftData(conBanchId, shift)
			fmt.Println("收到 shift =>", shift)
			break
		default:
			continue
		}

		// send
		Singleton().send(conBanchId)

    }

	// 離開房間
	fmt.Printf("\n使用者 %d 離開房間 %d\n", user.UserId, conBanchId)
	(*redis.Singleton()).LeaveShiftRoom(conBanchId, user.UserId)
}
 
func Conn() {
	// rabbitMQ.Conn()
	ip := os.Getenv("SOCKET_IP")
	port := os.Getenv("SOCKET_PORT")
    http.HandleFunc("/workAppSocket/shift", socketHandler)
    log.Fatal(http.ListenAndServe(ip + ":" + port, nil))
}