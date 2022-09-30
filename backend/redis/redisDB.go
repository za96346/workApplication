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

	panichandler "backend/panicHandler"
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
	defer panichandler.Recover()
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

// 0 => 全部, value =>  nil
//  1 =>  userId, value => int64
//  2 => account, value => string
func(dbObj *DB) SelectUser(selectKey int, value... interface{}) *[]table.UserTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[0]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.UserTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.UserTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.UserTable) bool {
				for _, filterItem := range value {
					if filterItem == v.Account {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.UserTable{}
	}
}

// 0 => 全部, value => nil
//  1 => 使用者id, value => int64
func(dbObj *DB) SelectUserPreference(selectKey int, value... interface{}) *[]table.UserPreferenceTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[1]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(V table.UserPreferenceTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(V table.UserPreferenceTable) bool {return true},
		)
	default:
		return &[]table.UserPreferenceTable{}
	}
}

// 0 => 全部, value => nil
//  1 => 公司id, value => int64
//  2 => 公司碼, value => string
func(dbObj *DB) SelectCompany(selectKey int, value... interface{}) *[]table.CompanyTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[2]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.CompanyTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.CompanyTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.CompanyTable) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyCode {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.CompanyTable{}
	}
}

// 0 => 全部, value => nil
//	1 => 公司Id, value => int64
// 	2 => id (banchId), value => int64
func(dbObj *DB) SelectCompanyBanch(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[3]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.CompanyBanchTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.CompanyBanchTable) bool {
				for _, filterItem := range value {
					if int64(filterItem.(int)) == v.CompanyId {
						return true
					}
				}
				return false
			},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.CompanyBanchTable) bool {return true},
		)
	default:
		return &[]table.CompanyBanchTable{}
	}
}

// 0 => all, value => nil
//  1 => 班表id, value => int64
func(dbObj *DB) SelectShift(selectKey int, value... interface{}) *[]table.ShiftTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[4]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.ShiftTable) bool {return true},
		)
	default:
		return &[]table.ShiftTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectShiftChange(selectKey int, value... interface{}) *[]table.ShiftChangeTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[5]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftChangeTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.ShiftChangeTable) bool {return true},
		)
	default:
		return &[]table.ShiftChangeTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectShiftOverTime(selectKey int, value... interface{}) *[]table.ShiftOverTimeTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[6]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftOverTimeTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.ShiftOverTimeTable) bool {return true},
		)
	default:
		return &[]table.ShiftOverTimeTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectDayOff(selectKey int, value... interface{}) *[]table.DayOffTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[7]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.DayOffTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.DayOffTable) bool {return true},
		)
	default:
		return &[]table.DayOffTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectForgetPunch(selectKey int, value... interface{}) *[]table.ForgetPunchTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[8]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ForgetPunchTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.ForgetPunchTable) bool {return true},
		)
	default:
		return &[]table.ForgetPunchTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectLateExcused(selectKey int, value... interface{}) *[]table.LateExcusedTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[9]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.LateExcusedTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.LateExcusedTable) bool {return true},
		)
	default:
		return &[]table.LateExcusedTable{}
	}
}

//  ------------------------------delete------------------------------

//使用者的唯一id
func(dbObj *DB) DeleteUser(deleteKey int, userId int64) bool {
	defer panichandler.Recover()
	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[0], strconv.FormatInt(userId, 10))
	return true
}

//使用者的唯一id
func(dbObj *DB) DeleteUserPreference(deleteKey int, userId int64) bool {
	defer panichandler.Recover()
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[1], strconv.FormatInt(userId, 10))
	return true
}

//公司的唯一id
func(dbObj *DB) DeleteCompany(deleteKey int, companyId int64) bool {
	defer panichandler.Recover()
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[2], strconv.FormatInt(companyId, 10))
	return true
}

// 公司部門的id
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id int64) bool {
	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[3], strconv.FormatInt(id, 10))
	return true
}

// 班表的唯一id
func(dbObj *DB) DeleteShift(deleteKey int, shiftId int64) bool {
	defer panichandler.Recover()
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[4], strconv.FormatInt(shiftId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[5], strconv.FormatInt(caseId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[6], strconv.FormatInt(caseId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteDayOff(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[7], strconv.FormatInt(caseId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[8], strconv.FormatInt(caseId, 10))
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[9], strconv.FormatInt(caseId, 10))
	return true
}



//  ------------------------------insert------------------------------
func(dbObj *DB) InsertUser(data *table.UserTable) bool {
	defer panichandler.Recover()
	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()
	key := strconv.FormatInt((*data).UserId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[0], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertUserPreference(data *table.UserPreferenceTable) bool {
	defer panichandler.Recover()
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	key := strconv.FormatInt((*data).UserId, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[1], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertCompany(data *table.CompanyTable) bool {
	defer panichandler.Recover()
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	key := strconv.FormatInt((*data).CompanyId, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[2], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertCompanyBanch(data *table.CompanyBanchTable) bool {
	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	key := strconv.FormatInt((*data).Id, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[3], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertShift(data *table.ShiftTable) bool {
	defer panichandler.Recover()
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	key := strconv.FormatInt((*data).ShiftId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[4], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertShiftChange(data *table.ShiftChangeTable) bool {
	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[5], key,jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertShiftOverTime(data *table.ShiftOverTimeTable) bool {
	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[6], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertDayOff(data *table.DayOffTable) bool {
	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[7], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertForgetPunch(data *table.ForgetPunchTable) bool {
	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[8], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertLateExcused(data *table.LateExcusedTable) bool {
	defer panichandler.Recover()
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
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[0])
}
func(dbObj *DB) DeleteKeyUserPreference(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[1])
}
func(dbObj *DB) DeleteKeyCompany(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[2])
}
func(dbObj *DB) DeleteKeyCompanyBanch(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[3])
}
func(dbObj *DB) DeleteKeyShift(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[4])
}
func(dbObj *DB) DeleteKeyShiftChange(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[5])
}
func(dbObj *DB) DeleteKeyShiftOverTime(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[6])
}
func(dbObj *DB) DeleteKeyDayOff(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[7])
}
func(dbObj *DB) DeleteKeyForgetPunch(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[8])
}
func(dbObj *DB) DeleteKeyLateExcused(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[9])
}


//  ------------------------------method------------------------------
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
	for _, v:= range res {
		returnValue = append(returnValue, v.(string))
	}
	return returnValue, err
}