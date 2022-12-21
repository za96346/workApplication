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
