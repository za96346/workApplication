package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreShiftOverTimeAll() {
	defer panichandler.Recover()
	(*(*dbObj).shiftOverTimeLock) = true
	(*dbObj).Redis.DeleteKeyShiftOverTime()
	arr := (*dbObj.Mysql).SelectShiftOverTime(0)
	forEach(arr, (*dbObj.Redis).InsertShiftOverTime)
	(*(*dbObj).shiftOverTimeLock) = false
}

func(dbObj *DB) InsertShiftOverTime(data *table.ShiftOverTimeTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (dbObj).Mysql.InsertShiftOverTime(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectShiftOverTime(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertShiftOverTime(&value)
	// 		}
			
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateShiftOverTime(updateKey int, data *table.ShiftOverTimeTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateShiftOverTime(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectShiftOverTime(1, int64((*data).CaseId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertShiftOverTime(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteShiftOverTime(deleteKey, caseId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteShiftOverTime(deleteKey, caseId)
	// 	}()	
	// }
	
	return res
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