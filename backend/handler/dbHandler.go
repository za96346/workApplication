package handler

import (
	"backend/mysql"
	"backend/redis"
	"backend/table"
	"fmt"

	// "fmt"
	"sync"
)

func Init() {
	(*mysql.Singleton()).Conn() //mysql 連接
	(*redis.Singleton()).Conn() // redis 連接
	// (*redis.Singleton()).RedisDb.FlushAll() // redis清空
	(*Singleton()).TakeAllFromMysql() // 從mysql 抓到 redis
}

var dbHandlerInstance *DB
var dbHandlerInstanceMux = new(sync.Mutex)

type DB struct {
	Redis *redis.DB
	Mysql *mysql.DB
	userLock bool
	userPreferenceLock bool
	compnayLock bool
	compnayBanchLock bool
	shiftLock bool
	shiftChangeLock bool
	shiftOverTimeLock bool
	dayOffLock bool
	forgetPunchLock bool
	lateExcusedLock bool
}
func Singleton() *DB {
	if dbHandlerInstance == nil {
		dbHandlerInstanceMux.Lock()
		if dbHandlerInstance == nil {
			defer dbHandlerInstanceMux.Unlock()
			dbHandlerInstance = &DB{
				Redis: redis.Singleton(),
				Mysql: mysql.Singleton(),
				userLock: false,
				userPreferenceLock: false,
				compnayLock: false,
				compnayBanchLock: false,
				shiftLock: false,
				shiftChangeLock: false,
				shiftOverTimeLock: false,
				dayOffLock: false,
				forgetPunchLock: false,
				lateExcusedLock: false,
			}
		}
	}
	return dbHandlerInstance
}

func(dbObj *DB) TakeAllFromMysql() {
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

	(*dbObj).SelectUser(0)
	(*dbObj).SelectCompany(0)
	(*dbObj).SelectUserPreference(0)
	(*dbObj).SelectCompanyBanch(0)
	(*dbObj).SelectShift(0)
	(*dbObj).SelectShiftChange(0)
	(*dbObj).SelectShiftOverTime(0)
	(*dbObj).SelectDayOff(0)
	(*dbObj).SelectForgetPunch(0)
	(*dbObj).SelectLateExcused(0)

	res := (*dbObj).Mysql.SelectUserSingle(1, 0)
	fmt.Println("選擇單個user", res)
}


func forEach[T any, callbackT any](data *[]T, callback func(*T) callbackT) {
	for _, value := range *data {
		_ = callback(&value)
	}
}

func selectAllHandler[
		callbackT any,
	](
		redisCallback func() *[]callbackT,
		mysqlCallback func() *[]callbackT,
		isLocked bool,
	) *[]callbackT {

	if (*redis.Singleton()).IsAlive() && !isLocked {
		// redis
		res := redisCallback()
		return res

	} else {
		// mysql
		res := mysqlCallback()
		return res
	}
}












//  ------------------------------clear and reStore to redis------------------------------

