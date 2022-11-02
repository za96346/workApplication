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
		}
	}
	return redisInstance
}

func(dbObj *DB) Conn(path string) { // 實體化redis.Client 並返回實體的位址
	defer panichandler.Recover()
	err := godotenv.Load(path)
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

//  ------------------------------select------------------------------
		//  select all  //

// 0 => 全部, value =>  nil
//  1 =>  userId, value => int64
//  2 => account, value => string
// 3 => companyCode, value => string
//  4 => banch, value = > int64
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
	case 3:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.UserTable) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyCode {
						return true
					}
				}
				return false
			},
		)
	case 4:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.UserTable) bool {
				for _, filterItem := range value {
					if filterItem == v.Banch {
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
					if filterItem == v.CompanyId {
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
//  2 => initiatorShiftId, value => int64
//  3 => requestedShiftId, value => int64
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
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftChangeTable) bool {
				for _, filterItem := range value {
					if filterItem == v.InitiatorShiftId {
						return true
					}
				}
				return false
			},
		)
	case 3:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftChangeTable) bool {
				for _, filterItem := range value {
					if filterItem == v.RequestedShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.ShiftChangeTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
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
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftOverTimeTable) bool {
				for _, filterItem := range value {
					if filterItem == v.ShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.ShiftOverTimeTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
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
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.DayOffTable) bool {
				for _, filterItem := range value {
					if filterItem == v.ShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.DayOffTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
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
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ForgetPunchTable) bool {
				for _, filterItem := range value {
					if filterItem == v.ShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.ForgetPunchTable{}
	}
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
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
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.LateExcusedTable) bool {
				for _, filterItem := range value {
					if filterItem == v.ShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.LateExcusedTable{}
	}
}

// 0 => all, value => nil
//  1 => styleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchStyle(selectKey int, value... interface{}) *[]table.BanchStyle {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[10]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.BanchStyle) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.BanchStyle) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.BanchStyle) bool {
				for _, filterItem := range value {
					if filterItem == v.BanchId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.BanchStyle{}
	}
}

// 0 => all, value => nil
//  1 => ruleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchRule(selectKey int, value... interface{}) *[]table.BanchRule {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[11]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.BanchRule) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.BanchRule) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.BanchRule) bool {
				for _, filterItem := range value {
					if filterItem == v.BanchId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.BanchRule{}
	}
}

// 0 => all, value => nil
//  1 => quitId, value => int64
//   2 => userId, value => int64
//   3 => companyCode, value => string 
//   4=> companyCode && userId ,  value string && int64
func(dbObj *DB) SelectQuitWorkUser(selectKey int, value... interface{}) *[]table.QuitWorkUser {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[12]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.QuitWorkUser) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.QuitWorkUser) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.QuitWorkUser) bool {
				for _, filterItem := range value {
					if filterItem == v.UserId {
						return true
					}
				}
				return false
			},
		)
	case 3:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.QuitWorkUser) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyCode {
						return true
					}
				}
				return false
			},
		)
	case 4:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.QuitWorkUser) bool {
				if v.CompanyCode == value[0] && v.UserId == value[1] {
					return true
				}
				return false
			},
		)
	default:
		return &[]table.QuitWorkUser{}
	}
}

// 0 => all, value => nil
//  1 => waitId, value => int64
//  2 => userId, value => int64
//  3 => companyId, value => int64
//  4 => comapnyId && userId, value => int64, int64
func(dbObj *DB) SelectWaitCompanyReply (selectKey int, value... interface{}) *[]table.WaitCompanyReply {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[13]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WaitCompanyReply) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.WaitCompanyReply) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WaitCompanyReply) bool {
				for _, filterItem := range value {
					if filterItem == v.UserId {
						return true
					}
				}
				return false
			},
		)
	case 3:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WaitCompanyReply) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyId {
						return true
					}
				}
				return false
			},
		)
	case 4:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WaitCompanyReply) bool {
				if value[0] == v.CompanyId && value[1] == v.UserId {
					return true
				}
				return false
			},
		)
	default:
		return &[]table.WaitCompanyReply{}
	}
}

// 0 => all, value => nil
//  1 => weekendId, value => int64
//  2 => companyId, value => int64
func(dbObj *DB) SelectWeekendSetting (selectKey int, value... interface{}) *[]table.WeekendSetting {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[14]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WeekendSetting) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.WeekendSetting) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WeekendSetting) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.WeekendSetting{}
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

	res := (*dbObj).SelectCompanyBanch(2, id)
	(*dbObj).RedisDb.HDel((*dbObj).table[3], strconv.FormatInt(id, 10))
	if len(*res) <= 0 {
		return true
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func ()  {
		defer wg.Done()
		banchStyle := (*dbObj).SelectBanchStyle(2, (*res)[0].Id)
		for _, v := range *banchStyle {
			(*dbObj).DeleteBanchStyle(0, v.StyleId)
		}
		
	}()
	go func ()  {
		defer wg.Done()
		banchRule := (*dbObj).SelectBanchRule(2, (*res)[0].Id)
		for _, v:= range *banchRule {
			(*dbObj).DeleteBanchRule(0, v.RuleId)
		}
	}()
	wg.Wait()
	return true
}

