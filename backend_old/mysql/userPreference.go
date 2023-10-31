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

// 0 => 全部, value => nil
//  1 => 使用者id, value => int64
func(dbObj *DB) SelectUserPreference(selectKey int, value... interface{}) *[]table.UserPreferenceTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).UserPreference.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).UserPreference.SelectSingleByUserId
		break
	default:
		querys = (*query.MysqlSingleton()).UserPreference.SelectAll
		break
	}
	userPreference := new(table.UserPreferenceTable)
	carry := []table.UserPreferenceTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&userPreference.UserId,
			&userPreference.Style,
			&userPreference.FontSize,
			&userPreference.SelfPhoto,
			&userPreference.CreateTime,
			&userPreference.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *userPreference)
		}
	}


	return &carry
}

//使用者的唯一id
func(dbObj *DB) DeleteUserPreference(deleteKey int, userId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).UserPreference.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(userId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateUserPreference(updateKey int, data *table.UserPreferenceTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	defer func ()  {
		(*dbObj).containers.userPreference = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).UserPreference.UpdateSingle
		(*dbObj).containers.userPreference = append(
			(*dbObj).containers.userPreference,
			(*data).Style,
			(*data).FontSize,
			(*data).SelfPhoto,
			(*data).LastModify,
			(*data).UserId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.userPreference...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertUserPreference(data *table.UserPreferenceTable) (bool, int64) {
	///
	defer panichandler.Recover()
	(*dbObj).userPreferenceMux.Lock()
	defer  (*dbObj).userPreferenceMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).UserPreference.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).UserId,
		(*data).Style,
		(*data).FontSize,
		(*data).SelfPhoto,
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