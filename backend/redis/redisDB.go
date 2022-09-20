package redis

import (
	"backend/database"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	// "strconv"
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/goinggo/mapstructure"
)
type db struct {
	RedisDb *redis.Client
	fake *[]interface{}
	table map[int]string 
}
var redisInstance *db

var dbSingletonMux = new(sync.Mutex)

func RedisSingleton() *db {
	if redisInstance == nil {
		dbSingletonMux.Lock()
		if redisInstance == nil {
			redisInstance = &db{
				table: make(map[int]string),
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
			defer dbSingletonMux.Unlock()
		}
	}
	return redisInstance
}

func(dbObj *db) Conn() { // 實體化redis.Client 並返回實體的位址
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
	_, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
}



func(dbObj *db) SelectUserAll() *[]interface{} {
	res := forEach((*dbObj).fake, (*dbObj).table[0], 2)
	return maptoStruct(res)
}
func(dbObj *db) SelectUserPreferenceAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[1], 2)
}
func(dbObj *db) SelectCompanyAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[2], 2)
}
func(dbObj *db) SelectCompanyBanchAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[3], 2)
}
func(dbObj *db) SelectShiftAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[4], 2)
}
func(dbObj *db) SelectShiftChangeAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[5], 2)
}
func(dbObj *db) SelectShiftOverTimeAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[6], 2)
}
func(dbObj *db) SelectDayOffAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[7], 2)
}
func(dbObj *db) SelectForgetPunchAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[8], 2)
}
func(dbObj *db) SelectLateExcusedAll() *[]interface{} {
	return forEach((*dbObj).fake, (*dbObj).table[9], 2)
}






//  ------------------------------insert------------------------------
func(dbObj *db) InsertUserAll(data *[]database.UserTable) bool {
	forEach(data, (*dbObj).table[0], 1)
	return true
}
func(dbObj *db) InsertUserPreferenceAll (data *[]database.UserPreferenceTable) bool {
	forEach(data, (*dbObj).table[1], 1)
	return true
}
func(dbObj *db) InsertCompanyAll(data *[]database.CompanyTable) bool {
	forEach(data, (*dbObj).table[2], 1)
	return true
}
func(dbObj *db) InsertCompanyBanchAll(data *[]database.CompanyBanchTable) bool {
	forEach(data, (*dbObj).table[3], 1)
	return true
}
func(dbObj *db) InsertShiftAll(data *[]database.ShiftTable) bool {
	forEach(data, (*dbObj).table[4], 1)
	return true
}
func(dbObj *db) InsertShiftChangeAll(data *[]database.ShiftChangeTable) bool {
	forEach(data, (*dbObj).table[5], 1)
	return true
}
func(dbObj *db) InsertShiftOverTimeAll(data *[]database.ShiftOverTimeTable) bool {
	forEach(data, (*dbObj).table[6], 1)
	return true
}
func(dbObj *db) InsertDayOffAll(data *[]database.DayOffTable) bool {
	forEach(data, (*dbObj).table[7], 1)
	return true
}
func(dbObj *db) InsertForgetPunch(data *[]database.ForgetPunchTable) bool {
	forEach(data, (*dbObj).table[8], 1)
	return true
}
func(dbObj *db) InsertLateExcusedAll(data *[]database.LateExcusedTable) bool {
	forEach(data, (*dbObj).table[9], 1)
	return true
}
func(dbObj *db) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func forEach[T any](data *[]T, key string, types int) *[]interface{} {
	switch types {
	case 1:
		for _, v := range (*data) {
			jsonData, _ := json.Marshal(v)
			_, err := (*RedisSingleton()).RedisDb.LPush(key, jsonData).Result()
			(*RedisSingleton()).checkErr(err)
		}
		return (*RedisSingleton()).fake
	case 2:
		jsonData, _ := (*RedisSingleton()).RedisDb.LRange(key, 0, -1).Result()
		var container []interface{}
		var list map[string]interface{}
		for _, res := range jsonData {
			json.Unmarshal([]byte(res), &list)
			container = append(container, list)
		}
		return &container
	default:
		return (*RedisSingleton()).fake
	}
}
func maptoStruct(arr *[]interface{}) *[]interface{} {
	var decode []interface{}
	var user database.UserTable
	for _, res := range *arr {
		decode = append(decode, mapstructure.Decode(res, &user))
		fmt.Println(res)
	}
	return &decode
}