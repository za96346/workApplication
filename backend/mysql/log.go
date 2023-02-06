package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (

	// "runtime"
	panichandler "backend/panicHandler"
	"backend/query"
	"backend/table"

	_ "github.com/go-sql-driver/mysql"
)
func(dbObj *DB) SelectLog(selectKey int, value... interface{}) *[]table.Log {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).Log.SelectAll
		break
	}
	log := new(table.Log)
	carry := []table.Log{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&log.LogId,
			&log.UserId,
			&log.UserName,
			&log.CompanyId,
			&log.CompanyCode,
			&log.Permession,
			&log.Routes,
			&log.Ip,
			&log.Params,
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

func(dbObj *DB) InsertLog(data *table.Log) (bool, int64) {
	defer panichandler.Recover()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Log.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).UserId,
		(*data).UserName,
		(*data).CompanyId,
		(*data).CompanyCode,
		(*data).Permession,
		(*data).Routes,
		(*data).Ip,
		(*data).Params,
		(*data).Msg,
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