func(dbObj *DB) restoreUserAll() {
	(*dbObj).userLock = true
	(*dbObj).Redis.DeleteKeyUser()
	arr := (*dbObj.Mysql).SelectUserAll(0)
	forEach(arr, (*dbObj.Redis).InsertUser)
	(*dbObj).userLock = false
}
func(dbObj *DB) restoreUserPreferenceAll() {
	(*dbObj).userPreferenceLock = true
	(*dbObj).Redis.DeleteKeyUserPreference()
	arr := (*dbObj.Mysql).SelectUserPreferenceAll(0)
	forEach(arr, (*dbObj.Redis).InsertUserPreference)
	(*dbObj).userPreferenceLock = false
}
func(dbObj *DB) restoreCompanyAll() {
	(*dbObj).compnayLock= true
	(*dbObj).Redis.DeleteKeyCompany()
	arr := (*dbObj.Mysql).SelectCompanyAll(0)
	forEach(arr, (*dbObj.Redis).InsertCompany)
	(*dbObj).compnayLock= false
}
func(dbObj *DB) restoreCompanyBanchAll() {
	(*dbObj).compnayBanchLock= true
	(*dbObj).Redis.DeleteKeyCompanyBanch()
	arr := (*dbObj.Mysql).SelectCompanyBanchAll(0)
	forEach(arr, (*dbObj.Redis).InsertCompanyBanch)
	(*dbObj).compnayBanchLock= false
}
func(dbObj *DB) restoreShiftAll() {
	(*dbObj).shiftLock= true
	(*dbObj).Redis.DeleteKeyShift()
	arr := (*dbObj.Mysql).SelectShiftAll(0)
	forEach(arr, (*dbObj.Redis).InsertShift)
	(*dbObj).shiftLock= false
}
func(dbObj *DB) restoreShiftChangeAll() {
	(*dbObj).shiftChangeLock= true
	(*dbObj).Redis.DeleteKeyShiftChange()
	arr := (*dbObj.Mysql).SelectShiftChangeAll(0)
	forEach(arr, (*dbObj.Redis).InsertShiftChange)
	(*dbObj).shiftChangeLock= false
}
func(dbObj *DB) restoreShiftOverTimeAll() {
	(*dbObj).shiftOverTimeLock= true
	(*dbObj).Redis.DeleteKeyShiftOverTime()
	arr := (*dbObj.Mysql).SelectShiftOverTimeAll(0)
	forEach(arr, (*dbObj.Redis).InsertShiftOverTime)
	(*dbObj).shiftOverTimeLock= false
}
func(dbObj *DB) restoreDayOffAll() {
	(*dbObj).dayOffLock= true
	(*dbObj).Redis.DeleteKeyDayOff()
	arr := (*dbObj.Mysql).SelectDayOffAll(0)
	forEach(arr, (*dbObj.Redis).InsertDayOff)
	(*dbObj).dayOffLock= false
}
func(dbObj *DB) restoreForgetPunchAll() {
	(*dbObj).forgetPunchLock= true
	(*dbObj).Redis.DeleteKeyForgetPunch()
	arr := (*dbObj.Mysql).SelectForgetPunchAll(0)
	forEach(arr, (*dbObj.Redis).InsertForgetPunch)
	(*dbObj).forgetPunchLock= false
}
func(dbObj *DB) restoreLateExcusedAll() {
	(*dbObj).lateExcusedLock= true
	(*dbObj).Redis.DeleteKeyLateExcused()
	arr := (*dbObj.Mysql).SelectLateExcusedAll(0)
	forEach(arr, (*dbObj.Redis).InsertLateExcused)
	(*dbObj).lateExcusedLock= false
}






//  ------------------------------insert------------------------------


