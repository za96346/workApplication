package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (

	// "runtime"
	panichandler "backend/panicHandler"
	"backend/mysql/query"
	"backend/mysql/table"

	_ "github.com/go-sql-driver/mysql"
)

// 0 => 根據banchid, value => banchId, year, month
func(dbObj *DB) SelectShiftEditLog(selectKey int, value... interface{}) *[]table.ShiftEditLog {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftEditLog.SelectByBanchId
		break
	}
	log := new(table.ShiftEditLog)
	carry := []table.ShiftEditLog{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&log.LogId,
			&log.Year,
			&log.Month,
			&log.BanchId,
			&log.Msg,
			&log.CreateTime,
			&log.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *log)
		}
	}
	return &carry
}

func(dbObj *DB) InsertShiftEditLog(data *table.ShiftEditLog) (bool, int64) {
	defer panichandler.Recover()
	(*dbObj).shiftEditLogMux.Lock()
	defer (*dbObj).shiftEditLogMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftEditLog.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).Year,
		(*data).Month,
		(*data).BanchId,
		(*data).Msg,
	)
	(*dbObj).checkErr(err)
	id, _:= res.LastInsertId()
	if err != nil {
		return false, id
	}
	return true, id
}