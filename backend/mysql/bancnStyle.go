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
//  1 => styleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchStyle(selectKey int, value... interface{}) *[]table.BanchStyle {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).BanchStyle.SelectAll
		break
	case 1:
		querys = (*query.MysqlSingleton()).BanchStyle.SelectSingleByStyleId
		break
	case 2:
		querys = (*query.MysqlSingleton()).BanchStyle.SelectAllByBanchId
		break
	default:
		querys = (*query.MysqlSingleton()).BanchStyle.SelectAll
		break
	}
	banchStyle := new(table.BanchStyle)
	carry := []table.BanchStyle{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&banchStyle.StyleId,
			&banchStyle.BanchId,
			&banchStyle.Icon,
			&banchStyle.RestTime,
			&banchStyle.TimeRangeName,
			&banchStyle.OnShiftTime,
			&banchStyle.OffShiftTime,
			&banchStyle.CreateTime,
			&banchStyle.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *banchStyle)
		}
	}
	return &carry
}

// style的唯一id
func(dbObj *DB) DeleteBanchStyle(deleteKey int, styleId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).banchStyleMux.Lock()
	defer (*dbObj).banchStyleMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).BanchStyle.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(styleId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateBanchStyle(updateKey int, data *table.BanchStyle, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).banchStyleMux.Lock()
	defer (*dbObj).banchStyleMux.Unlock()
	defer func ()  {
		(*dbObj).containers.banchStyle = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).BanchStyle.UpdateSingle
		(*dbObj).containers.banchStyle = append(
			(*dbObj).containers.banchStyle,
			(*data).Icon,
			(*data).RestTime,
			(*data).TimeRangeName,
			(*data).OnShiftTime,
			(*data).OffShiftTime,
			(*data).LastModify,
			(*data).StyleId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.banchStyle...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertBanchStyle(data *table.BanchStyle) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).banchStyleMux.Lock()
	defer (*dbObj).banchStyleMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).BanchStyle.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).BanchId,
		(*data).Icon,
		(*data).RestTime,
		(*data).TimeRangeName,
		(*data).OnShiftTime,
		(*data).OffShiftTime,
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