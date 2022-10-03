package redis

import (
	panichandler "backend/panicHandler"
	"fmt"
	"time"
)

func(dbObj *DB) InsertToken(token string) {
	defer panichandler.Recover()
	(*dbObj).RedisDb.SAdd("token", token)
	(*dbObj).RedisDb.Expire("token", time.Minute * 30)
}

func(dbObj *DB) ClearToken(token string) {
	defer panichandler.Recover()
	(*dbObj).RedisDb.SRem(token, token) // 刪除集合裡面的一個值
}
func(dbObj *DB) IsTokenExited(token string) bool {
	defer panichandler.Recover()
	exists, _ := (*dbObj).RedisDb.SIsMember("token", token).Result() // 判斷集合裡面是否有職
	if exists {
		fmt.Println("存在集合中")
		return true
	} else {
		fmt.Println("不存在集合中")
		return false
	}
}