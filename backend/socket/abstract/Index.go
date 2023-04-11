package abstract

import (
	panichandler "backend/panicHandler"
	"fmt"

	"github.com/gorilla/websocket"
)

type RealConnT[T any] struct {
	Key string
	Conn *websocket.Conn
	OtherProps T
}

type RealMsgT[T any] struct {
	RoomKey string
	OtherProps T
}

type Socket[ConnT any, MsgT any] struct {
	connQuene chan RealConnT[ConnT]  // 連線隊列
	msgQuene (chan RealMsgT[MsgT]) // 發送訊息對列 訊息
	connInstanceCollection map[string]*websocket.Conn // 儲存 每個連線資料
	EnterRoomCallBack func(RealConnT[ConnT]) // 進入 room 資訊的紀錄 回乎
	SendMessageCallBack func(RealMsgT[MsgT], func(ConnId string)) // 發送消息 callBack
}

// 建構事
func(sk *Socket[ConnT, MsgT]) Contruct(workMount int) {
	defer panichandler.Recover()
	(*sk).msgQuene = make(chan RealMsgT[MsgT])
	(*sk).connInstanceCollection = make(map[string]*websocket.Conn)
	(*sk).connQuene = make(chan RealConnT[ConnT])
	(*sk).worker(workMount)
}

// buffer pool
// 此工作者 是關乎到可同時讀取 connQuene and MsgQuene
func(sk *Socket[ConnT, MsgT]) worker(workMount int) {
	defer panichandler.Recover()
	for i := 0; i < workMount; i++ {
		go (*sk).listenConnQuene()
		go (*sk).listenMsgQuene()
	}
}

// 監聽 連線 隊列
func (sk *Socket[ConnT, MsgT]) listenConnQuene () {
	defer panichandler.Recover()
	for v := range (*sk).connQuene {
		(*sk).EnterRoomCallBack(v) // 執行room callback
	}
}

// 監聽 發送訊息 隊列
func (sk *Socket[ConnT, MsgT]) listenMsgQuene () {
	defer panichandler.Recover()
	for v := range (*sk).msgQuene {
		fmt.Println("發送的訊息", v)
		(*sk).SendMessageCallBack(v, func (ConnId string)  {
			if (*sk).connInstanceCollection[ConnId] != nil {
				go (*sk).connInstanceCollection[ConnId].WriteJSON(v)
			}
				
		})
	}
}

// 進入房間
func (sk *Socket[ConnT, MsgT]) EnterRoom (connId string, conn *websocket.Conn, otherProps ConnT) {
	// 放入隊列
	(*sk).connQuene <- RealConnT[ConnT]{
		Key: connId,
		Conn: conn,
		OtherProps: otherProps,
	}
	(*sk).connInstanceCollection[connId] = conn // 放入 連線 儲存池
}

// 離開房間
func (sk *Socket[ConnT, MsgT]) LeaveRoom (connId string) {
	// 
	delete(sk.connInstanceCollection, connId)
}

// 發送消息
func (sk *Socket[ConnT, MsgT]) SendMessage (message RealMsgT[MsgT]) {
	(*sk).msgQuene <- message
}