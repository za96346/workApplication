package handler

import (
	"backend/mysql"
	"backend/redis"
	"backend/table"
	"fmt"

	// "fmt"
	"sync"
	// "time"
	"backend/panicHandler"
)

func Init(path string) {
	(*mysql.Singleton()).Conn(path) //mysql 連接
	(*redis.Singleton()).Conn(path) // redis 連接
	(*redis.Singleton()).RedisDb.FlushAll() // redis清空
	(*Singleton()).TakeAllFromMysql() // 從mysql 抓到 redis
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
	if (*redis.Singleton()).IsAlive() && !(*isLocked) {
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












//  ------------------------------clear and reStore to redis------------------------------

func(dbObj *DB) restoreUserAll() {
	defer panichandler.Recover()
	(*(*dbObj).userLock) = true
	(*dbObj).Redis.DeleteKeyUser()
	arr := (*dbObj.Mysql).SelectUser(0)
	forEach(arr, (*dbObj.Redis).InsertUser)
	(*(*dbObj).userLock) = false
}
func(dbObj *DB) restoreUserPreferenceAll() {
	defer panichandler.Recover()
	(*(*dbObj).userPreferenceLock) = true
	(*dbObj).Redis.DeleteKeyUserPreference()
	arr := (*dbObj.Mysql).SelectUserPreference(0)
	forEach(arr, (*dbObj.Redis).InsertUserPreference)
	(*(*dbObj).userPreferenceLock) = false
}
func(dbObj *DB) restoreCompanyAll() {
	defer panichandler.Recover()
	(*(*dbObj).compnayLock)= true
	(*dbObj).Redis.DeleteKeyCompany()
	arr := (*dbObj.Mysql).SelectCompany(0)
	forEach(arr, (*dbObj.Redis).InsertCompany)
	(*(*dbObj).compnayLock) = false
}
func(dbObj *DB) restoreCompanyBanchAll() {
	defer panichandler.Recover()
	(*(*dbObj).compnayBanchLock)= true
	(*dbObj).Redis.DeleteKeyCompanyBanch()
	arr := (*dbObj.Mysql).SelectCompanyBanch(0)
	forEach(arr, (*dbObj.Redis).InsertCompanyBanch)
	(*(*dbObj).compnayBanchLock) = false
}
func(dbObj *DB) restoreShiftAll() {
	defer panichandler.Recover()
	(*(*dbObj).shiftLock) = true
	(*dbObj).Redis.DeleteKeyShift()
	arr := (*dbObj.Mysql).SelectShift(0)
	forEach(arr, (*dbObj.Redis).InsertShift)
	(*(*dbObj).shiftLock) = false
}
func(dbObj *DB) restoreShiftChangeAll() {
	defer panichandler.Recover()
	(*(*dbObj).shiftChangeLock) = true
	(*dbObj).Redis.DeleteKeyShiftChange()
	arr := (*dbObj.Mysql).SelectShiftChange(0)
	forEach(arr, (*dbObj.Redis).InsertShiftChange)
	(*(*dbObj).shiftChangeLock) = false
}
func(dbObj *DB) restoreShiftOverTimeAll() {
	defer panichandler.Recover()
	(*(*dbObj).shiftOverTimeLock) = true
	(*dbObj).Redis.DeleteKeyShiftOverTime()
	arr := (*dbObj.Mysql).SelectShiftOverTime(0)
	forEach(arr, (*dbObj.Redis).InsertShiftOverTime)
	(*(*dbObj).shiftOverTimeLock) = false
}
func(dbObj *DB) restoreDayOffAll() {
	defer panichandler.Recover()
	(*(*dbObj).dayOffLock) = true
	(*dbObj).Redis.DeleteKeyDayOff()
	arr := (*dbObj.Mysql).SelectDayOff(0)
	forEach(arr, (*dbObj.Redis).InsertDayOff)
	(*(*dbObj).dayOffLock) = false
}
func(dbObj *DB) restoreForgetPunchAll() {
	defer panichandler.Recover()
	(*(*dbObj).forgetPunchLock) = true
	(*dbObj).Redis.DeleteKeyForgetPunch()
	arr := (*dbObj.Mysql).SelectForgetPunch(0)
	forEach(arr, (*dbObj.Redis).InsertForgetPunch)
	(*(*dbObj).forgetPunchLock)= false
}
func(dbObj *DB) restoreLateExcusedAll() {
	defer panichandler.Recover()
	(*(*dbObj).lateExcusedLock) = true
	(*dbObj).Redis.DeleteKeyLateExcused()
	arr := (*dbObj.Mysql).SelectLateExcused(0)
	forEach(arr, (*dbObj.Redis).InsertLateExcused)
	(*(*dbObj).lateExcusedLock) = false
}

func(dbObj *DB) restoreBanchStyleAll() {
	defer panichandler.Recover()
	(*(*dbObj).banchStyleLock) = true
	(*dbObj).Redis.DeleteKeyBanchStyle()
	arr := (*dbObj.Mysql).SelectBanchStyle(0)
	forEach(arr, (*dbObj.Redis).InsertBanchStyle)
	(*(*dbObj).banchStyleLock) = false
}

func(dbObj *DB) restoreBanchRuleAll() {
	defer panichandler.Recover()
	(*(*dbObj).banchRuleLock) = true
	(*dbObj).Redis.DeleteKeyBanchRule()
	arr := (*dbObj.Mysql).SelectBanchRule(0)
	forEach(arr, (*dbObj.Redis).InsertBanchRule)
	(*(*dbObj).banchRuleLock) = false
}

func(dbObj *DB) restoreQuitWorkUserAll() {
	defer panichandler.Recover()
	(*(*dbObj).quitWorkUserLock) = true
	(*dbObj).Redis.DeleteKeyQuitWorkUser()
	arr := (*dbObj.Mysql).SelectQuitWorkUser(0)
	forEach(arr, (*dbObj.Redis).InsertQuitWorkUser)
	(*(*dbObj).quitWorkUserLock) = false
}






//  ------------------------------insert------------------------------


func(dbObj *DB) InsertUser(data *table.UserTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertUser(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectUser(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertUser(&value)
			}
		}()
	}
	return isOk, id
}
func(dbObj *DB) InsertUserPreference(data *table.UserPreferenceTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, _ := (*dbObj).Mysql.InsertUserPreference(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectUserPreference(1, (*data).UserId)
			for _, value := range *res {
				(*dbObj).Redis.InsertUserPreference(&value)
			}
		}()
	}
	return isOk, (*data).UserId
}
func(dbObj *DB) InsertCompany(data *table.CompanyTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertCompany(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectCompany(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertCompany(&value)	
			}
		}()
	}
	return isOk, id
}
func(dbObj *DB) InsertCompanyBanch(data *table.CompanyBanchTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertCompanyBanch(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectCompanyBanch(2, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertCompanyBanch(&value)
			}
		}()
	}
	return isOk, id
}
func(dbObj *DB) InsertShift(data *table.ShiftTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertShift(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShift(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertShift(&value)
			}
			
		}()
	}
	return isOk, id
}
func(dbObj *DB) InsertShiftChange(data *table.ShiftChangeTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertShiftChange(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShiftChange(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertShiftChange(&value)
			}
		}()
	}
	return isOk, id
}
func(dbObj *DB) InsertShiftOverTime(data *table.ShiftOverTimeTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (dbObj).Mysql.InsertShiftOverTime(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShiftOverTime(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertShiftOverTime(&value)
			}
			
		}()
	}
	return isOk, id
}
func(dbObj *DB) InsertDayOff(data *table.DayOffTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertDayOff(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectDayOff(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertDayOff(&value)
			}
		}()
	}
	return isOk, id
}
func(dbObj *DB) InsertForgetPunch(data *table.ForgetPunchTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertForgetPunch(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectForgetPunch(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertForgetPunch(&value)
			}
			
		}()
	}
	return isOk, id
}
func(dbObj *DB) InsertLateExcused(data *table.LateExcusedTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertLateExcused(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectLateExcused(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertLateExcused(&value)
			}
		}()
	}
	return isOk, id
}

