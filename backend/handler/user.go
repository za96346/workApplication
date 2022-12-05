package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreUserAll() {
	defer panichandler.Recover()
	(*(*dbObj).userLock) = true
	(*dbObj).Redis.DeleteKeyUser()
	arr := (*dbObj.Mysql).SelectUser(0)
	forEach(arr, (*dbObj.Redis).InsertUser)
	(*(*dbObj).userLock) = false
}

func(dbObj *DB) InsertUser(data *table.UserTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertUser(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectUser(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertUser(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateUser(updateKey int, data *table.UserTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateUser(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectUser(1, int64((*data).UserId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertUser(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

// func(dbObj *DB) DeleteUser(deleteKey int, userId int64) bool {
// 	defer panichandler.Recover()
// 	res := (*dbObj).Mysql.DeleteUser(deleteKey, userId)
// 	if res {
// 		go func ()  {
// 			(*dbObj).TakeAllFromMysql()
// 		}()
// 	}
// 	return res
// }



// 0 => 全部, value => nil
//  1 =>  userId, value => int64
//  2 => account, value => string
// 3 => companyCode, value => string
//  4 => banch, value = > int64
// . 5 => companyCode, userId, value => string, int64
func(dbObj *DB) SelectUser(selectKey int, value... interface{}) *[]table.UserTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.UserTable {
			return (*dbObj.Redis).SelectUser(selectKey, value...)
		},
		func() *[]table.UserTable {
			return (*dbObj.Mysql).SelectUser(selectKey, value...)
		},
		(*dbObj).userLock,
	)
}