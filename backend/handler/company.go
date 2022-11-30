package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreCompanyAll() {
	defer panichandler.Recover()
	(*(*dbObj).compnayLock)= true
	(*dbObj).Redis.DeleteKeyCompany()
	arr := (*dbObj.Mysql).SelectCompany(0)
	forEach(arr, (*dbObj.Redis).InsertCompany)
	(*(*dbObj).compnayLock) = false
}

func(dbObj *DB) InsertCompany(data *table.CompanyTable) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertCompany(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectCompany(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertCompany(&value)	
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateCompany(updateKey int, data *table.CompanyTable) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateCompany(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectCompany(1, int64((*data).CompanyId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertCompany(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

// func(dbObj *DB) DeleteCompany(deleteKey int, companyId int64) bool {
// 	defer panichandler.Recover()
// 	res := (*dbObj).Mysql.DeleteCompany(deleteKey, companyId)
// 	if res {
// 		go func ()  {
// 			(*dbObj).TakeAllFromMysql()
// 		}()	
// 	}
// 	return res
// }



// 0 => 全部, value => nil
//  1 => 公司id, value => int64
//  2 => 公司碼, value => string
func(dbObj *DB) SelectCompany(selectKey int, value... interface{}) *[]table.CompanyTable {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.CompanyTable {
			return (*dbObj.Redis).SelectCompany(selectKey, value...)
		},
		func() *[]table.CompanyTable {
			return (*dbObj.Mysql).SelectCompany(selectKey, value...)
		},
		(*dbObj).compnayLock,
	)
}