func(dbObj *DB) InsertBanchStyle(data *table.BanchStyle) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertBanchStyle(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectBanchStyle(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertBanchStyle(&value)
			}
		}()
	}
	return isOk, id
}

func(dbObj *DB) InsertBanchRule(data *table.BanchRule) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertBanchRule(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectBanchRule(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertBanchRule(&value)
			}
		}()
	}
	return isOk, id
}

func(dbObj *DB) InsertQuitWorkUser(data *table.QuitWorkUser) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertQuitWorkUser(data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectQuitWorkUser(1, id)
			for _, value := range *res {
				(*dbObj).Redis.InsertQuitWorkUser(&value)
			}
		}()
	}
	return isOk, id
}


//  ------------------------------update------------------------------


func(dbObj *DB) UpdateUser(updateKey int, data *table.UserTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateUser(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectUser(1, int64((*data).UserId))
			for _, v := range *res {
				(*dbObj).Redis.InsertUser(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateUserPreference(updateKey int, data *table.UserPreferenceTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateUserPreference(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectUserPreference(1, int64((*data).UserId))
			for _, v := range *res {
				(*dbObj).Redis.InsertUserPreference(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateCompany(updateKey int, data *table.CompanyTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateCompany(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectCompany(1, int64((*data).CompanyId))
			for _, v := range *res {
				(*dbObj).Redis.InsertCompany(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateCompanyBanch(updateKey int, data *table.CompanyBanchTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateCompanyBanch(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectCompanyBanch(2, int64((*data).Id))
			for _, v := range *res {
				(*dbObj).Redis.InsertCompanyBanch(&v)
			}

		}()
	}
	return isOk
}
func(dbObj *DB) UpdateShift(updateKey int, data *table.ShiftTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateShift(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShift(1, int64((*data).ShiftId))
			for _, v := range *res {
				(*dbObj).Redis.InsertShift(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateShiftChange(updateKey int, data *table.ShiftChangeTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateShiftChange(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShiftChange(1, int64((*data).CaseId))
			for _, v := range *res {
				(*dbObj).Redis.InsertShiftChange(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateShiftOverTime(updateKey int, data *table.ShiftOverTimeTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateShiftOverTime(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShiftOverTime(1, int64((*data).CaseId))
			for _, v := range *res {
				(*dbObj).Redis.InsertShiftOverTime(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateDayOff(updateKey int, data *table.DayOffTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateDayOff(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectDayOff(1, int64((*data).CaseId))
			for _, v := range *res {
				(*dbObj).Redis.InsertDayOff(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateForgetPunch(updateKey int, data *table.ForgetPunchTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateForgetPunch(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectForgetPunch(1, int64((*data).CaseId))
			for _, v := range *res {
				(*dbObj).Redis.InsertForgetPunch(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateLateExcused(updateKey int, data *table.LateExcusedTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateLateExcused(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectLateExcused(1, int64((*data).CaseId))
			for _, v := range *res {
				(*dbObj).Redis.InsertLateExcused(&v)
			}
		}()
	}
	return isOk
}
func(dbObj *DB) UpdateBanchStyle(updateKey int, data *table.BanchStyle) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateBanchStyle(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectBanchStyle(1, int64((*data).StyleId))
			for _, v := range *res {
				(*dbObj).Redis.InsertBanchStyle(&v)
			}
		}()
	}
	return isOk
}

func(dbObj *DB) UpdateBanchRule(updateKey int, data *table.BanchRule) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateBanchRule(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectBanchRule(1, int64((*data).RuleId))
			for _, v := range *res {
				(*dbObj).Redis.InsertBanchRule(&v)
			}
		}()
	}
	return isOk
}

func(dbObj *DB) UpdateQuitWorkUser(updateKey int, data *table.QuitWorkUser) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateQuitWorkUser(updateKey, data)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectQuitWorkUser(1, int64((*data).QuitId))
			for _, v := range *res {
				(*dbObj).Redis.InsertQuitWorkUser(&v)
			}
		}()
	}
	return isOk
}

//  ------------------------------delete------------------------------


// func(dbObj *DB) DeleteUser(deleteKey int, userId int64) bool {
// 	defer panichandler.Recover()
// 	res := (*dbObj).Mysql.DeleteUser(deleteKey, userId)
// 	if res {
// 		go func ()  {
// 			(*dbObj).TakeAllFromMysql()
// 		}()
// 	}
// 	return res
// }
func(dbObj *DB) DeleteUserPreference(deleteKey int, userId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteUserPreference(deleteKey, userId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteUserPreference(deleteKey, userId)
		}()	
	}
	return res
}
// func(dbObj *DB) DeleteCompany(deleteKey int, companyId int64) bool {
// 	defer panichandler.Recover()
// 	res := (*dbObj).Mysql.DeleteCompany(deleteKey, companyId)
// 	if res {
// 		go func ()  {
// 			(*dbObj).TakeAllFromMysql()
// 		}()	
// 	}
// 	return res
// }
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteCompanyBanch(deleteKey, id)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteCompanyBanch(deleteKey, id)
		}()	
	}
	return res
}
func(dbObj *DB) DeleteShift(deleteKey int, shiftId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteShift(deleteKey, shiftId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteShift(deleteKey, shiftId)
		}()	
	}
	return res
}
func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteShiftChange(deleteKey, caseId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteShiftChange(deleteKey, caseId)	
		}()	
	}
	return res
}
func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteShiftOverTime(deleteKey, caseId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteShiftOverTime(deleteKey, caseId)
		}()	
	}
	
	return res
}
func(dbObj *DB) DeleteDayOff(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteDayOff(deleteKey, caseId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteDayOff(deleteKey, caseId)
		}()	
	}
	
	return res
}
func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteForgetPunch(deleteKey, caseId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteForgetPunch(deleteKey, caseId)
		}()	
	}
	
	return res
}
func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteLateExcused(deleteKey, caseId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteLateExcused(deleteKey, caseId)
		}()	
	}
	
	return res
}

func(dbObj *DB) DeleteBanchStyle(deleteKey int, styleId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteBanchStyle(deleteKey, styleId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteBanchStyle(deleteKey, styleId)
		}()	
	}
	
	return res
}

func(dbObj *DB) DeleteBanchRule(deleteKey int, ruleId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteBanchRule(deleteKey, ruleId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteBanchRule(deleteKey, ruleId)
		}()	
	}
	
	return res
}

func(dbObj *DB) DeleteQuitWorkUser(deleteKey int, quitId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteQuitWorkUser(deleteKey, quitId)
	if res {
		go func ()  {
			(*dbObj).Redis.DeleteQuitWorkUser(deleteKey, quitId)
		}()	
	}
	
	return res
}




//  ------------------------------select------------------------------

// 0 => 全部, value => nil
//  1 =>  userId, value => int64
//  2 => account, value => string
// 3 => companyCode, value => string
func(dbObj *DB) SelectUser(selectKey int, value... interface{}) *[]table.UserTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.UserTable {
			return (*dbObj.Redis).SelectUser(selectKey, value...)
		},
		func() *[]table.UserTable {
			return (*dbObj.Mysql).SelectUser(selectKey, value...)
		},
		(*dbObj).userLock,
	)
}

// 0 => 全部, value => nil
//  1 => 使用者id, value => int64
func(dbObj *DB) SelectUserPreference(selectKey int, value... interface{}) *[]table.UserPreferenceTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.UserPreferenceTable {
			return (*dbObj.Redis).SelectUserPreference(selectKey, value...)
		},
		func() *[]table.UserPreferenceTable {
			return (*dbObj.Mysql).SelectUserPreference(selectKey, value...)
		},
		(*dbObj).userPreferenceLock,
	)
}

// 0 => 全部, value => nil
//  1 => 公司id, value => int64
//  2 => 公司碼, value => string
func(dbObj *DB) SelectCompany(selectKey int, value... interface{}) *[]table.CompanyTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.CompanyTable {
			return (*dbObj.Redis).SelectCompany(selectKey, value...)
		},
		func() *[]table.CompanyTable {
			return (*dbObj.Mysql).SelectCompany(selectKey, value...)
		},
		(*dbObj).compnayLock,
	)
}

// 0 => 全部, value => nil
//	1 => 公司Id, value => int64
// 	2 => id (banchId), value => int64
func(dbObj *DB) SelectCompanyBanch(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.CompanyBanchTable {
			return (*dbObj.Redis).SelectCompanyBanch(selectKey, value...)
		},
		func() *[]table.CompanyBanchTable {
			return (*dbObj.Mysql).SelectCompanyBanch(selectKey, value...)
		},
		(*dbObj).compnayBanchLock,
	)
}

