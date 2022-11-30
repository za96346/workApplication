package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreBanchRuleAll() {
	defer panichandler.Recover()
	(*(*dbObj).banchRuleLock) = true
	(*dbObj).Redis.DeleteKeyBanchRule()
	arr := (*dbObj.Mysql).SelectBanchRule(0)
	forEach(arr, (*dbObj.Redis).InsertBanchRule)
	(*(*dbObj).banchRuleLock) = false
}

func(dbObj *DB) InsertBanchRule(data *table.BanchRule) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertBanchRule(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectBanchRule(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertBanchRule(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateBanchRule(updateKey int, data *table.BanchRule) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateBanchRule(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectBanchRule(1, int64((*data).RuleId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertBanchRule(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteBanchRule(deleteKey int, ruleId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteBanchRule(deleteKey, ruleId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteBanchRule(deleteKey, ruleId)
	// 	}()	
	// }
	
	return res
}

// 0 => all, value => nil
//  1 => ruleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchRule(selectKey int, value... interface{}) *[]table.BanchRule {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.BanchRule {
			return (*dbObj.Redis).SelectBanchRule(selectKey, value...)
		},
		func() *[]table.BanchRule {
			return (*dbObj.Mysql).SelectBanchRule(selectKey, value...)
		},
		(*dbObj).banchRuleLock,
	)
}