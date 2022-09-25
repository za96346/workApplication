package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	// "strconv"
	"context"
	"time"

	"backend/table"

	"github.com/go-redis/redis"
	// "github.com/goinggo/mapstructure"
	"github.com/joho/godotenv"
)
type DB struct {
	RedisDb *redis.Client // 要先使用連線方法後才能使用這個
	fake *[]interface{}
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
}
var redisInstance *DB

var dbSingletonMux = new(sync.Mutex)

func Singleton() *DB {
	if redisInstance == nil {
		dbSingletonMux.Lock()
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

func(dbObj *DB) Conn() { // 實體化redis.Client 並返回實體的位址
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

func(dbObj *DB) IsAlive() bool {
	_, err := (*dbObj).RedisDb.Ping().Result()
	if err != nil {
		return false
	}
	return true
}

//  ------------------------------select------------------------------
		//  select all  //

//0 => all
func(dbObj *DB) SelectUser(selectKey int, value... interface{}) *[]table.UserTable {
	switch selectKey {
	case 0:
		return forEach[table.UserTable]((*dbObj).fake, (*dbObj).table[0], 2)
	default:
		return &[]table.UserTable{}
	}
}

//0 => all
func(dbObj *DB) SelectUserPreference(selectKey int, value... interface{}) *[]table.UserPreferenceTable {
	switch selectKey {
	case 0:
		return forEach[table.UserPreferenceTable]((*dbObj).fake, (*dbObj).table[1], 2)
	default:
		return &[]table.UserPreferenceTable{}
	}
}

//0 => all
func(dbObj *DB) SelectCompany(selectKey int, value... interface{}) *[]table.CompanyTable {
	switch selectKey {
	case 0:
		return forEach[table.CompanyTable]((*dbObj).fake, (*dbObj).table[2], 2)
	default:
		return &[]table.CompanyTable{}
	}
}

//0 => all
func(dbObj *DB) SelectCompanyBanch(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	switch selectKey {
	case 0:
		return forEach[table.CompanyBanchTable]((*dbObj).fake, (*dbObj).table[3], 2)
	default:
		return &[]table.CompanyBanchTable{}
	}
}

//0 => all
func(dbObj *DB) SelectShift(selectKey int, value... interface{}) *[]table.ShiftTable {
	switch selectKey {
	case 0:
		return forEach[table.ShiftTable]((*dbObj).fake, (*dbObj).table[4], 2)
	default:
		return &[]table.ShiftTable{}
	}
}

//0 => all
func(dbObj *DB) SelectShiftChange(selectKey int, value... interface{}) *[]table.ShiftChangeTable {
	switch selectKey {
	case 0:
		return forEach[table.ShiftChangeTable]((*dbObj).fake, (*dbObj).table[5], 2)
	default:
		return &[]table.ShiftChangeTable{}
	}
}

//0 => all
func(dbObj *DB) SelectShiftOverTime(selectKey int, value... interface{}) *[]table.ShiftOverTimeTable {
	switch selectKey {
	case 0:
		return forEach[table.ShiftOverTimeTable]((*dbObj).fake, (*dbObj).table[6], 2)
	default:
		return &[]table.ShiftOverTimeTable{}
	}
}

//0 => all
func(dbObj *DB) SelectDayOff(selectKey int, value... interface{}) *[]table.DayOffTable {
	switch selectKey {
	case 0:
		return forEach[table.DayOffTable]((*dbObj).fake, (*dbObj).table[7], 2)
	default:
		return &[]table.DayOffTable{}
	}
}

//0 => all
func(dbObj *DB) SelectForgetPunch(selectKey int, value... interface{}) *[]table.ForgetPunchTable {
	switch selectKey {
	case 0:
		return forEach[table.ForgetPunchTable]((*dbObj).fake, (*dbObj).table[8], 2)
	default:
		return &[]table.ForgetPunchTable{}
	}
}

//0 => all
func(dbObj *DB) SelectLateExcused(selectKey int, value... interface{}) *[]table.LateExcusedTable {
	switch selectKey {
	case 0:
		return forEach[table.LateExcusedTable]((*dbObj).fake, (*dbObj).table[9], 2)
	default:
		return &[]table.LateExcusedTable{}
	}
}

//  ------------------------------delete------------------------------

//使用者的唯一id
func(dbObj *DB) DeleteUser(deleteKey int, userId int64) bool {
	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[0], strconv.FormatInt(userId, 10))
	return true
}

//使用者的唯一id
func(dbObj *DB) DeleteUserPreference(deleteKey int, userId int64) bool {
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[1], strconv.FormatInt(userId, 10))
	return true
}

//公司的唯一id
func(dbObj *DB) DeleteCompany(deleteKey int, companyId int64) bool {
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[2], strconv.FormatInt(companyId, 10))
	return true
}

