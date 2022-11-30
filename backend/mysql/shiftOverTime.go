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
func(dbObj *DB) SelectShiftOverTime(selectKey int, value... interface{}) *[]table.ShiftOverTimeTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectSingleByCaseId
		break
	case 2:
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectAllByShiftId
		break
	default:
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectAll
		break
	}
	shiftOverTime := new(table.ShiftOverTimeTable)
	carry := []table.ShiftOverTimeTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&shiftOverTime.CaseId,
			&shiftOverTime.ShiftId,
			&shiftOverTime.InitiatorOnOverTime,
			&shiftOverTime.InitiatorOffOverTime,
			&shiftOverTime.Reason,
			&shiftOverTime.CaseProcess,
			&shiftOverTime.SpecifyTag,
			&shiftOverTime.CreateTime,
			&shiftOverTime.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *shiftOverTime)
		}
	}
	return &carry
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftOverTime.Delete)
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
func(dbObj *DB) UpdateShiftOverTime(updateKey int, data *table.ShiftOverTimeTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	defer func ()  {
		(*dbObj).containers.shiftOverTime = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftOverTime.UpdateSingle
		(*dbObj).containers.shiftOverTime = append(
			(*dbObj).containers.shiftOverTime,
			(*data).InitiatorOnOverTime,
			(*data).InitiatorOffOverTime,
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
	_, err = stmt.Exec((*dbObj).containers.shiftOverTime...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertShiftOverTime(data *table.ShiftOverTimeTable) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftOverTime.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).ShiftId,
		(*data).InitiatorOnOverTime,
		(*data).InitiatorOffOverTime,
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