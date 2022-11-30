package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => all, value => nil
//  1 => waitId, value => int64
//  2 => userId, value => int64
//  3 => companyId, value => int64
//  4 => comapnyId && userId, value => int64, int64
func(dbObj *DB) SelectWaitCompanyReply (selectKey int, value... interface{}) *[]table.WaitCompanyReply {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[13]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WaitCompanyReply) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.WaitCompanyReply) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.WaitCompanyReply) bool {
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
			func(v table.WaitCompanyReply) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyId {
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
			func(v table.WaitCompanyReply) bool {
				if value[0] == v.CompanyId && value[1] == v.UserId {
					return true
				}
				return false
			},
		)
	default:
		return &[]table.WaitCompanyReply{}
	}
}

// wait company reply 的唯一id
func(dbObj *DB) DeleteWaitCompanyReply(deleteKey int, waitId int64) bool {
	defer panichandler.Recover()
	(*dbObj).waitCompanyReply.Lock()
	defer (*dbObj).waitCompanyReply.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[13], strconv.FormatInt(waitId, 10))
	return true
}

func(dbObj *DB) InsertWaitCompanyReply(data *table.WaitCompanyReply) bool {
	defer panichandler.Recover()
	(*dbObj).waitCompanyReply.Lock()
	defer (*dbObj).waitCompanyReply.Unlock()
	key := strconv.FormatInt((*data).WaitId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[13], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

func(dbObj *DB) DeleteKeyWaitCompanyReply(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[13])
}