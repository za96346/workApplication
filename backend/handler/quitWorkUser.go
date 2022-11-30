package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreQuitWorkUserAll() {
	defer panichandler.Recover()
	(*(*dbObj).quitWorkUserLock) = true
	(*dbObj).Redis.DeleteKeyQuitWorkUser()
	arr := (*dbObj.Mysql).SelectQuitWorkUser(0)
	forEach(arr, (*dbObj.Redis).InsertQuitWorkUser)
	(*(*dbObj).quitWorkUserLock) = false
}

func(dbObj *DB) InsertQuitWorkUser(data *table.QuitWorkUser) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertQuitWorkUser(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectQuitWorkUser(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertQuitWorkUser(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateQuitWorkUser(updateKey int, data *table.QuitWorkUser) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateQuitWorkUser(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectQuitWorkUser(1, int64((*data).QuitId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertQuitWorkUser(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteQuitWorkUser(deleteKey int, quitId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteQuitWorkUser(deleteKey, quitId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteQuitWorkUser(deleteKey, quitId)
	// 	}()	
	// }
	
	return res
}

// 0 => all, value => nil
//  1 => quitId, value => int64
//   2 => userId, value => int64
//   3 => companyCode, value => string 
//   4=> companyCode && userId ,  value string && int64
func(dbObj *DB) SelectQuitWorkUser(selectKey int, value... interface{}) *[]table.QuitWorkUser {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.QuitWorkUser {
			return (*dbObj.Redis).SelectQuitWorkUser(selectKey, value...)
		},
		func() *[]table.QuitWorkUser {
			return (*dbObj.Mysql).SelectQuitWorkUser(selectKey, value...)
		},
		(*dbObj).quitWorkUserLock,
	)
}