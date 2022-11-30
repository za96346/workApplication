package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"
	"sync"

)

// 0 => 全部, value => nil
//	1 => 公司Id, value => int64
// 	2 => id (banchId), value => int64
func(dbObj *DB) SelectCompanyBanch(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[3]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.CompanyBanchTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.CompanyBanchTable) bool {
				for _, filterItem := range value {
					if filterItem == v.CompanyId {
						return true
					}
				}
				return false
			},
		)
	case 2:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.CompanyBanchTable) bool {return true},
		)
	default:
		return &[]table.CompanyBanchTable{}
	}
}

// 公司部門的id
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id int64) bool {
	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()

	res := (*dbObj).SelectCompanyBanch(2, id)
	(*dbObj).RedisDb.HDel((*dbObj).table[3], strconv.FormatInt(id, 10))
	if len(*res) <= 0 {
		return true
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func ()  {
		defer wg.Done()
		banchStyle := (*dbObj).SelectBanchStyle(2, (*res)[0].Id)
		for _, v := range *banchStyle {
			(*dbObj).DeleteBanchStyle(0, v.StyleId)
		}
		
	}()
	go func ()  {
		defer wg.Done()
		banchRule := (*dbObj).SelectBanchRule(2, (*res)[0].Id)
		for _, v:= range *banchRule {
			(*dbObj).DeleteBanchRule(0, v.RuleId)
		}
	}()
	wg.Wait()
	return true
}
func(dbObj *DB) InsertCompanyBanch(data *table.CompanyBanchTable) bool {
	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	key := strconv.FormatInt((*data).Id, 10)
	jsonData, _ := json.Marshal(data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[3], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}
func(dbObj *DB) DeleteKeyCompanyBanch(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[3])
}