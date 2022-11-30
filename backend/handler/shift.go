package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreShiftAll() {
	defer panichandler.Recover()
	(*(*dbObj).shiftLock) = true
	(*dbObj).Redis.DeleteKeyShift()
	arr := (*dbObj.Mysql).SelectShift(0)
	forEach(arr, (*dbObj.Redis).InsertShift)
	(*(*dbObj).shiftLock) = false
}

func(dbObj *DB) InsertShift(data *table.ShiftTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertShift(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectShift(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertShift(&value)
	// 		}
			
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateShift(updateKey int, data *table.ShiftTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateShift(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectShift(1, int64((*data).ShiftId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertShift(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteShift(deleteKey int, shiftId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteShift(deleteKey, shiftId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteShift(deleteKey, shiftId)
	// 	}()	
	// }
	return res
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