// 0 => all, value => nil
//  1 => 班表id, value => int64
func(dbObj *DB) SelectShift(selectKey int, value... interface{}) *[]table.ShiftTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.ShiftTable {
			return (*dbObj.Redis).SelectShift(selectKey, value...)
		},
		func() *[]table.ShiftTable {
			return (*dbObj.Mysql).SelectShift(selectKey, value...)
		},
		(*dbObj).shiftLock,
	)
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => initiatorShiftId, value => int64
//  3 => requestedShiftId, value => int64
func(dbObj *DB) SelectShiftChange(selectKey int, value... interface{}) *[]table.ShiftChangeTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.ShiftChangeTable {
			return (*dbObj.Redis).SelectShiftChange(selectKey, value...)
		},
		func() *[]table.ShiftChangeTable {
			return (*dbObj.Mysql).SelectShiftChange(selectKey, value...)
		},
		(*dbObj).shiftChangeLock,
	)
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
func(dbObj *DB) SelectShiftOverTime(selectKey int, value... interface{}) *[]table.ShiftOverTimeTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.ShiftOverTimeTable {
			return (*dbObj.Redis).SelectShiftOverTime(selectKey, value...)
		},
		func() *[]table.ShiftOverTimeTable {
			return (*dbObj.Mysql).SelectShiftOverTime(selectKey, value...)
		},
		(*dbObj).shiftOverTimeLock,
	)
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
func(dbObj *DB) SelectDayOff(selectKey int, value... interface{}) *[]table.DayOffTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.DayOffTable {
			return (*dbObj.Redis).SelectDayOff(selectKey, value...)
		},
		func() *[]table.DayOffTable {
			return (*dbObj.Mysql).SelectDayOff(selectKey, value...)
		},
		(*dbObj).dayOffLock,
	)
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
func(dbObj *DB) SelectForgetPunch(selectKey int, value... interface{}) *[]table.ForgetPunchTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.ForgetPunchTable {
			return (*dbObj.Redis).SelectForgetPunch(selectKey, value...)
		},
		func() *[]table.ForgetPunchTable {
			return (*dbObj.Mysql).SelectForgetPunch(selectKey, value...)
		},
		(*dbObj).forgetPunchLock,
	)
}

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
func(dbObj *DB) SelectLateExcused(selectKey int, value... interface{}) *[]table.LateExcusedTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.LateExcusedTable {
			return (*dbObj.Redis).SelectLateExcused(selectKey, value...)
		},
		func() *[]table.LateExcusedTable {
			return (*dbObj.Mysql).SelectLateExcused(selectKey, value...)
		},
		(*dbObj).lateExcusedLock,
	)
}

