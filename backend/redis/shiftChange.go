package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => initiatorShiftId, value => int64
//  3 => requestedShiftId, value => int64
func(dbObj *DB) SelectShiftChange(selectKey int, value... interface{}) *[]table.ShiftChangeTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[5]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftChangeTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.ShiftChangeTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftChangeTable) bool {
				for _, filterItem := range value {
					if filterItem == v.InitiatorShiftId {
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
			func(v table.ShiftChangeTable) bool {
				for _, filterItem := range value {
					if filterItem == v.RequestedShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.ShiftChangeTable{}
	}
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[5], strconv.FormatInt(caseId, 10))
	return true
}

func(dbObj *DB) InsertShiftChange(data *table.ShiftChangeTable) bool {
	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[5], key,jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}
func(dbObj *DB) DeleteKeyShiftChange(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[5])
}