// 公司部門的id
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id int64) bool {
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[3], strconv.FormatInt(id, 10))
	return true
}

// 班表的唯一id
func(dbObj *DB) DeleteShift(deleteKey int, shiftId int64) bool {
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[4], strconv.FormatInt(shiftId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId int64) bool {
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[5], strconv.FormatInt(caseId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId int64) bool {
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[6], strconv.FormatInt(caseId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteDayOff(deleteKey int, caseId int64) bool {
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[7], strconv.FormatInt(caseId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId int64) bool {
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[8], strconv.FormatInt(caseId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId int64) bool {
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[9], strconv.FormatInt(caseId, 10))
	return true
}



//  ------------------------------insert------------------------------
func(dbObj *DB) InsertUser(data *table.UserTable) bool {
	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()
	key := strconv.FormatInt((*data).UserId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[0], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertUserPreference(data *table.UserPreferenceTable) bool {
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	key := strconv.FormatInt((*data).UserId, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[1], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertCompany(data *table.CompanyTable) bool {
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	key := strconv.FormatInt((*data).CompanyId, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[2], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertCompanyBanch(data *table.CompanyBanchTable) bool {
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	key := strconv.FormatInt((*data).Id, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[3], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertShift(data *table.ShiftTable) bool {
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	key := strconv.FormatInt((*data).ShiftId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[4], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertShiftChange(data *table.ShiftChangeTable) bool {
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[5], key,jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertShiftOverTime(data *table.ShiftOverTimeTable) bool {
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[6], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertDayOff(data *table.DayOffTable) bool {
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[7], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertForgetPunch(data *table.ForgetPunchTable) bool {
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[8], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertLateExcused(data *table.LateExcusedTable) bool {
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[9], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}



//  ------------------------------delete key------------------------------
func(dbObj *DB) DeleteKeyUser(){
	(*dbObj).RedisDb.Del((*dbObj).table[0])
}
func(dbObj *DB) DeleteKeyUserPreference(){
	(*dbObj).RedisDb.Del((*dbObj).table[1])
}
func(dbObj *DB) DeleteKeyCompany(){
	(*dbObj).RedisDb.Del((*dbObj).table[2])
}
func(dbObj *DB) DeleteKeyCompanyBanch(){
	(*dbObj).RedisDb.Del((*dbObj).table[3])
}
func(dbObj *DB) DeleteKeyShift(){
	(*dbObj).RedisDb.Del((*dbObj).table[4])
}
func(dbObj *DB) DeleteKeyShiftChange(){
	(*dbObj).RedisDb.Del((*dbObj).table[5])
}
func(dbObj *DB) DeleteKeyShiftOverTime(){
	(*dbObj).RedisDb.Del((*dbObj).table[6])
}
func(dbObj *DB) DeleteKeyDayOff(){
	(*dbObj).RedisDb.Del((*dbObj).table[7])
}
func(dbObj *DB) DeleteKeyForgetPunch(){
	(*dbObj).RedisDb.Del((*dbObj).table[8])
}
func(dbObj *DB) DeleteKeyLateExcused(){
	(*dbObj).RedisDb.Del((*dbObj).table[9])
}


//  ------------------------------method------------------------------
func(dbObj *DB) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func forEach[T any](data *[]any, key string, types int) *[]T {
	var container []T
	list := new(T)
	switch types {
	case 2:
		jsonData, _ := (*Singleton()).RedisDb.HGetAll(key).Result()

		for _, v := range jsonData {
			json.Unmarshal([]byte(v), list)
			// mapstructure.Decode(res, list)
			container = append(container, *list)
		}
		return &container
	default:
		return &container
	}
}