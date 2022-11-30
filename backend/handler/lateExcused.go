package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreLateExcusedAll() {
	defer panichandler.Recover()
	(*(*dbObj).lateExcusedLock) = true
	(*dbObj).Redis.DeleteKeyLateExcused()
	arr := (*dbObj.Mysql).SelectLateExcused(0)
	forEach(arr, (*dbObj.Redis).InsertLateExcused)
	(*(*dbObj).lateExcusedLock) = false
}

func(dbObj *DB) InsertLateExcused(data *table.LateExcusedTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertLateExcused(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectLateExcused(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertLateExcused(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateLateExcused(updateKey int, data *table.LateExcusedTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateLateExcused(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectLateExcused(1, int64((*data).CaseId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertLateExcused(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteLateExcused(deleteKey, caseId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteLateExcused(deleteKey, caseId)
	// 	}()	
	// }
	
	return res
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