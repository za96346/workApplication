package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => all, value => nil
//  1 => ruleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchRule(selectKey int, value... interface{}) *[]table.BanchRule {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[11]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.BanchRule) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.BanchRule) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.BanchRule) bool {
				for _, filterItem := range value {
					if filterItem == v.BanchId {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.BanchRule{}
	}
}

// style的唯一id
func(dbObj *DB) DeleteBanchRule(deleteKey int, ruleId int64) bool {
	defer panichandler.Recover()
	(*dbObj).banchRuleMux.Lock()
	defer (*dbObj).banchRuleMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[11], strconv.FormatInt(ruleId, 10))
	return true
}


func(dbObj *DB) InsertBanchRule(data *table.BanchRule) bool {
	defer panichandler.Recover()
	(*dbObj).banchRuleMux.Lock()
	defer (*dbObj).banchRuleMux.Unlock()
	key := strconv.FormatInt((*data).RuleId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[11], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyBanchRule(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[11])
}