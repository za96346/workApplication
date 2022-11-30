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

// . 0 => all, value => nil
// . 1 => userId, value => int64
// . 2 => year && month, value => int && int
func(dbObj *DB) SelectWorkTime (selectKey int, value... interface{}) *[]table.WorkTime {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).WorkTime.SelectAll
		break
	case 1:
		querys = (*query.MysqlSingleton()).WorkTime.SelectAllByUserId
		break
	case 2:
		querys = (*query.MysqlSingleton()).WorkTime.SelectAllByTime
		break
	}
	workTime := new(table.WorkTime)
	carry := []table.WorkTime{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&workTime.WorkTimeId,
			&workTime.UserId,
			&workTime.Year,
			&workTime.Month,
			&workTime.WorkHours,
			&workTime.TimeOff,
			&workTime.UsePaidVocation,
			&workTime.CreateTime,
			&workTime.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *workTime)
		}
	}
	return &carry
}

// workTime 的唯一id
func(dbObj *DB) DeleteWorkTime (deleteKey int, workTimeId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).workTime.Lock()
	defer (*dbObj).workTime.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).WorkTime.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(workTimeId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all, by workTimeId int64
func(dbObj *DB) UpdateWorkTime (updateKey int, data *table.WorkTime, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).workTime.Lock()
	defer (*dbObj).workTime.Unlock()
	defer func ()  {
		(*dbObj).containers.worktime = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).WorkTime.UpdateSingle
		(*dbObj).containers.worktime = append(
			(*dbObj).containers.worktime,
			(*data).Year,
			(*data).Month,
			(*data).WorkHours,
			(*data).TimeOff,
			(*data).UsePaidVocation,
			(*data).LastModify,
			(*data).WorkTimeId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.worktime...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertWorkTime (data *table.WorkTime) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).workTime.Lock()
	defer (*dbObj).workTime.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).WorkTime.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).UserId,
		(*data).Year,
		(*data).Month,
		(*data).WorkHours,
		(*data).TimeOff,
		(*data).UsePaidVocation,
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