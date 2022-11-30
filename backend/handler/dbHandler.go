package handler

import (
	"backend/mysql"
	"backend/redis"
	"fmt"

	// "fmt"
	"sync"
	// "time"
	"backend/panicHandler"
)

func Init(path string) {
	(*mysql.Singleton()).Conn(path) //mysql 連接
	(*redis.Singleton()).Conn(path) // redis 連接
	(*redis.Singleton()).RedisDb.FlushDB() // redis清空
	// (*Singleton()).TakeAllFromMysql() // 從mysql 抓到 redis
	// (*Singleton()).DeleteCompanyBanch(0, int64(1))
	// (*Singleton()).DeleteShift(0, 2)
}

var dbHandlerInstance *DB
var dbHandlerInstanceMux = new(sync.Mutex)

type DB struct {
	Redis *redis.DB
	Mysql *mysql.DB
	userLock *bool
	userPreferenceLock *bool
	compnayLock *bool
	compnayBanchLock *bool
	shiftLock *bool
	shiftChangeLock *bool
	shiftOverTimeLock *bool
	dayOffLock *bool
	forgetPunchLock *bool
	lateExcusedLock *bool
	banchStyleLock *bool
	banchRuleLock *bool
	quitWorkUserLock *bool
	waitCompanyReplyLock *bool
	weekendSettingLock *bool
	workTime *bool
	paidVocation *bool
}

func newBool(b bool) *bool {
    return &b
}

func Singleton() *DB {
	defer panichandler.Recover()
	if dbHandlerInstance == nil {
		dbHandlerInstanceMux.Lock()
		defer dbHandlerInstanceMux.Unlock()
		if dbHandlerInstance == nil {
			dbHandlerInstance = &DB{
				Redis: redis.Singleton(),
				Mysql: mysql.Singleton(),
				userLock: newBool(false),
				userPreferenceLock: newBool(false),
				compnayLock: newBool(false),
				compnayBanchLock: newBool(false),
				shiftLock: newBool(false),
				shiftChangeLock: newBool(false),
				shiftOverTimeLock: newBool(false),
				dayOffLock: newBool(false),
				forgetPunchLock: newBool(false),
				lateExcusedLock: newBool(false),
				banchStyleLock: newBool(false),
				banchRuleLock: newBool(false),
				quitWorkUserLock: newBool(false),
				waitCompanyReplyLock: newBool(false),
				weekendSettingLock: newBool(false),
				workTime: newBool(false),
				paidVocation: newBool(false),
			}
		}
	}
	return dbHandlerInstance
}

func(dbObj *DB) TakeAllFromMysql() {
	defer panichandler.Recover()
	(*dbObj).restoreUserAll()
	(*dbObj).restoreUserPreferenceAll()
	(*dbObj).restoreCompanyAll()
	(*dbObj).restoreCompanyBanchAll()
	(*dbObj).restoreShiftAll()
	(*dbObj).restoreShiftChangeAll()
	(*dbObj).restoreShiftOverTimeAll()
	(*dbObj).restoreDayOffAll()
	(*dbObj).restoreForgetPunchAll()
	(*dbObj).restoreLateExcusedAll()
	(*dbObj).restoreBanchStyleAll()
	(*dbObj).restoreBanchRuleAll()
	(*dbObj).restoreQuitWorkUserAll()
	(*dbObj).restoreWaitCompanyReplyAll()
	(*dbObj).restoreWeekendSettingAll()

}


func forEach[T any, callbackT any](data *[]T, callback func(*T) callbackT) {
	defer panichandler.Recover()
	for _, value := range *data {
		_ = callback(&value)
	}
}

func selectAllHandler[
		callbackT any,
	](
		redisCallback func() *[]callbackT,
		mysqlCallback func() *[]callbackT,
		isLocked *bool,
	) *[]callbackT {

	defer panichandler.Recover()
	if (*redis.Singleton()).IsAlive() && !(*isLocked) && false {
		// redis
		fmt.Println("從 redis 拿取資料 開始")
		res := redisCallback()
		return res

	} else {
		// mysql
		fmt.Println("從 mysql 拿取資料 開始")
		res := mysqlCallback()
		return res
	}
}