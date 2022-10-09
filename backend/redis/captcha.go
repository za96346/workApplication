package redis

import (
	panichandler "backend/panicHandler"
	"fmt"
	"strconv"
	"time"
)

func(dbObj *DB) InsertEmailCaptcha(email string, captcha int) {
	defer panichandler.Recover()
	(*dbObj).RedisOfCaptcha.Del(email)
	err := (*dbObj).RedisOfCaptcha.SetNX(email, captcha, time.Minute * 3)
	fmt.Println(err)
}

func(dbObj *DB) SelectEmailCaptcha(email string) int {
	defer panichandler.Recover()
	exists, _ := (*dbObj).RedisOfCaptcha.Get(email).Result() // 判斷集合裡面是否有職
	if len(exists) > 0 {
		fmt.Println("email captcha存活", exists)
		v, _:= strconv.Atoi(exists)
		return v
	} else {
		fmt.Println("email captcha過期")
		return -1
	}
}