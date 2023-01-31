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

// 0 => admin, value => user companyCode string, quitworkuser companyCode,year int,month int
// 1=> manager, value => user companyCode string, quitworkuser companyCode, banchId int64, banchName string,year int,month int
// 2 => person, value => user userId int64 ,year int,month int
// 3 => copy admin, value => performanceId int64, user companyCode string, quitworkuser companyCode
// 4 => copy manage. value => performanceId int64, user companyCode string, quitworkuser companyCode, banchId int64, banchName string
// 5 => copy person, value => performanceId int64, user userId int64, quitworkuser userId int64
func(dbObj *DB) SelectPerformance(selectKey int, value... interface{}) *[]table.PerformanceExtend {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).Performance.SelectAllByAdmin
		break
	case 1:
		querys = (*query.MysqlSingleton()).Performance.SelectAllByManager
		break
	case 2:
		querys = (*query.MysqlSingleton()).Performance.SelectAllByPerson
		break
	case 3:
		querys = (*query.MysqlSingleton()).Performance.SelectSingleByAdmin
		break
	case 4:
		querys = (*query.MysqlSingleton()).Performance.SelectSingleByManager
		break
	case 5:
		querys = (*query.MysqlSingleton()).Performance.SelectSingleByPerson
		break
	}
	per := new(table.PerformanceExtend)
	carry := []table.PerformanceExtend{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&per.PerformanceId,
			&per.UserId,
			&per.Year,
			&per.Month,
			&per.BanchId,
			&per.Goal,
			&per.Attitude,
			&per.Efficiency,
			&per.Professional,
			&per.Directions,
			&per.BeLate,
			&per.DayOffNotOnRule,
			&per.BanchName,
			&per.CreateTime,
			&per.LastModify,
			&per.UserName,
			&per.CompanyId,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *per)
		}
	}

	return &carry
}

// 0 => admin => performanceId, companyCode, companyCode
// 1 => manager => performanceId, companyCode, companyCode, performanceBanchId, performanceBanchName
// 2 => personal => performanceId, userId
func(dbObj *DB) UpdatePerformance(updateKey int, data *table.Performance, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).performanceMux.Lock()
	defer (*dbObj).performanceMux.Unlock()
	defer func ()  {
		(*dbObj).containers.performance = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).Performance.UpdateByAdmin
		(*dbObj).containers.performance = append(
			(*dbObj).containers.performance,
			(*data).BanchId,
			(*data).Goal,
			(*data).Attitude,
			(*data).Efficiency,
			(*data).Professional,
			(*data).Directions,
			(*data).BeLate,
			(*data).DayOffNotOnRule,
			(*data).BanchName,
			(*data).LastModify,
		)
		break;
	case 1:
		querys = (*query.MysqlSingleton()).Performance.UpdateByManager
		(*dbObj).containers.performance = append(
			(*dbObj).containers.performance,
			(*data).BanchId,
			(*data).Goal,
			(*data).Attitude,
			(*data).Efficiency,
			(*data).Professional,
			(*data).Directions,
			(*data).BeLate,
			(*data).DayOffNotOnRule,
			(*data).BanchName,
			(*data).LastModify,
		)
		break;
	case 2:
		querys = (*query.MysqlSingleton()).Performance.UpdateByPerson
		(*dbObj).containers.performance = append(
			(*dbObj).containers.performance,
			(*data).Goal,
			(*data).LastModify,
		)
		break;
	
	}
	(*dbObj).containers.performance = append(
		(*dbObj).containers.performance,
		value...
	)
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.performance...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}
//  
//   1 admin, performanceId
//   2 manager, performanceId, p.banchId, p.userId != ?
func(dbObj *DB) DeletePerformance(deleteKey int, value...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).performanceMux.Lock()
	defer (*dbObj).performanceMux.Unlock()
	querys := ""
	switch deleteKey {
	case 0:
		querys = (*query.MysqlSingleton()).Performance.DeleteByAdmin
		break
	case 1:
		querys = (*query.MysqlSingleton()).Performance.DeleteByManage
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


func(dbObj *DB) InsertPerformance(data *table.Performance) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).performanceMux.Lock()
	defer (*dbObj).performanceMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Performance.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)

	res, err := stmt.Exec(
		(*data).UserId,
		(*data).Year,
		(*data).Month,
		(*data).BanchId,
		(*data).Goal,
		(*data).Attitude,
		(*data).Efficiency,
		(*data).Professional,
		(*data).Directions,
		(*data).BeLate,
		(*data).DayOffNotOnRule,
		(*data).BanchName,
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

// 0 => admin, value => user companyCode, qu companyCode, startYear, endYear, userName * 3
// 1 => manage, value => user companyCode, qu companyCode, banchId, banchName, startYear, endYear, userName * 3
// 2 => person, value => user companyCode, qu companyCode, userId, startYear, endYear, userName * 3
func(dbObj *DB) SelectYearPerformance(selectKey int, value... interface{}) *[]table.YearPerformance {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).Performance.SelectYearPerformanceByAdmin
		break
	case 1:
		querys = (*query.MysqlSingleton()).Performance.SelectYearPerformanceByManage
		break
	case 2:
		querys = (*query.MysqlSingleton()).Performance.SelectYearPerformanceByPerson
		break
	}
	per := new(table.YearPerformance)
	carry := []table.YearPerformance{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&per.UserId,
			&per.Year,
			&per.UserName,
			&per.Avg,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *per)
		}
	}

	return &carry
}