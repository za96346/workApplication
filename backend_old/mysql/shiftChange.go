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
//  2 => initiatorShiftId, value => int64
//  3 => requestedShiftId, value => int64
func(dbObj *DB) SelectShiftChange(selectKey int, value... interface{}) *[]table.ShiftChangeTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftChange.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).ShiftChange.SelectSingleByCaseId
		break
	case 2:
		querys = (*query.MysqlSingleton()).ShiftChange.SelectAllByInitiatorShiftId
		break
	case 3:
		querys = (*query.MysqlSingleton()).ShiftChange.SelectAllByRequestedShiftId
		break
	default:
		querys = (*query.MysqlSingleton()).ShiftChange.SelectAll
		break
	}
	shiftChange := new(table.ShiftChangeTable)
	carry := []table.ShiftChangeTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&shiftChange.CaseId,
			&shiftChange.InitiatorShiftId,
			&shiftChange.RequestedShiftId,
			&shiftChange.Reason,
			&shiftChange.CaseProcess,
			&shiftChange.SpecifyTag,
			&shiftChange.CreateTime,
			&shiftChange.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *shiftChange)
		}
	}
	return &carry
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftChange.Delete)
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
func(dbObj *DB) UpdateShiftChange(updateKey int, data *table.ShiftChangeTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	defer func ()  {
		(*dbObj).containers.shiftChange = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftChange.UpdateSingle
		(*dbObj).containers.shiftChange = append(
			(*dbObj).containers.shiftChange,
			(*data).InitiatorShiftId,
			(*data).RequestedShiftId,
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
	_, err = stmt.Exec((*dbObj).containers.shiftChange...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertShiftChange(data *table.ShiftChangeTable) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftChange.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).InitiatorShiftId,
		(*data).RequestedShiftId,
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