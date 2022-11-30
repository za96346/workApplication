package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreBanchStyleAll() {
	defer panichandler.Recover()
	(*(*dbObj).banchStyleLock) = true
	(*dbObj).Redis.DeleteKeyBanchStyle()
	arr := (*dbObj.Mysql).SelectBanchStyle(0)
	forEach(arr, (*dbObj.Redis).InsertBanchStyle)
	(*(*dbObj).banchStyleLock) = false
}

func(dbObj *DB) InsertBanchStyle(data *table.BanchStyle) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertBanchStyle(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectBanchStyle(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertBanchStyle(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateBanchStyle(updateKey int, data *table.BanchStyle) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateBanchStyle(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectBanchStyle(1, int64((*data).StyleId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertBanchStyle(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteBanchStyle(deleteKey int, styleId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteBanchStyle(deleteKey, styleId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteBanchStyle(deleteKey, styleId)
	// 	}()	
	// }
	
	return res
}

// 0 => all, value => nil
//  1 => styleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchStyle(selectKey int, value... interface{}) *[]table.BanchStyle {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.BanchStyle {
			return (*dbObj.Redis).SelectBanchStyle(selectKey, value...)
		},
		func() *[]table.BanchStyle {
			return (*dbObj.Mysql).SelectBanchStyle(selectKey, value...)
		},
		(*dbObj).banchStyleLock,
	)
}