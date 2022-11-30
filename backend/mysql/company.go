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
//  1 => 公司id, value => int64
//  2 => 公司碼, value => string
func(dbObj *DB) SelectCompany(selectKey int, value... interface{}) *[]table.CompanyTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).Company.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).Company.SelectSingleByCompanyId
		break
	case 2:
		// value need string
		querys = (*query.MysqlSingleton()).Company.SelectSingleByCompanyCode
		break
	default:
		querys = (*query.MysqlSingleton()).Company.SelectAll
		break
	}
	company := new(table.CompanyTable)
	carry := []table.CompanyTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&company.CompanyId,
			&company.CompanyCode,
			&company.CompanyName,
			&company.CompanyLocation,
			&company.CompanyPhoneNumber,
			&company.BossId,
			&company.SettlementDate,
			&company.TermStart,
			&company.TermEnd,
			&company.CreateTime,
			&company.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *company)
		}
	}

	return &carry
}

//公司的唯一id (關聯資料表 companyBanch 也上鎖)
func(dbObj *DB) DeleteCompany(deleteKey int, companyId interface{}) bool {
	defer panichandler.Recover()

	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()

	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()

	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Company.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(companyId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateCompany(updateKey int, data *table.CompanyTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	defer func ()  {
		(*dbObj).containers.company = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).Company.UpdateSingle
		(*dbObj).containers.company = append(
			(*dbObj).containers.company,
			(*data).CompanyName,
			(*data).CompanyLocation,
			(*data).CompanyPhoneNumber,
			(*data).BossId,
			(*data).SettlementDate,
			(*data).TermStart,
			(*data).TermEnd,
			(*data).LastModify,
			(*data).CompanyId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.company...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertCompany(data *table.CompanyTable) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Company.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)

	res, err := stmt.Exec(
		(*data).CompanyCode,
		(*data).CompanyName,
		(*data).CompanyLocation,
		(*data).CompanyPhoneNumber,
		(*data).BossId,
		(*data).SettlementDate,
		(*data).TermStart,
		(*data).TermEnd,
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