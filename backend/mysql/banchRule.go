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
//  1 => ruleId, value => int64
//  2=> banchId, value => int64
func(dbObj *DB) SelectBanchRule(selectKey int, value... interface{}) *[]table.BanchRule {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).BanchRule.SelectAll
		break
	case 1:
		querys = (*query.MysqlSingleton()).BanchRule.SelectSingleByRuleId
		break
	case 2:
		querys = (*query.MysqlSingleton()).BanchRule.SelectAllByBanchId
		break
	default:
		querys = (*query.MysqlSingleton()).BanchRule.SelectAll
		break
	}
	banchRule := new(table.BanchRule)
	carry := []table.BanchRule{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&banchRule.RuleId,
			&banchRule.BanchId,
			&banchRule.MaxPeople,
			&banchRule.MinPeople,
			&banchRule.WeekDay,
			&banchRule.WeekType,
			&banchRule.OnShiftTime,
			&banchRule.OffShiftTime,
			&banchRule.CreateTime,
			&banchRule.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *banchRule)
		}
	}
	return &carry
}

// rule的唯一id
func(dbObj *DB) DeleteBanchRule(deleteKey int, ruleId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).banchRuleMux.Lock()
	defer (*dbObj).banchRuleMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).BanchRule.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(ruleId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateBanchRule(updateKey int, data *table.BanchRule, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).banchRuleMux.Lock()
	defer (*dbObj).banchRuleMux.Unlock()
	defer func ()  {
		(*dbObj).containers.banchRule = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).BanchRule.UpdateSingle
		(*dbObj).containers.banchRule= append(
			(*dbObj).containers.banchRule,
			(*data).MaxPeople,
			(*data).MinPeople,
			(*data).WeekDay,
			(*data).WeekType,
			(*data).OnShiftTime,
			(*data).OffShiftTime,
			(*data).LastModify,
			(*data).RuleId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.banchRule...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

func(dbObj *DB) InsertBanchRule(data *table.BanchRule) (bool, int64) {

	defer panichandler.Recover()
	(*dbObj).banchRuleMux.Lock()
	defer (*dbObj).banchRuleMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).BanchRule.InsertAll)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	res, err := stmt.Exec(
		(*data).BanchId,
		(*data).MaxPeople,
		(*data).MinPeople,
		(*data).WeekDay,
		(*data).WeekType,
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