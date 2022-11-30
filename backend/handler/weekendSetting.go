package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreWeekendSettingAll() {
	defer panichandler.Recover()
	(*(*dbObj).weekendSettingLock) = true
	(*dbObj).Redis.DeleteKeyWeekendSetting()
	arr := (*dbObj.Mysql).SelectWeekendSetting(0)
	forEach(arr, (*dbObj.Redis).InsertWeekendSetting)
	(*(*dbObj).weekendSettingLock) = false
}

func(dbObj *DB) InsertWeekendSetting(data *table.WeekendSetting) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertWeekendSetting(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectWeekendSetting(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertWeekendSetting(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateWeekendSetting(updateKey int, data *table.WeekendSetting) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateWeekendSetting(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectWeekendSetting(1, int64((*data).WeekendId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertWeekendSetting(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteWeekendSetting(deleteKey int, weekendId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteWeekendSetting(deleteKey, weekendId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteWeekendSetting(deleteKey, weekendId)
	// 	}()	
	// }
	
	return res
}

// 0 => all, value => nil
//  1 => weekendId, value => int64
//  2 => companyId, value => int64
func(dbObj *DB) SelectWeekendSetting(selectKey int, value... interface{}) *[]table.WeekendSetting {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.WeekendSetting {
			return (*dbObj.Redis).SelectWeekendSetting(selectKey, value...)
		},
		func() *[]table.WeekendSetting {
			return (*dbObj.Mysql).SelectWeekendSetting(selectKey, value...)
		},
		(*dbObj).weekendSettingLock,
	)
}