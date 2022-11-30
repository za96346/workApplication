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
func(dbObj *DB) SelectLateExcused(selectKey int, value... interface{}) *[]table.LateExcusedTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[9]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.LateExcusedTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.LateExcusedTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.LateExcusedTable) bool {
				for _, filterItem := range value {
					if filterItem == v.ShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.LateExcusedTable{}
	}
}

// 案件的唯一id
func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[9], strconv.FormatInt(caseId, 10))
	return true
}

func(dbObj *DB) InsertLateExcused(data *table.LateExcusedTable) bool {
	defer panichandler.Recover()
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[9], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyLateExcused(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[9])
}