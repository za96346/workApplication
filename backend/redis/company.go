package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"

)

// 0 => 全部, value => nil
//  1 => 公司id, value => int64
//  2 => 公司碼, value => string
func(dbObj *DB) SelectCompany(selectKey int, value... interface{}) *[]table.CompanyTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[2]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.CompanyTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.CompanyTable) bool {return true},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.CompanyTable) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyCode {
						return true
					}
				}
				return false
			},
		)
	default:
		return &[]table.CompanyTable{}
	}
}

//公司的唯一id
func(dbObj *DB) DeleteCompany(deleteKey int, companyId int64) bool {
	defer panichandler.Recover()
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	(*dbObj).RedisDb.HDel((*dbObj).table[2], strconv.FormatInt(companyId, 10))
	return true
}

func(dbObj *DB) InsertCompany(data *table.CompanyTable) bool {
	defer panichandler.Recover()
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	key := strconv.FormatInt((*data).CompanyId, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[2], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}
func(dbObj *DB) DeleteKeyCompany(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[2])
}