func(dbObj *DB) InsertUser(data *table.UserTable) (bool, int64) {
	isOk, id := (*dbObj).Mysql.InsertUser(
		(*data).CompanyCode,
		(*data).Account,
		(*data).Password,
		(*data).OnWorkDay,
		(*data).Banch,
		(*data).Permession,
		(*data).WorkState,
		(*data).CreateTime,
		(*data).LastModify,
		int((*data).MonthSalary),
		int((*data).PartTimeSalary),
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectUserSingle(0, id)
			(*dbObj).Redis.InsertUser(res)
		}()
		return true, id
	}
	return false, id
}
func(dbObj *DB) InsertUserPreference(data *table.UserPreferenceTable) (bool, int64) {
	isOk, _ := (*dbObj).Mysql.InsertUserPreference(
		(*data).UserId,
		(*data).Style,
		(*data).FontSize,
		(*data).SelfPhoto,
		(*data).CreateTime,
		(*data).LastModify,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectUserPreferenceSingle(0, (*data).UserId)
			(*dbObj).Redis.InsertUserPreference(res)
		}()
		return true, (*data).UserId
	}
	return false, (*data).UserId
}
func(dbObj *DB) InsertCompany(data *table.CompanyTable) (bool, int64) {
	isOk, id := (*dbObj).Mysql.InsertCompany(
		(*data).CompanyCode,
		(*data).CompanyName,
		(*data).CompanyLocation,
		(*data).CompanyPhoneNumber,
		(*data).TermStart,
		(*data).TermEnd,
		(*data).CreateTime,
		(*data).LastModify,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectCompanySingle(0, id)
			(*dbObj).Redis.InsertCompany(res)	
		}()
		return true, id
	}
	return false, id
}
func(dbObj *DB) InsertCompanyBanch(data *table.CompanyBanchTable) (bool, int64) {
	isOk, id := (*dbObj).Mysql.InsertCompanyBanch(
		(*data).CompanyId,
		(*data).BanchName,
		(*data).BanchShiftStyle,
		(*data).CreateTime,
		(*data).LastModify,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectCompanyBanchSingle(1, id)
			(*dbObj).Redis.InsertCompanyBanch(res)
		}()
		return true, id
	}
	return false, id
}
func(dbObj *DB) InsertShift(data *table.ShiftTable) (bool, int64) {
	isOk, id := (*dbObj).Mysql.InsertShift(
		(*data).UserId,
		(*data).OnShiftTime,
		(*data).OffShiftTime,
		(*data).PunchIn,
		(*data).PunchOut,
		(*data).CreateTime,
		(*data).LastModify,
		(*data).SpecifyTag,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShiftSingle(0, id)
			(*dbObj).Redis.InsertShift(res)
		}()
		return true, id
	}
	return false, id
}
func(dbObj *DB) InsertShiftChange(data *table.ShiftChangeTable) (bool, int64) {
	isOk, id := (*dbObj).Mysql.InsertShiftChange(
		(*data).InitiatorShiftId,
		(*data).RequestedShiftId,
		(*data).Reason,
		(*data).CaseProcess,
		(*data).SpecifyTag,
		(*data).CreateTime,
		(*data).LastModify,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShiftChangeSingle(0, id)
			(*dbObj).Redis.InsertShiftChange(res)
		}()
		return true, id
	}
	return false, id
}
func(dbObj *DB) InsertShiftOverTime(data *table.ShiftOverTimeTable) (bool, int64) {
	isOk, id := (dbObj).Mysql.InsertShiftOverTime(
		(*data).ShiftId,
		(*data).InitiatorOnOverTime,
		(*data).InitiatorOffOverTime,
		(*data).Reason,
		(*data).CaseProcess,
		(*data).SpecifyTag,
		(*data).CreateTime,
		(*data).LastModify,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectShiftOverTimeSingle(0, id)
			(*dbObj).Redis.InsertShiftOverTime(res)
		}()
		return true, id
	}
	return false, id
}
func(dbObj *DB) InsertDayOff(data *table.DayOffTable) (bool, int64) {
	isOk, id := (*dbObj).Mysql.InsertDayOff(
		(*data).ShiftId,
		(*data).DayOffType,
		(*data).Reason,
		(*data).CaseProcess,
		(*data).SpecifyTag,
		(*data).CreateTime,
		(*data).LastModify,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectDayOffSingle(0, id)
			(*dbObj).Redis.InsertDayOff(res)
		}()
		return true, id
	}
	return false, id
}
func(dbObj *DB) InsertForgetPunch(data *table.ForgetPunchTable) (bool, int64) {
	isOk, id := (*dbObj).Mysql.InsertForgetPunch(
		(*data).ShiftId,
		(*data).TargetPunch,
		(*data).Reason,
		(*data).CaseProcess,
		(*data).SpecifyTag,
		(*data).CreateTime,
		(*data).LastModify,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectForgetPunchSingle(0, id)
			(*dbObj).Redis.InsertForgetPunch(res)
		}()
		return true, id
	}
	return false, id
}
func(dbObj *DB) InsertLateExcusedAll(data *table.LateExcusedTable) (bool, int64) {
	isOk, id := (*dbObj).Mysql.InsertLateExcused(
		(*data).ShiftId,
		(*data).LateExcusedType,
		(*data).Reason,
		(*data).CaseProcess,
		(*data).SpecifyTag,
		(*data).CreateTime,
		(*data).LastModify,
	)
	if isOk {
		go func ()  {
			res := (*dbObj).Mysql.SelectLateExcusedSingle(0, id)
			(*dbObj).Redis.InsertLateExcused(res)
		}()
		return true, id
	}
	return false, id
}


