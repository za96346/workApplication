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
//  1 => caseId, value => int64
//  2 => shiftId, value => int64
func(dbObj *DB) SelectLateExcused(selectKey int, value... interface{}) *[]table.LateExcusedTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).LateExcused.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).LateExcused.SelectSingleByCaseId
		break
	case 2:
		querys = (*query.MysqlSingleton()).LateExcused.SelectAllByShiftId
		break
	default:
		querys = (*query.MysqlSingleton()).LateExcused.SelectAll
		break
	}
	lateExcused := new(table.LateExcusedTable)
	carry := []table.LateExcusedTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&lateExcused.CaseId,
			&lateExcused.ShiftId,
			&lateExcused.LateExcusedType,
			&lateExcused.Reason,
			&lateExcused.CaseProcess,
			&lateExcused.SpecifyTag,
			&lateExcused.CreateTime,
			&lateExcused.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *lateExcused)
		}
	}
	return &carry
}

// 案件的唯一id
func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).LateExcused.Delete)
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
func(dbObj *DB) UpdateLateExcused(updateKey int, data *table.LateExcusedTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	defer func ()  {
		(*dbObj).containers.lateExcused = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).LateExcused.UpdateSingle
		(*dbObj).containers.lateExcused = append(
			(*dbObj).containers.lateExcused,
			(*data).LateExcusedType,
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
	_, err = stmt.Exec((*dbObj).containers.lateExcused...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertLateExcused(data *table.LateExcusedTable) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).LateExcused.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).ShiftId,
		(*data).LateExcusedType,
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