// 0 => all, value => nil
//  1 => styleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchStyle(selectKey int, value... interface{}) *[]table.BanchStyle {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.BanchStyle {
			return (*dbObj.Redis).SelectBanchStyle(selectKey, value...)
		},
		func() *[]table.BanchStyle {
			return (*dbObj.Mysql).SelectBanchStyle(selectKey, value...)
		},
		(*dbObj).banchStyleLock,
	)
}

// 0 => all, value => nil
//  1 => ruleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchRule(selectKey int, value... interface{}) *[]table.BanchRule {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.BanchRule {
			return (*dbObj.Redis).SelectBanchRule(selectKey, value...)
		},
		func() *[]table.BanchRule {
			return (*dbObj.Mysql).SelectBanchRule(selectKey, value...)
		},
		(*dbObj).banchRuleLock,
	)
}

// 0 => all, value => nil
//  1 => quitId, value => int64
//   2 => userId, value => int64
//   3 => companyCode, value => string 
//   4=> companyCode && userId ,  value string && int64
func(dbObj *DB) SelectQuitWorkUser(selectKey int, value... interface{}) *[]table.QuitWorkUser {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.QuitWorkUser {
			return (*dbObj.Redis).SelectQuitWorkUser(selectKey, value...)
		},
		func() *[]table.QuitWorkUser {
			return (*dbObj.Mysql).SelectQuitWorkUser(selectKey, value...)
		},
		(*dbObj).quitWorkUserLock,
	)
}