//  ------------------------------update------------------------------


func(dbObj *DB) UpdateUser(){
	
}
func(dbObj *DB) UpdateUserPreference(){
	
}
func(dbObj *DB) UpdateCompany(){
	
}
func(dbObj *DB) UpdateCompanyBanch(){
	
}
func(dbObj *DB) UpdateShift(){
	
}
func(dbObj *DB) UpdateShiftChange(){
	
}
func(dbObj *DB) UpdateShiftOverTime(){
	
}
func(dbObj *DB) UpdateDayOff(){
	
}
func(dbObj *DB) UpdateForgetPunch(){
	
}
func(dbObj *DB) UpdateLateExcused(){
	
}

//  ------------------------------delete------------------------------


func(dbObj *DB) DeleteUser(deleteKey int, userId int64){
	(*dbObj).Mysql.DeleteUser(deleteKey, userId)
	(*dbObj).Redis.DeleteUser(deleteKey, userId)

}
func(dbObj *DB) DeleteUserPreference(deleteKey int, userId int64) {
	(*dbObj).Mysql.DeleteUserPreference(deleteKey, userId)
	(*dbObj).Redis.DeleteUserPreference(deleteKey, userId)
}
func(dbObj *DB) DeleteCompany(deleteKey int, companyId int64){
	(*dbObj).Mysql.DeleteCompany(deleteKey, companyId)
	(*dbObj).Redis.DeleteCompany(deleteKey, companyId)
}
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id int64){
	(*dbObj).Mysql.DeleteCompanyBanch(deleteKey, id)
	(*dbObj).Redis.DeleteCompanyBanch(deleteKey, id)
}
func(dbObj *DB) DeleteShift(deleteKey int, shiftId int64){
	(*dbObj).Mysql.DeleteShift(deleteKey, shiftId)
	(*dbObj).Redis.DeleteShift(deleteKey, shiftId)
}
func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId int64){
	(*dbObj).Mysql.DeleteShiftChange(deleteKey, caseId)
	(*dbObj).Redis.DeleteShiftChange(deleteKey, caseId)
}
func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId int64){
	(*dbObj).Mysql.DeleteShiftOverTime(deleteKey, caseId)
	(*dbObj).Redis.DeleteShiftOverTime(deleteKey, caseId)
}
func(dbObj *DB) DeleteDayOff(deleteKey int, caseId int64){
	(*dbObj).Mysql.DeleteDayOff(deleteKey, caseId)
	(*dbObj).Redis.DeleteDayOff(deleteKey, caseId)
}
func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId int64){
	(*dbObj).Mysql.DeleteForgetPunch(deleteKey, caseId)
	(*dbObj).Redis.DeleteForgetPunch(deleteKey, caseId)
}
func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId int64){
	(*dbObj).Mysql.DeleteLateExcused(deleteKey, caseId)
	(*dbObj).Redis.DeleteLateExcused(deleteKey, caseId)	
}




//  ------------------------------select------------------------------

//0 => all
func(dbObj *DB) SelectUser(selectKey int, value... interface{}) *[]table.UserTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.UserTable {
				return (*dbObj.Redis).SelectUser(selectKey)
			},
			func() *[]table.UserTable {
				return (*dbObj.Mysql).SelectUserAll(selectKey)
			},
			(*dbObj).userLock,
		)
	default:
		return &[]table.UserTable{}
	}
}

//0 => all
func(dbObj *DB) SelectUserPreference(selectKey int, value... interface{}) *[]table.UserPreferenceTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.UserPreferenceTable {
				return (*dbObj.Redis).SelectUserPreference(selectKey)
			},
			func() *[]table.UserPreferenceTable {
				return (*dbObj.Mysql).SelectUserPreferenceAll(selectKey)
			},
			(*dbObj).userPreferenceLock,
		)
	default:
		return &[]table.UserPreferenceTable{}
	}
}

