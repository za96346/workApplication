package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) InsertPaidVocation (data *table.PaidVocation) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertPaidVocation(data)
	return isOk, id
}

func(dbObj *DB) UpdatePaidVocation (updateKey int, data *table.PaidVocation) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdatePaidVocation(updateKey, data)
	return isOk
}

func(dbObj *DB) DeletePaidVocation (deleteKey int, paidVocationId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeletePaidVocation(deleteKey, paidVocationId)
	return res
}

// . 0 => all, value => nil
// . 1 => userId, value => int64
// . 2 => year, value => int
func(dbObj *DB) SelectPaidVocation (selectKey int, value... interface{}) *[]table.PaidVocation {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.PaidVocation {
			return &[]table.PaidVocation{}
		},
		func() *[]table.PaidVocation {
			return (*dbObj.Mysql).SelectPaidVocation(selectKey, value...)
		},
		(*dbObj).paidVocation,
	)
}