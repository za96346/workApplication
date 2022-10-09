package redis

import (
	panichandler "backend/panicHandler"
	"fmt"
	"time"
)

func(dbObj *DB) InsertToken(token string) {
	defer panichandler.Recover()
	(*dbObj).RedisOfToken.SetNX(token, token,time.Minute * 30)
}

func(dbObj *DB) ClearToken(token string) {
	defer panichandler.Recover()
	// (*dbObj).RedisOfToken.(token) // 刪除集合裡面的一個值
}
func(dbObj *DB) IsTokenExited(token string) bool {
	defer panichandler.Recover()
	exists, _ := (*dbObj).RedisOfToken.Get(token).Result() // 判斷集合裡面是否有職
	if len(exists) > 0 {
		fmt.Println("token存活")
		return true
	} else {
		fmt.Println("token過期")
		return false
	}
}