//0 => all
func(dbObj *DB) SelectCompany(selectKey int, value... interface{}) *[]table.CompanyTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.CompanyTable {
				return (*dbObj.Redis).SelectCompany(selectKey)
			},
			func() *[]table.CompanyTable {
				return (*dbObj.Mysql).SelectCompanyAll(selectKey)
			},
			(*dbObj).compnayLock,
		)
	default:
		return &[]table.CompanyTable{}
	}
}

//0 => all
func(dbObj *DB) SelectCompanyBanch(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.CompanyBanchTable {
				return (*dbObj.Redis).SelectCompanyBanch(selectKey)
			},
			func() *[]table.CompanyBanchTable {
				return (*dbObj.Mysql).SelectCompanyBanchAll(selectKey)
			},
			(*dbObj).compnayBanchLock,
		)
	default:
		return &[]table.CompanyBanchTable{}
	}
}

//0 => all
func(dbObj *DB) SelectShift(selectKey int, value... interface{}) *[]table.ShiftTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.ShiftTable {
				return (*dbObj.Redis).SelectShift(selectKey)
			},
			func() *[]table.ShiftTable {
				return (*dbObj.Mysql).SelectShiftAll(selectKey)
			},
			(*dbObj).shiftLock,
		)
	default:
		return &[]table.ShiftTable{}
	}
}

//0 => all
func(dbObj *DB) SelectShiftChange(selectKey int, value... interface{}) *[]table.ShiftChangeTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.ShiftChangeTable {
				return (*dbObj.Redis).SelectShiftChange(selectKey)
			},
			func() *[]table.ShiftChangeTable {
				return (*dbObj.Mysql).SelectShiftChangeAll(selectKey)
			},
			(*dbObj).shiftChangeLock,
		)
	default:
		return &[]table.ShiftChangeTable{}
	}
}

//0 => all
func(dbObj *DB) SelectShiftOverTime(selectKey int, value... interface{}) *[]table.ShiftOverTimeTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.ShiftOverTimeTable {
				return (*dbObj.Redis).SelectShiftOverTime(selectKey)
			},
			func() *[]table.ShiftOverTimeTable {
				return (*dbObj.Mysql).SelectShiftOverTimeAll(selectKey)
			},
			(*dbObj).shiftOverTimeLock,
		)
	default:
		return &[]table.ShiftOverTimeTable{}
	}
}

//0 => all
func(dbObj *DB) SelectDayOff(selectKey int, value... interface{}) *[]table.DayOffTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.DayOffTable {
				return (*dbObj.Redis).SelectDayOff(selectKey)
			},
			func() *[]table.DayOffTable {
				return (*dbObj.Mysql).SelectDayOffAll(selectKey)
			},
			(*dbObj).dayOffLock,
		)
	default:
		return &[]table.DayOffTable{}
	}
}

//0 => all
func(dbObj *DB) SelectForgetPunch(selectKey int, value... interface{}) *[]table.ForgetPunchTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.ForgetPunchTable {
				return (*dbObj.Redis).SelectForgetPunch(selectKey)
			},
			func() *[]table.ForgetPunchTable {
				return (*dbObj.Mysql).SelectForgetPunchAll(selectKey)
			},
			(*dbObj).forgetPunchLock,
		)
	default:
		return &[]table.ForgetPunchTable{}
	}
}

//0 => all
func(dbObj *DB) SelectLateExcused(selectKey int, value... interface{}) *[]table.LateExcusedTable {
	switch selectKey {
	case 0:
		return selectAllHandler(
			func() *[]table.LateExcusedTable {
				return (*dbObj.Redis).SelectLateExcused(selectKey)
			},
			func() *[]table.LateExcusedTable {
				return (*dbObj.Mysql).SelectLateExcusedAll(selectKey)
			},
			(*dbObj).lateExcusedLock,
		)
	default:
		return &[]table.LateExcusedTable{}
	}
}
