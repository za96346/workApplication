package socket

import (
	"backend/handler"
	"backend/methods"
	panichandler "backend/panicHandler"
	"backend/redis"
	"backend/response"
	"backend/table"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

var lock = new(sync.Mutex)
var instance *Manager

type Manager struct {
	Conn chan struct {
		BanchId int64
		Value response.Member
	}
	SendMsg chan Msg
	ConnLine map[int64]*websocket.Conn
}
type Msg struct {
	BanchId int64
	User []response.Member
	Data []response.Shift
}

func Singleton () *Manager {
	defer panichandler.Recover()
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Manager{
				Conn: make(chan struct{
					BanchId int64;
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

func (mg *Manager) Worker() {
	for i := 0; i < 5; i++ {
		go (*mg).enterRoom()
		go (*mg).sendMsg()
	}
}

func (mg *Manager) enterRoom () {
	for v := range (*mg).Conn {
		fmt.Printf("\n使用者編號 %d 進入部門房間 %d", v.Value.UserId, v.BanchId)
		(*redis.Singleton()).EnterShiftRoom(v.BanchId, v.Value)
	}
}

func (mg *Manager) sendMsg () {
	for v := range (*mg).SendMsg {
		userAll := (redis.Singleton().GetShiftRoomUser(v.BanchId))
		for _, user := range *userAll {
			(*mg).ConnLine[user.UserId].WriteJSON(v)
		}
	}
}

func (mg *Manager) TokenPrase (tokenParams string) (table.UserTable, bool) {
	// 判斷 token 是否過期
	if !handler.Singleton().Redis.IsTokenExited(tokenParams) {
		return table.UserTable{}, false
	}
		
	// 解析 token
	userInfo, err := handler.ParseToken(tokenParams)
	if err != nil {
		return table.UserTable{}, false
	}
	
	(*handler.Singleton()).Redis.ResetExpireTime(tokenParams)

	user := (*handler.Singleton()).SelectUser(1, int64(userInfo["UserId"].(float64)))
	if methods.IsNotExited(user) {
		return table.UserTable{}, false
	}

	return (*user)[0], true
}