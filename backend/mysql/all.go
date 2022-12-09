package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (

	// "runtime"
	"backend/query"
	"backend/response"

	"backend/panicHandler"

	_ "github.com/go-sql-driver/mysql"
)

// 0 => all, value => nil
//  1 => ruleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectAllUser(selectKey int, value... interface{}) *[]response.User {
	defer panichandler.Recover()
	querys := (*query.MysqlSingleton()).User.SelectAllNeedJoin
	user := new(response.User)
	carry := []response.User{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&user.UserId,
			&user.CompanyCode,
			&user.UserName,
			&user.EmployeeNumber,
			&user.OnWorkDay,
			&user.Banch,
			&user.Permession,
			&user.WorkState,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *user)
		}
	}
	return &carry
}


// .  userId companyCode int64 string
func(dbObj *DB) InsertQuitWorkUserBySelectUser(userId int64, companyCode string) (bool, int64) {
	///
		defer panichandler.Recover()
		(*dbObj).quitWorkUserMux.Lock()
		defer (*dbObj).quitWorkUserMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).QuitWorkUser.InsertBySelectUser)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(userId, companyCode)
		(*dbObj).checkErr(err)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id

}