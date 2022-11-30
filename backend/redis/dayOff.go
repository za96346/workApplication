package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
func(dbObj *DB) SelectDayOff(selectKey int, value... interface{}) *[]table.DayOffTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[7]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.DayOffTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.DayOffTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.DayOffTable) bool {
				for _, filterItem := range value {
					if filterItem == v.ShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.DayOffTable{}
	}
}

// 案件的唯一id
func(dbObj *DB) DeleteDayOff(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[7], strconv.FormatInt(caseId, 10))
	return true
}

func(dbObj *DB) InsertDayOff(data *table.DayOffTable) bool {
	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[7], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyDayOff(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[7])
}