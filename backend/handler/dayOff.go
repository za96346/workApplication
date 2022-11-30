package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreDayOffAll() {
	defer panichandler.Recover()
	(*(*dbObj).dayOffLock) = true
	(*dbObj).Redis.DeleteKeyDayOff()
	arr := (*dbObj.Mysql).SelectDayOff(0)
	forEach(arr, (*dbObj.Redis).InsertDayOff)
	(*(*dbObj).dayOffLock) = false
}

func(dbObj *DB) InsertDayOff(data *table.DayOffTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertDayOff(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectDayOff(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertDayOff(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateDayOff(updateKey int, data *table.DayOffTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateDayOff(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectDayOff(1, int64((*data).CaseId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertDayOff(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteDayOff(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteDayOff(deleteKey, caseId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteDayOff(deleteKey, caseId)
	// 	}()	
	// }
	
	return res
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