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

func(dbObj *DB) UpdateWorkTime (updateKey int, data *table.WorkTime) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateWorkTime(updateKey, data)
	return isOk
}

func(dbObj *DB) DeleteWorkTime (deleteKey int, workTimeId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteWorkTime(deleteKey, workTimeId)
	return res
}

// . 0 => all, value => nil
// . 1 => userId, value => int64
// . 2 => year && month, value => int && int
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