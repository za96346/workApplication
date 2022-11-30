package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => 全部, value =>  nil
//  1 =>  userId, value => int64
//  2 => account, value => string
// 3 => companyCode, value => string
//  4 => banch, value = > int64
func(dbObj *DB) SelectUser(selectKey int, value... interface{}) *[]table.UserTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[0]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.UserTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.UserTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.UserTable) bool {
				for _, filterItem := range value {
					if filterItem == v.Account {
						return true
					}
				}
				return false
			},
		)
	case 3:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.UserTable) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyCode {
						return true
					}
				}
				return false
			},
		)
	case 4:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.UserTable) bool {
				for _, filterItem := range value {
					if filterItem == v.Banch {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.UserTable{}
	}
}

//使用者的唯一id
func(dbObj *DB) DeleteUser(deleteKey int, userId int64) bool {
	defer panichandler.Recover()
	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[0], strconv.FormatInt(userId, 10))
	return true
}
func(dbObj *DB) InsertUser(data *table.UserTable) bool {
	defer panichandler.Recover()
	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()
	key := strconv.FormatInt((*data).UserId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[0], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyUser(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[0])
}