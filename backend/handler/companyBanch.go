package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreCompanyBanchAll() {
	defer panichandler.Recover()
	(*(*dbObj).compnayBanchLock)= true
	(*dbObj).Redis.DeleteKeyCompanyBanch()
	arr := (*dbObj.Mysql).SelectCompanyBanch(0)
	forEach(arr, (*dbObj.Redis).InsertCompanyBanch)
	(*(*dbObj).compnayBanchLock) = false
}

func(dbObj *DB) InsertCompanyBanch(data *table.CompanyBanchTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertCompanyBanch(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectCompanyBanch(2, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertCompanyBanch(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateCompanyBanch(updateKey int, data *table.CompanyBanchTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateCompanyBanch(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectCompanyBanch(2, int64((*data).Id))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertCompanyBanch(&v)
	// 		}

	// 	}()
	// }
	return isOk
}
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteCompanyBanch(deleteKey, id)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteCompanyBanch(deleteKey, id)
	// 	}()	
	// }
	return res
}


// 0 => 全部, value => nil
//	1 => 公司Id, value => int64
// 	2 => id (banchId), value => int64
func(dbObj *DB) SelectCompanyBanch(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.CompanyBanchTable {
			return (*dbObj.Redis).SelectCompanyBanch(selectKey, value...)
		},
		func() *[]table.CompanyBanchTable {
			return (*dbObj.Mysql).SelectCompanyBanch(selectKey, value...)
		},
		(*dbObj).compnayBanchLock,
	)
}