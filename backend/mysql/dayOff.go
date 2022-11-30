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
func(dbObj *DB) SelectDayOff(selectKey int, value... interface{}) *[]table.DayOffTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).DayOff.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).DayOff.SelectSingleByCaseId
		break
	case 2:
		querys = (*query.MysqlSingleton()).DayOff.SelectAllByShiftId
		break
	default:
		querys = (*query.MysqlSingleton()).DayOff.SelectAll
		break
	}
	dayOff := new(table.DayOffTable)
	carry := []table.DayOffTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&dayOff.CaseId,
			&dayOff.ShiftId,
			&dayOff.DayOffType,
			&dayOff.Reason,
			&dayOff.CaseProcess,
			&dayOff.SpecifyTag,
			&dayOff.CreateTime,
			&dayOff.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *dayOff)
		}
	}
	return &carry
}

// 案件的唯一id
func(dbObj *DB) DeleteDayOff(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).DayOff.Delete)
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
func(dbObj *DB) UpdateDayOff(updateKey int, data *table.DayOffTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	defer func ()  {
		(*dbObj).containers.dayOff = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).DayOff.UpdateSingle
		(*dbObj).containers.dayOff = append(
			(*dbObj).containers.dayOff,
			(*data).DayOffType,
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
	_, err = stmt.Exec((*dbObj).containers.dayOff...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertDayOff(data *table.DayOffTable) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).DayOff.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).ShiftId,
		(*data).DayOffType,
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