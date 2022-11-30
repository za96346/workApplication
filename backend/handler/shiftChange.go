package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreShiftChangeAll() {
	defer panichandler.Recover()
	(*(*dbObj).shiftChangeLock) = true
	(*dbObj).Redis.DeleteKeyShiftChange()
	arr := (*dbObj.Mysql).SelectShiftChange(0)
	forEach(arr, (*dbObj.Redis).InsertShiftChange)
	(*(*dbObj).shiftChangeLock) = false
}

func(dbObj *DB) InsertShiftChange(data *table.ShiftChangeTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertShiftChange(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectShiftChange(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertShiftChange(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateShiftChange(updateKey int, data *table.ShiftChangeTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateShiftChange(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectShiftChange(1, int64((*data).CaseId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertShiftChange(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteShiftChange(deleteKey, caseId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteShiftChange(deleteKey, caseId)	
	// 	}()	
	// }
	return res
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