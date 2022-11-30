package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => all, value => nil
//  1 => weekendId, value => int64
//  2 => companyId, value => int64
func(dbObj *DB) SelectWeekendSetting (selectKey int, value... interface{}) *[]table.WeekendSetting {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[14]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WeekendSetting) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.WeekendSetting) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WeekendSetting) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.WeekendSetting{}
	}
}

// weekend setting 的唯一id
func(dbObj *DB) DeleteWeekendSetting(deleteKey int, weekendId int64) bool {
	defer panichandler.Recover()
	(*dbObj).weekendSetting.Lock()
	defer (*dbObj).weekendSetting.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[14], strconv.FormatInt(weekendId, 10))
	return true
}

func(dbObj *DB) InsertWeekendSetting(data *table.WeekendSetting) bool {
	defer panichandler.Recover()
	(*dbObj).weekendSetting.Lock()
	defer (*dbObj).weekendSetting.Unlock()
	key := strconv.FormatInt((*data).WeekendId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[14], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyWeekendSetting(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[14])
}
