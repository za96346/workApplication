package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => all, value => nil
//  1 => styleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchStyle(selectKey int, value... interface{}) *[]table.BanchStyle {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[10]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.BanchStyle) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.BanchStyle) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.BanchStyle) bool {
				for _, filterItem := range value {
					if filterItem == v.BanchId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.BanchStyle{}
	}
}

// style的唯一id
func(dbObj *DB) DeleteBanchStyle(deleteKey int, styleId int64) bool {
	defer panichandler.Recover()
	(*dbObj).banchStyleMux.Lock()
	defer (*dbObj).banchStyleMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[10], strconv.FormatInt(styleId, 10))
	return true
}

func(dbObj *DB) InsertBanchStyle(data *table.BanchStyle) bool {
	defer panichandler.Recover()
	(*dbObj).banchStyleMux.Lock()
	defer (*dbObj).banchStyleMux.Unlock()
	key := strconv.FormatInt((*data).StyleId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[10], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyBanchStyle(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[10])
}