package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	// "strconv"
	"context"
	"time"

	panichandler "backend/panicHandler"

	"github.com/go-redis/redis"
	// "github.com/goinggo/mapstructure"
	"github.com/joho/godotenv"
)
type DB struct {
	RedisDb *redis.Client // 要先使用連線方法後才能使用這個
	RedisOfToken *redis.Client // token
	RedisOfCaptcha *redis.Client // captcha
	RedisOfShiftSocket *redis.Client // shift socket
	RedisOfShiftData *redis.Client // shift data
	table map[int]string
	companyMux *sync.Mutex
	userPreferenceMux *sync.RWMutex
	userMux *sync.RWMutex
	companyBanchMux *sync.RWMutex
	shiftMux *sync.RWMutex
	shiftChangeMux *sync.RWMutex
	shiftOverTimeMux *sync.RWMutex
	dayOffMux *sync.RWMutex
	lateExcusedMux *sync.RWMutex
	forgetPunchMux *sync.RWMutex
	banchStyleMux *sync.RWMutex
	banchRuleMux *sync.RWMutex
	quitWorkUserMux *sync.RWMutex
	waitCompanyReply *sync.RWMutex
	weekendSetting *sync.RWMutex
}
var redisInstance *DB

var dbSingletonMux = new(sync.Mutex)

func Singleton() *DB {
	defer panichandler.Recover()
	if redisInstance == nil {
		dbSingletonMux.Lock()
		defer dbSingletonMux.Unlock()
		if redisInstance == nil {
			redisInstance = &DB{
				table: make(map[int]string),
				companyMux: new(sync.Mutex),
				userPreferenceMux: new(sync.RWMutex),
				userMux: new(sync.RWMutex),
				companyBanchMux: new(sync.RWMutex),
				shiftMux: new(sync.RWMutex),
				shiftChangeMux: new(sync.RWMutex),
				shiftOverTimeMux: new(sync.RWMutex),
				dayOffMux: new(sync.RWMutex),
				lateExcusedMux: new(sync.RWMutex),
				forgetPunchMux: new(sync.RWMutex),
				banchStyleMux: new(sync.RWMutex),
				banchRuleMux: new(sync.RWMutex),
				quitWorkUserMux: new(sync.RWMutex),
				waitCompanyReply: new(sync.RWMutex),
				weekendSetting: new(sync.RWMutex),
			}
			(*redisInstance).table[0] = "user"
			(*redisInstance).table[1] = "userPreference"
			(*redisInstance).table[2] = "company"
			(*redisInstance).table[3] = "companyBanch"
			(*redisInstance).table[4] = "shift"
			(*redisInstance).table[5] = "shiftChange"
			(*redisInstance).table[6] = "shiftOverTime"
			(*redisInstance).table[7] = "dayOff"
			(*redisInstance).table[8] = "forgetPunch"
			(*redisInstance).table[9] = "lateExcused"
			(*redisInstance).table[10] = "banchStyle"
			(*redisInstance).table[11] = "banchRule"
			(*redisInstance).table[12] = "quitWorkUser"
			(*redisInstance).table[13] = "waitCompanyReply"
			(*redisInstance).table[14] = "weekendSetting"
			(*redisInstance).Conn()
		}
	}
	return redisInstance
}

func(dbObj *DB) Conn() { // 實體化redis.Client 並返回實體的位址
	defer panichandler.Recover()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	redisIp := os.Getenv("REDIS_DB_IP")
	redisPort := os.Getenv("REDIS_DB_PORT")
	redisPassword := os.Getenv("REDIS_DB_PASSWORD")
	// redisPool := os.Getenv("REDIS_DB_POOL")
	

	(*dbObj).RedisDb = redis.NewClient(&redis.Options{
		Addr: redisIp + ":" + redisPort,
		Password: redisPassword, // no password set
		DB: 0,  // use default DB
		PoolSize:    64,
        MinIdleConns: 16,
	})
	(*dbObj).RedisOfToken = redis.NewClient(&redis.Options{
		Addr: redisIp + ":" + redisPort,
		Password: redisPassword, // no password set
		DB: 1,  // use default DB
		PoolSize:    64,
        MinIdleConns: 16,
	})
	(*dbObj).RedisOfCaptcha = redis.NewClient(&redis.Options{
		Addr: redisIp + ":" + redisPort,
		Password: redisPassword, // no password set
		DB: 2,  // use default DB
		PoolSize:    64,
        MinIdleConns: 16,
	})
	(*dbObj).RedisOfShiftSocket = redis.NewClient(&redis.Options{
		Addr: redisIp + ":" + redisPort,
		Password: redisPassword, // no password set
		DB: 3,  // use default DB
		PoolSize:    64,
        MinIdleConns: 16,
	})
	(*dbObj).RedisOfShiftData = redis.NewClient(&redis.Options{
		Addr: redisIp + ":" + redisPort,
		Password: redisPassword, // no password set
		DB: 4,  // use default DB
		PoolSize:    64,
        MinIdleConns: 16,
	})
	_, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
}

func(dbObj *DB) IsAlive() bool {
	_, err := (*dbObj).RedisDb.Ping().Result()
	if err != nil {
		return false
	}
	return true
}

func(dbObj *DB) checkErr(err error) {
	defer panichandler.Recover()
	if err != nil {
		fmt.Println(err)
	}
}

func forEach[T any](callback func() ([]string, error), filterCallBack func(T) bool) *[]T {
	defer panichandler.Recover()
	var container []T
	list := new(T)
	jsonData, _ := callback()

	for _, v := range jsonData {
		json.Unmarshal([]byte(v), list)
		if filterCallBack(*list) {
			container = append(container, *list)
		}
	}
	return &container

}

// redis hmget tableKey => string, field..多選
func(dbObj *DB) hmGet(tableKey string, field... interface{}) ([]string, error) {
	transArr := []string{}
	for _, v := range field {
		transArr = append(transArr, fmt.Sprintf("%v", v))// format to string
	}
	returnValue := []string{}
	res, err := (*dbObj).RedisDb.HMGet(tableKey, transArr...).Result()
	for _, v := range res {
		switch x := v.(type) {
		case string:
			returnValue = append(returnValue, x)
			break
		case nil:
			break;
		}
	}
	return returnValue, err
}