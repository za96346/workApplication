package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (

	// "runtime"
	"backend/query"
	"backend/table"

	_ "github.com/go-sql-driver/mysql"
)

func(dbObj *DB) InsertLog(data *table.Log) (bool, int64) {
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Log.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
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