package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreForgetPunchAll() {
	defer panichandler.Recover()
	(*(*dbObj).forgetPunchLock) = true
	(*dbObj).Redis.DeleteKeyForgetPunch()
	arr := (*dbObj.Mysql).SelectForgetPunch(0)
	forEach(arr, (*dbObj.Redis).InsertForgetPunch)
	(*(*dbObj).forgetPunchLock)= false
}

func(dbObj *DB) InsertForgetPunch(data *table.ForgetPunchTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertForgetPunch(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectForgetPunch(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertForgetPunch(&value)
	// 		}
			
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateForgetPunch(updateKey int, data *table.ForgetPunchTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateForgetPunch(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectForgetPunch(1, int64((*data).CaseId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertForgetPunch(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteForgetPunch(deleteKey, caseId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteForgetPunch(deleteKey, caseId)
	// 	}()	
	// }
	
	return res
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