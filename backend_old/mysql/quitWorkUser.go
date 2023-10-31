package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (

	// "runtime"
	"backend/mysql/query"
	"backend/mysql/table"

	"backend/panicHandler"

	_ "github.com/go-sql-driver/mysql"
)

// 0 => all, value => nil
//  1 => quitId, value => int64
//   2 => userId, value => int64
//   3 => companyCode, value => string
//   4=> companyCode && userId ,  value string && int64
func(dbObj *DB) SelectQuitWorkUser(selectKey int, value... interface{}) *[]table.QuitWorkUser {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).QuitWorkUser.SelectAll
		break
	case 1:
		querys = (*query.MysqlSingleton()).QuitWorkUser.SelectSingleByQuitId
		break
	case 2:
		 querys = (*query.MysqlSingleton()).QuitWorkUser.SelectSingleByUserId
		 break
	case 3:
		querys = (*query.MysqlSingleton()).QuitWorkUser.SelectAllByCompanyCode
		break
	case 4:
		querys = (*query.MysqlSingleton()).QuitWorkUser.SelectSingleByCompanyCodeAndUserId
		break
	default:
		querys = (*query.MysqlSingleton()).QuitWorkUser.SelectAll
		break
	}
	quitWorkUser := new(table.QuitWorkUser)
	carry := []table.QuitWorkUser{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&quitWorkUser.QuitId,
			&quitWorkUser.UserId,
			&quitWorkUser.CompanyCode,
			&quitWorkUser.UserName,
			&quitWorkUser.EmployeeNumber,
			&quitWorkUser.Account,
			&quitWorkUser.OnWorkDay,
			&quitWorkUser.Banch,
			&quitWorkUser.Permession,
			&quitWorkUser.MonthSalary,
			&quitWorkUser.PartTimeSalary,
			&quitWorkUser.CreateTime,
			&quitWorkUser.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *quitWorkUser)
		}
	}
	return &carry
}

// 0, quit work suer 的唯一id
// .  1 => userId int64, companyCode string, companyCode string
func(dbObj *DB) DeleteQuitWorkUser(deleteKey int, value... interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).quitWorkUserMux.Lock()
	defer (*dbObj).quitWorkUserMux.Unlock()
	querys := ""
	switch deleteKey {
	case 0:
			querys = (*query.MysqlSingleton()).QuitWorkUser.Delete
			break
	case 1:
		querys = (*query.MysqlSingleton()).QuitWorkUser.DeleteByJoinUser
		break

	}
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(value...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateQuitWorkUser(updateKey int, data *table.QuitWorkUser, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).quitWorkUserMux.Lock()
	defer (*dbObj).quitWorkUserMux.Unlock()
	defer func ()  {
		(*dbObj).containers.quitWorkUser = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).QuitWorkUser.UpdateSingle
		(*dbObj).containers.quitWorkUser= append(
			(*dbObj).containers.quitWorkUser,
			(*data).UserId,
			(*data).CompanyCode,
			(*data).UserName,
			(*data).EmployeeNumber,
			(*data).Account,
			(*data).OnWorkDay,
			(*data).Banch,
			(*data).Permession,
			(*data).MonthSalary,
			(*data).PartTimeSalary,
			(*data).CreateTime,
			(*data).LastModify,
			(*data).QuitId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.quitWorkUser...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertQuitWorkUser(data *table.QuitWorkUser) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).quitWorkUserMux.Lock()
	defer (*dbObj).quitWorkUserMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).QuitWorkUser.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).UserId,
		(*data).CompanyCode,
		(*data).UserName,
		(*data).EmployeeNumber,
		(*data).Account,
		(*data).OnWorkDay,
		(*data).Banch,
		(*data).Permession,
		(*data).MonthSalary,
		(*data).PartTimeSalary,
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