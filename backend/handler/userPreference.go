package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreUserPreferenceAll() {
	defer panichandler.Recover()
	(*(*dbObj).userPreferenceLock) = true
	(*dbObj).Redis.DeleteKeyUserPreference()
	arr := (*dbObj.Mysql).SelectUserPreference(0)
	forEach(arr, (*dbObj.Redis).InsertUserPreference)
	(*(*dbObj).userPreferenceLock) = false
}

func(dbObj *DB) InsertUserPreference(data *table.UserPreferenceTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, _ := (*dbObj).Mysql.InsertUserPreference(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectUserPreference(1, (*data).UserId)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertUserPreference(&value)
	// 		}
	// 	}()
	// }
	return isOk, (*data).UserId
}

func(dbObj *DB) UpdateUserPreference(updateKey int, data *table.UserPreferenceTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateUserPreference(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectUserPreference(1, int64((*data).UserId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertUserPreference(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteUserPreference(deleteKey int, userId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteUserPreference(deleteKey, userId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteUserPreference(deleteKey, userId)
	// 	}()	
	// }
	return res
}

// 0 => 全部, value => nil
//  1 => 使用者id, value => int64
func(dbObj *DB) SelectUserPreference(selectKey int, value... interface{}) *[]table.UserPreferenceTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.UserPreferenceTable {
			return (*dbObj.Redis).SelectUserPreference(selectKey, value...)
		},
		func() *[]table.UserPreferenceTable {
			return (*dbObj.Mysql).SelectUserPreference(selectKey, value...)
		},
		(*dbObj).userPreferenceLock,
	)
}