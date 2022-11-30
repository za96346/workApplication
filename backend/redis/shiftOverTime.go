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
func(dbObj *DB) SelectShiftOverTime(selectKey int, value... interface{}) *[]table.ShiftOverTimeTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[6]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftOverTimeTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.ShiftOverTimeTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftOverTimeTable) bool {
				for _, filterItem := range value {
					if filterItem == v.ShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.ShiftOverTimeTable{}
	}
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[6], strconv.FormatInt(caseId, 10))
	return true
}

func(dbObj *DB) InsertShiftOverTime(data *table.ShiftOverTimeTable) bool {
	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[6], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}
func(dbObj *DB) DeleteKeyShiftOverTime(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[6])
}