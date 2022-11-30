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

// 0 => all, value => nil
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
func(dbObj *DB) SelectForgetPunch(selectKey int, value... interface{}) *[]table.ForgetPunchTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectSingleByCaseId
		break
	case 2:
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectAllByShiftId
		break
	default:
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectAll
		break
	}
	forgetPunch := new(table.ForgetPunchTable)
	carry := []table.ForgetPunchTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&forgetPunch.CaseId,
			&forgetPunch.ShiftId,
			&forgetPunch.TargetPunch,
			&forgetPunch.Reason,
			&forgetPunch.CaseProcess,
			&forgetPunch.SpecifyTag,
			&forgetPunch.CreateTime,
			&forgetPunch.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *forgetPunch)
		}
	}
	return &carry
}

// 案件的唯一id
func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ForgetPunch.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}


// 0 => all
func(dbObj *DB) UpdateForgetPunch(updateKey int, data *table.ForgetPunchTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	defer func ()  {
		(*dbObj).containers.forgetPunch = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).ForgetPunch.UpdateSingle
		(*dbObj).containers.forgetPunch = append(
			(*dbObj).containers.forgetPunch,
			(*data).TargetPunch,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
			(*data).LastModify,
			(*data).CaseId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.forgetPunch...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertForgetPunch(data *table.ForgetPunchTable) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ForgetPunch.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).ShiftId,
		(*data).TargetPunch,
		(*data).Reason,
		(*data).CaseProcess,
		(*data).SpecifyTag,
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