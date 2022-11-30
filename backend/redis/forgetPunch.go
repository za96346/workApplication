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
func(dbObj *DB) SelectForgetPunch(selectKey int, value... interface{}) *[]table.ForgetPunchTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[8]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ForgetPunchTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.ForgetPunchTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ForgetPunchTable) bool {
				for _, filterItem := range value {
					if filterItem == v.ShiftId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.ForgetPunchTable{}
	}
}

// 案件的唯一id
func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId int64) bool {
	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[8], strconv.FormatInt(caseId, 10))
	return true
}

func(dbObj *DB) InsertForgetPunch(data *table.ForgetPunchTable) bool {
	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	key := strconv.FormatInt((*data).CaseId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[8], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyForgetPunch(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[8])
}