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
// . 2 => year, value => int
func(dbObj *DB) SelectPaidVocation (selectKey int, value... interface{}) *[]table.PaidVocation {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).PaidVocation.SelectAll
		break
	case 1:
		querys = (*query.MysqlSingleton()).PaidVocation.SelectAllByUserId
		break
	case 2:
		querys = (*query.MysqlSingleton()).PaidVocation.SelectAllByTime
		break
	}
	paidVocation := new(table.PaidVocation)
	carry := []table.PaidVocation{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&paidVocation.PaidVocationId,
			&paidVocation.UserId,
			&paidVocation.Year,
			&paidVocation.Count,
			&paidVocation.CreateTime,
			&paidVocation.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *paidVocation)
		}
	}
	return &carry
}

// paidVocation 的唯一id
func(dbObj *DB) DeletePaidVocation (deleteKey int, paidVocationId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).paidVocation.Lock()
	defer (*dbObj).paidVocation.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).PaidVocation.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(paidVocationId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all, by paidVocationId int64
func(dbObj *DB) UpdatePaidVocation (updateKey int, data *table.PaidVocation, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).paidVocation.Lock()
	defer (*dbObj).paidVocation.Unlock()
	defer func ()  {
		(*dbObj).containers.paidVocation = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).PaidVocation.UpdateSingle
		(*dbObj).containers.paidVocation = append(
			(*dbObj).containers.paidVocation,
			(*data).Year,
			(*data).Count,
			(*data).LastModify,
			(*data).PaidVocationId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.paidVocation...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertPaidVocation (data *table.PaidVocation) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).paidVocation.Lock()
	defer (*dbObj).paidVocation.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).PaidVocation.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).UserId,
		(*data).Year,
		(*data).Count,
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