// 班表的唯一id
func(dbObj *DB) DeleteShift(deleteKey int, shiftId int64) bool {
	defer panichandler.Recover()
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()

	(*dbObj).RedisDb.HDel((*dbObj).table[4], strconv.FormatInt(shiftId, 10))

	wg := new(sync.WaitGroup)
	wg.Add(6)
	// (*dbObj).SelectShiftChange(1)
	go func ()  {
		defer wg.Done()
		res := (*dbObj).SelectShiftOverTime(2, shiftId)
		for _, v := range *res {
			(*dbObj).DeleteShiftOverTime(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res := (*dbObj).SelectDayOff(2, shiftId)
		for _, v := range *res {
			(*dbObj).DeleteDayOff(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res := (*dbObj).SelectLateExcused(2, shiftId)
		for _, v := range *res {
			(*dbObj).DeleteLateExcused(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res := (*dbObj).SelectForgetPunch(2, shiftId)
		for _, v := range *res {
			(*dbObj).DeleteForgetPunch(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res1 := (*dbObj).SelectShiftChange(2, shiftId)
		for _, v := range *res1 {
			(*dbObj).DeleteShiftChange(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res1 := (*dbObj).SelectShiftChange(3, shiftId)
		for _, v := range *res1 {
			(*dbObj).DeleteShiftChange(0, v.CaseId)
		}
	}()
	wg.Wait()
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

// style的唯一id
func(dbObj *DB) DeleteBanchStyle(deleteKey int, styleId int64) bool {
	defer panichandler.Recover()
	(*dbObj).banchStyleMux.Lock()
	defer (*dbObj).banchStyleMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[10], strconv.FormatInt(styleId, 10))
	return true
}

// style的唯一id
func(dbObj *DB) DeleteBanchRule(deleteKey int, ruleId int64) bool {
	defer panichandler.Recover()
	(*dbObj).banchRuleMux.Lock()
	defer (*dbObj).banchRuleMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[11], strconv.FormatInt(ruleId, 10))
	return true
}

// quit work suer 的唯一id
func(dbObj *DB) DeleteQuitWorkUser(deleteKey int, quitId int64) bool {
	defer panichandler.Recover()
	(*dbObj).quitWorkUserMux.Lock()
	defer (*dbObj).quitWorkUserMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[12], strconv.FormatInt(quitId, 10))
	return true
}

// wait company reply 的唯一id
func(dbObj *DB) DeleteWaitCompanyReply(deleteKey int, waitId int64) bool {
	defer panichandler.Recover()
	(*dbObj).waitCompanyReply.Lock()
	defer (*dbObj).waitCompanyReply.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[13], strconv.FormatInt(waitId, 10))
	return true
}

// weekend setting 的唯一id
func(dbObj *DB) DeleteWeekendSetting(deleteKey int, weekendId int64) bool {
	defer panichandler.Recover()
	(*dbObj).weekendSetting.Lock()
	defer (*dbObj).weekendSetting.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[14], strconv.FormatInt(weekendId, 10))
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

func(dbObj *DB) InsertBanchStyle(data *table.BanchStyle) bool {
	defer panichandler.Recover()
	(*dbObj).banchStyleMux.Lock()
	defer (*dbObj).banchStyleMux.Unlock()
	key := strconv.FormatInt((*data).StyleId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[10], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertBanchRule(data *table.BanchRule) bool {
	defer panichandler.Recover()
	(*dbObj).banchRuleMux.Lock()
	defer (*dbObj).banchRuleMux.Unlock()
	key := strconv.FormatInt((*data).RuleId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[11], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertQuitWorkUser(data *table.QuitWorkUser) bool {
	defer panichandler.Recover()
	(*dbObj).quitWorkUserMux.Lock()
	defer (*dbObj).quitWorkUserMux.Unlock()
	key := strconv.FormatInt((*data).QuitId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[12], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertWaitCompanyReply(data *table.WaitCompanyReply) bool {
	defer panichandler.Recover()
	(*dbObj).waitCompanyReply.Lock()
	defer (*dbObj).waitCompanyReply.Unlock()
	key := strconv.FormatInt((*data).WaitId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[13], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) InsertWeekendSetting(data *table.WeekendSetting) bool {
	defer panichandler.Recover()
	(*dbObj).weekendSetting.Lock()
	defer (*dbObj).weekendSetting.Unlock()
	key := strconv.FormatInt((*data).WeekendId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[14], key, jsonData).Result()
	
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

func(dbObj *DB) DeleteKeyBanchStyle(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[10])
}
func(dbObj *DB) DeleteKeyBanchRule(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[11])
}
func(dbObj *DB) DeleteKeyQuitWorkUser(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[12])
}
func(dbObj *DB) DeleteKeyWaitCompanyReply(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[13])
}
func(dbObj *DB) DeleteKeyWeekendSetting(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[14])
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