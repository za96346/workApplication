package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => 全部, value => nil
//  1 => 使用者id, value => int64
func(dbObj *DB) SelectUserPreference(selectKey int, value... interface{}) *[]table.UserPreferenceTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[1]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(V table.UserPreferenceTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(V table.UserPreferenceTable) bool {return true},
		)
	default:
		return &[]table.UserPreferenceTable{}
	}
}

//使用者的唯一id
func(dbObj *DB) DeleteUserPreference(deleteKey int, userId int64) bool {
	defer panichandler.Recover()
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[1], strconv.FormatInt(userId, 10))
	return true
}

func(dbObj *DB) InsertUserPreference(data *table.UserPreferenceTable) bool {
	defer panichandler.Recover()
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	key := strconv.FormatInt((*data).UserId, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[1], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyUserPreference(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[1])
}