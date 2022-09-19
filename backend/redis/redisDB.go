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
)
type db struct {
	RedisDb *redis.Client

}
type dbInterface interface{
	Conn()
	InsertUserAll() bool
	InsertUserPreferenceAll() bool
	checkErr()
}
var redisInstance *db

var dbSingletonMux = new(sync.Mutex)

func RedisSingleton() *db {
	if redisInstance == nil {
		dbSingletonMux.Lock()
		if redisInstance == nil {
			redisInstance = &db{}
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
	})
	_, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
}
func(dbObj *db) InsertUserAll(data *[]database.UserTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("user", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
	
	// resData, _ = (*dbObj).RedisDb.Get("user").Result()
}
func(dbObj *db) InsertUserPreferenceAll(data *[]database.UserPreferenceTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("userPreference", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) InsertCompanyAll(data *[]database.CompanyTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("company", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) InsertCompanyBanchAll(data *[]database.CompanyBanchTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("companyBanch", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) InsertShiftAll(data *[]database.ShiftTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("shift", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) InsertShiftChangeAll(data *[]database.ShiftChangeTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("shiftChange", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) InsertShiftOverTimeAll(data *[]database.ShiftOverTimeTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("shiftOverTime", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) InsertDayOffAll(data *[]database.DayOffTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("dayOff", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) InsertForgetPunch(data *[]database.ForgetPunchTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("forgetPunch", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) InsertLateExcusedAll(data *[]database.LateExcusedTable) bool {
	for _, v := range (*data) {
		jsonData, _ := json.Marshal(v)
		_, err := (*dbObj).RedisDb.LPush("lateExcused", jsonData).Result()
		(*dbObj).checkErr(err)
	}
	return true
}
func(dbObj *db) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}