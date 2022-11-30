package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => all, value => nil
//  1 => quitId, value => int64
//   2 => userId, value => int64
//   3 => companyCode, value => string 
//   4=> companyCode && userId ,  value string && int64
func(dbObj *DB) SelectQuitWorkUser(selectKey int, value... interface{}) *[]table.QuitWorkUser {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[12]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.QuitWorkUser) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.QuitWorkUser) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.QuitWorkUser) bool {
				for _, filterItem := range value {
					if filterItem == v.UserId {
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
			func(v table.QuitWorkUser) bool {
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
			func(v table.QuitWorkUser) bool {
				if v.CompanyCode == value[0] && v.UserId == value[1] {
					return true
				}
				return false
			},
		)
	default:
		return &[]table.QuitWorkUser{}
	}
}

// quit work suer 的唯一id
func(dbObj *DB) DeleteQuitWorkUser(deleteKey int, quitId int64) bool {
	defer panichandler.Recover()
	(*dbObj).quitWorkUserMux.Lock()
	defer (*dbObj).quitWorkUserMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[12], strconv.FormatInt(quitId, 10))
	return true
}

func(dbObj *DB) InsertQuitWorkUser(data *table.QuitWorkUser) bool {
	defer panichandler.Recover()
	(*dbObj).quitWorkUserMux.Lock()
	defer (*dbObj).quitWorkUserMux.Unlock()
	key := strconv.FormatInt((*data).QuitId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[12], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}
func(dbObj *DB) DeleteKeyQuitWorkUser(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[12])
}