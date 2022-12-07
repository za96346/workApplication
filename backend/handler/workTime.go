package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) InsertWorkTime (data *table.WorkTime) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertWorkTime(data)
	return isOk, id
}

//  0 => all, by workTimeId int64, companyCode string
func(dbObj *DB) UpdateWorkTime (updateKey int, data *table.WorkTime, value... interface{}) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateWorkTime(updateKey, data, value...)
	return isOk
}


//   workTime 的唯一id
//   workTimeId && companyCode, int64 && string
func(dbObj *DB) DeleteWorkTime (deleteKey int, value... interface{}) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteWorkTime(deleteKey, value...)
	return res
}

//0 => companyCode, value => string
//  . 1 => userId, companyCode, value => int64, string
//  . 2 => year && month && companyCode, value => int && int && string
//  . 3 => year && month && userId && companyCode, value => int, int, int64 && string
func(dbObj *DB) SelectWorkTime (selectKey int, value... interface{}) *[]table.WorkTime {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.WorkTime {
			return &[]table.WorkTime{}
		},
		func() *[]table.WorkTime {
			return (*dbObj.Mysql).SelectWorkTime(selectKey, value...)
		},
		(*dbObj).workTime,
	)
}