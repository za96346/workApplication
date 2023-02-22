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

// 0 => all, value =>banchId, companyId, year, month
//  1 => 班表id, value => int64
// 2=> 班表總共, value =>banchId, companyId, year, month
func(dbObj *DB) SelectShift(selectKey int, value... interface{}) *[]table.ShiftExtend {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).Shift.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).Shift.SelectSingleByShiftId
		break
	default:
		querys = (*query.MysqlSingleton()).Shift.SelectAll
		break
	}
	shift := new(table.ShiftExtend)
	carry := []table.ShiftExtend{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&shift.ShiftId,
			&shift.UserId,
			&shift.BanchStyleId,
			&shift.BanchId,
			&shift.Year,
			&shift.Month,
			&shift.Icon,
			&shift.OnShiftTime,
			&shift.OffShiftTime,
			&shift.RestTime,
			&shift.PunchIn,
			&shift.PunchOut,
			&shift.SpecifyTag,
			&shift.CreateTime,
			&shift.LastModify,
			&shift.UserName,
			&shift.Permission,
			&shift.Banch,
			&shift.EmployeeNumber,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *shift)
		}
	}
	return &carry
}

// 班表的唯一id (關聯資料表	shiftovertime shiftchange forgetpunch lateexcused dayoff 都上鎖)
func(dbObj *DB) DeleteShift(deleteKey int, shiftId interface{}) bool {
	defer panichandler.Recover()

	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()

	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()

	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()

	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()

	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()

	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()

	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Shift.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(shiftId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateShift(updateKey int, data *table.ShiftTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	defer func ()  {
		(*dbObj).containers.shift = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).Shift.UpdateSingle
		(*dbObj).containers.shift = append(
			(*dbObj).containers.shift,
			(*data).BanchStyleId,
			(*data).Icon,
			(*data).OnShiftTime,
			(*data).OffShiftTime,
			(*data).RestTime,
			(*data).PunchIn,
			(*data).PunchOut,
			(*data).SpecifyTag,
			(*data).LastModify,
			(*data).ShiftId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.shift...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertShift(data *table.ShiftTable) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Shift.InsertAll)
	defer stmt.Close()


	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).UserId,
		(*data).BanchStyleId,
		(*data).BanchId,
		(*data).Year,
		(*data).Month,
		(*data).Icon,
		(*data).OnShiftTime,
		(*data).OffShiftTime,
		(*data).RestTime,
		(*data).PunchIn,
		(*data).PunchOut,
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