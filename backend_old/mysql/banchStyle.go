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

// 0 => all, value => nil
//  1 => styleId, value => int64
//  2=> banchId, value => int64
// . 3 => banchId and companyCode, value => int64 string
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
	case 3:
		querys = (*query.MysqlSingleton()).BanchStyle.SelectByCompanyCode
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
			&banchStyle.DelFlag,
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

// . 0 styleId int64
// . 1 styleId int64, companyCode string
func(dbObj *DB) DeleteBanchStyle(deleteKey int, value... interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).banchStyleMux.Lock()
	defer (*dbObj).banchStyleMux.Unlock()
	querys := ""
	switch deleteKey {
	case 0:
		querys = (*query.MysqlSingleton()).BanchStyle.Delete
		break
	case 1:
		querys = (*query.MysqlSingleton()).BanchStyle.DeleteByCompanyCode
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

// 0 => all
// styleId  1 => companyCode string
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
	case 1:
		querys = (*query.MysqlSingleton()).BanchStyle.UpdateByCompanyCode
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
		break
	}
	(*dbObj).containers.banchStyle = append(
		(*dbObj).containers.banchStyle,
		value...
	)
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

func(dbObj *DB) InsertBanchStyle(data *table.BanchStyle, value... interface{}) (bool, int64) {

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