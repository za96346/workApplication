package service

import (
	"backend/logger"
	"backend/redis"
	"net/http"

	"sync"

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

var lock = new(sync.Mutex)
var Redis = redis.Singleton()
var Log = logger.Logger()
