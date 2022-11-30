package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (

	// "runtime"
	"backend/query"
	"backend/table"

	"backend/panicHandler"

	_ "github.com/go-sql-driver/mysql"
)

// 0 => 全部, value => nil
//	1 => 公司Id, value => int64
// 	2 => id (banchId), value => int64
func(dbObj *DB) SelectCompanyBanch(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectSingleByCompanyId
		break
	case 2:
		// value need int
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectSingleById
		break;
	default:
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectAll
		break
	}
	companyBanch := new(table.CompanyBanchTable)
	carry := []table.CompanyBanchTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&companyBanch.Id,
			&companyBanch.CompanyId,
			&companyBanch.BanchName,
			&companyBanch.BanchShiftStyle,
			&companyBanch.CreateTime,
			&companyBanch.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *companyBanch)
		}
	}

	return &carry 
}

// 公司部門的id
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).CompanyBanch.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(id)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateCompanyBanch(updateKey int, data *table.CompanyBanchTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	defer func ()  {
		(*dbObj).containers.companyBanch = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).CompanyBanch.UpdateSingle
		(*dbObj).containers.companyBanch = append(
			(*dbObj).containers.companyBanch,
			(*data).BanchName,
			(*data).BanchShiftStyle,
			(*data).LastModify,
			(*data).Id,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.companyBanch...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertCompanyBanch(data *table.CompanyBanchTable) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).CompanyBanch.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).CompanyId,
		(*data).BanchName,
		(*data).BanchShiftStyle,
		(*data).CreateTime,
		(*data).LastModify,
	)
	(*dbObj).checkErr(err)
	id, _:= res.LastInsertId()
	if err != nil {
		return false, id
	}
	return true, id
}