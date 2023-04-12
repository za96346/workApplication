package abstract

import (
	panichandler "backend/panicHandler"
	"fmt"

	"github.com/gorilla/websocket"
)

type RealConnT[T any] struct {
	RoomKey string // 房間id
	Conn *websocket.Conn // 連線實例
	OtherProps T // 連線值
}

type RealMsgT[T any] struct {
	RoomKey string // 房間id
	OtherProps T // 訊息值
}

type Socket[ConnT any, MsgT any] struct {
	connQuene chan RealConnT[ConnT]  // 連線隊列
	msgQuene (chan RealMsgT[MsgT]) // 發送訊息對列 訊息
	connInstanceCollection map[string]*websocket.Conn // 儲存 每個連線資料
	EnterRoomCallBack func(ConnT, string) // 進入 room 資訊的紀錄 回乎
	SendMessageCallBack func(MsgT, string, func(string, any)) // 發送消息 callBack
}

// 建構事
func Instance[ConnT any, MsgT any](workMount int) *Socket[ConnT, MsgT] {
	defer panichandler.Recover()
	socket := new(Socket[ConnT, MsgT])

	// 出使化
	(*socket).msgQuene = make(chan RealMsgT[MsgT])
	(*socket).connInstanceCollection = make(map[string]*websocket.Conn)
	(*socket).connQuene = make(chan RealConnT[ConnT])
	(*socket).worker(workMount)
	(*socket).EnterRoomCallBack = func(rct ConnT, connKey string) {
		fmt.Println("請指派 進入房間的 call back")
	}
	(*socket).SendMessageCallBack = func(mt MsgT, s string, f func(ConnId string,  Msg any)) {
		fmt.Println("請指派 傳送訊息的 call back")
	}

	//
	return socket
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
		(*sk).EnterRoomCallBack(v.OtherProps, v.RoomKey) // 執行room callback
	}
}

// 監聽 發送訊息 隊列
func (sk *Socket[ConnT, MsgT]) listenMsgQuene () {
	defer panichandler.Recover()
	for v := range (*sk).msgQuene {
		// fmt.Println("發送的訊息", v.OtherProps)
		(*sk).SendMessageCallBack(v.OtherProps, v.RoomKey, func (ConnId string, Msg any)  {
			if (*sk).connInstanceCollection[ConnId] != nil {
				(*sk).connInstanceCollection[ConnId].WriteJSON(Msg)
			}
		})
	}
}

// 進入房間
func (sk *Socket[ConnT, MsgT]) EnterRoom (
	connId string, // 連線id
	conn *websocket.Conn, // 連線實力
	otherProps ConnT, // 連線 值
	Roomkey string, // 房間key
) {
	// 放入隊列
	(*sk).connQuene <- RealConnT[ConnT]{
		RoomKey: Roomkey,
		Conn: conn,
		OtherProps: otherProps,
	}
	(*sk).connInstanceCollection[connId] = conn // 放入 連線 儲存池
}

// 離開房間
func (sk *Socket[ConnT, MsgT]) LeaveRoom (connId string) {
	// 
	// (*sk).connInstanceCollection[connId].Close()
	delete((*sk).connInstanceCollection, connId)
}

// 發送消息
func (sk *Socket[ConnT, MsgT]) SendMessage (message MsgT, RoomKey string) {
	(*sk).msgQuene <- RealMsgT[MsgT] {
		RoomKey: RoomKey,
		OtherProps: message,
	}
}