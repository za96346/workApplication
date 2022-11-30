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
//  1 => waitId, value => int64
//  2 => userId, value => int64
//  3 => companyId, value => int64
//  4 => comapnyId && userId, value => int64, int64
//  5 => companyId, value => int64  this is only mysql query
func(dbObj *DB) SelectWaitCompanyReply(selectKey int, value... interface{}) *[]table.WaitCompanyReply {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).WaitCompanyReply.SelectAll
		break
	case 1:
		querys = (*query.MysqlSingleton()).WaitCompanyReply.SelectSingleByWaitId
		break
	case 2:
		querys = (*query.MysqlSingleton()).WaitCompanyReply.SelectAllByUserId
		break
	case 3:
		querys = (query.MysqlSingleton()).WaitCompanyReply.SelectAllByCompanyId
		break
	case 4:
		querys = (query.MysqlSingleton()).WaitCompanyReply.SelectAllByCompanyIdAndUserId
		break
	case 5:
		querys = (query.MysqlSingleton()).WaitCompanyReply.SelectAllJoinUserTable
		break
	default:
		querys = (*query.MysqlSingleton()).WaitCompanyReply.SelectAll
		break
	}
	WaitCompanyReply := new(table.WaitCompanyReply)
	carry := []table.WaitCompanyReply{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		if selectKey == 5 {
			err = res.Scan(
				&WaitCompanyReply.WaitId,
				&WaitCompanyReply.UserId,
				&WaitCompanyReply.UserName,
				&WaitCompanyReply.CompanyId,
				&WaitCompanyReply.SpecifyTag,
				&WaitCompanyReply.IsAccept,
				&WaitCompanyReply.CreateTime,
				&WaitCompanyReply.LastModify,
			)
		} else  {
			err = res.Scan(
				&WaitCompanyReply.WaitId,
				&WaitCompanyReply.UserId,
				&WaitCompanyReply.CompanyId,
				&WaitCompanyReply.SpecifyTag,
				&WaitCompanyReply.IsAccept,
				&WaitCompanyReply.CreateTime,
				&WaitCompanyReply.LastModify,
			)
		}
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *WaitCompanyReply)
		}
	}
	return &carry
}

// wait company reply 的唯一id
func(dbObj *DB) DeleteWaitCompanyReply(deleteKey int, waitId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).waitCompanyReply.Lock()
	defer (*dbObj).waitCompanyReply.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).WaitCompanyReply.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(waitId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all, by waitId => int64
func(dbObj *DB) UpdateWaitCompanyReply(updateKey int, data *table.WaitCompanyReply, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).waitCompanyReply.Lock()
	defer (*dbObj).waitCompanyReply.Unlock()
	defer func ()  {
		(*dbObj).containers.waitCompanyReply = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).WaitCompanyReply.UpdateSingle
		(*dbObj).containers.waitCompanyReply= append(
			(*dbObj).containers.waitCompanyReply,
			(*data).SpecifyTag,
			(*data).IsAccept,
			(*data).LastModify,
			(*data).WaitId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.waitCompanyReply...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertWaitCompanyReply (data *table.WaitCompanyReply) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).waitCompanyReply.Lock()
	defer (*dbObj).waitCompanyReply.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).WaitCompanyReply.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).UserId,
		(*data).CompanyId,
		(*data).SpecifyTag,
		(*data).IsAccept,
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