package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (
	"database/sql"

	// "runtime"
	"backend/query"
	"backend/table"

	"backend/panicHandler"

	_ "github.com/go-sql-driver/mysql"
)

// 0 => 全部, value => nil
//  1 =>  userId, value => int64
//  2 => account, value => string
// 3 => companyCode, value => string
//  4 => banch, value = > int64
// . 5 => companyCode, userId, value => string, int64
func(dbObj *DB) SelectUser(selectKey int, value... interface{}) *[]table.UserExtend {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).User.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).User.SelectSingleByUserId
		break
	case 2:
		// value need string
		querys = (*query.MysqlSingleton()).User.SelectSingleByAccount
		break
	case 3:
		// value need string
		querys = (*query.MysqlSingleton()).User.SelectAllByCompanyCode
		break
	case 4:
		//value need int64
		querys = (*query.MysqlSingleton()).User.SelectAllByBanchId
		break
	case 5:
		querys = (*query.MysqlSingleton()).User.SelectAllByUserIdAndCompanyCode
		break
	default:
		querys = (*query.MysqlSingleton()).User.SelectAll
		break
	}
	user := new(table.UserExtend)
	carry := []table.UserExtend{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	companyCode := new(sql.NullString)
	for res.Next() {
		err = res.Scan(
			&user.UserId,
			companyCode,
			&user.UserName,
			&user.EmployeeNumber,
			&user.Account,
			&user.Password,
			&user.OnWorkDay,
			&user.Banch,
			&user.Permession,
			&user.CreateTime,
			&user.LastModify,
			&user.MonthSalary,
			&user.PartTimeSalary,
			&user.BanchName,
			&user.CompanyId,
			&user.CompanyName,
		)
		(*dbObj).checkErr(err)
		if companyCode.String == "" {
			user.CompanyCode = ""
		} else {
			user.CompanyCode = companyCode.String
		}
		
		if err == nil {
			carry = append(carry, *user)
		}
	}
	return &carry
}

//使用者的唯一id (關聯資料表userpreference 也上鎖)
func(dbObj *DB) DeleteUser(deleteKey int, userId interface{}) bool {
	defer panichandler.Recover()

	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()

	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()

	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).User.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(userId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
//  1 => CompanyCode, value string
// . 2 => boss
func(dbObj *DB) UpdateUser(updateKey int, data *table.UserTable, value... interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()
	defer func ()  {
		(*dbObj).containers.user = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).User.UpdateSingle
		(*dbObj).containers.user = append(
			(*dbObj).containers.user,
			(*data).EmployeeNumber,
			(*data).CompanyCode,
			(*data).Password,
			(*data).UserName,
			(*data).OnWorkDay,
			(*data).Banch,
			(*data).Permession,
			(*data).LastModify,
			(*data).MonthSalary,
			(*data).PartTimeSalary,
			(*data).UserId,
		)
		break
	case 1:
		querys = (*query.MysqlSingleton()).User.UpdateCompanyUser
		(*dbObj).containers.user = append(
			(*dbObj).containers.user,
			(*data).EmployeeNumber,
			(*data).CompanyCode,
			(*data).OnWorkDay,
			(*data).Banch,
			(*data).Permession,
			(*data).LastModify,
			(*data).UserId,
		)
		break
	case 2:
		querys = (*query.MysqlSingleton()).User.UpdateBoss
		(*dbObj).containers.user = append(
			(*dbObj).containers.user,
			(*data).CompanyCode,
			(*data).Banch,
			(*data).Permession,
			(*data).LastModify,
			(*data).UserId,
		)
		break
	}
	(*dbObj).containers.user = append((*dbObj).containers.user, value...)
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.user...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}
func(dbObj *DB) InsertUser(data *table.UserTable) (bool, int64) {
	///
		defer panichandler.Recover()
		(*dbObj).userMux.Lock()
		defer (*dbObj).userMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).User.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).CompanyCode,
			(*data).Account,
			(*data).Password,
			(*data).UserName,
			(*data).EmployeeNumber,
			(*data).OnWorkDay,
			(*data).Banch,
			(*data).Permession,
			(*data).CreateTime,
			(*data).LastModify,
			(*data).MonthSalary,
			(*data).PartTimeSalary,
		)
		(*dbObj).checkErr(err)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id

}