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
//  1 => weekendId, value => int64
//  2 => companyId, value => int64
func(dbObj *DB) SelectWeekendSetting(selectKey int, value... interface{}) *[]table.WeekendSetting {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).WeekendSetting.SelectAll
		break
	case 1:
		querys = (*query.MysqlSingleton()).WeekendSetting.SelectSingleByWeekendId
		break
	case 2:
		querys = (*query.MysqlSingleton()).WeekendSetting.SelectAllByCompanyId
		break
	default:
		querys = (*query.MysqlSingleton()).WeekendSetting.SelectAll
		break
	}
	weekendSetting := new(table.WeekendSetting)
	carry := []table.WeekendSetting{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&weekendSetting.WeekendId,
			&weekendSetting.CompanyId,
			&weekendSetting.Date,
			&weekendSetting.CreateTime,
			&weekendSetting.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			weekendSetting.Date = weekendSetting.Date[0 : 10]
			carry = append(carry, *weekendSetting)
		}
	}
	return &carry
}

// weekend setting 的唯一id
func(dbObj *DB) DeleteWeekendSetting(deleteKey int, weekendId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).weekendSetting.Lock()
	defer (*dbObj).weekendSetting.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).WeekendSetting.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(weekendId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all, by weekendId int64
func(dbObj *DB) UpdateWeekendSetting(updateKey int, data *table.WeekendSetting, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).weekendSetting.Lock()
	defer (*dbObj).weekendSetting.Unlock()
	defer func ()  {
		(*dbObj).containers.weekendSetting = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).WeekendSetting.UpdateSingle
		(*dbObj).containers.weekendSetting = append(
			(*dbObj).containers.weekendSetting,
			(*data).Date,
			(*data).LastModify,
			(*data).WeekendId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.weekendSetting...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertWeekendSetting (data *table.WeekendSetting) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).weekendSetting.Lock()
	defer (*dbObj).weekendSetting.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).WeekendSetting.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).CompanyId,
		(*data).Date,
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