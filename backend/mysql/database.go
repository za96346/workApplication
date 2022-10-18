package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "runtime"
	"backend/query"
	"backend/table"
	"sync"
	"time"

	"backend/panicHandler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var dbSingletonMux = new(sync.Mutex)
var dbInstance *DB

type DB struct {
	companyMux *sync.RWMutex
	userPreferenceMux *sync.RWMutex
	userMux *sync.RWMutex
	companyBanchMux *sync.RWMutex
	shiftMux *sync.RWMutex
	shiftChangeMux *sync.RWMutex
	shiftOverTimeMux *sync.RWMutex
	dayOffMux *sync.RWMutex
	lateExcusedMux *sync.RWMutex
	banchStyleMux *sync.RWMutex
	banchRuleMux *sync.RWMutex
	forgetPunchMux *sync.RWMutex
	MysqlDB *sql.DB // 要先使用連線方法後才能使用這個
	containers
}

type containers struct {
	user []interface{}
	userPreference []interface{}
	company []interface{}
	companyBanch []interface{}
	shift []interface{}
	shiftOverTime []interface{}
	shiftChange []interface{}
	forgetPunch []interface{}
	dayOff []interface{}
	lateExcused []interface{}
	banchStyle []interface{}
	banchRule []interface{}
}

func Singleton() *DB {
	defer panichandler.Recover()
	if dbInstance == nil {
		dbSingletonMux.Lock()
		defer dbSingletonMux.Unlock()
		if dbInstance == nil {
			dbInstance = &DB{
				companyMux: new(sync.RWMutex),
				userPreferenceMux: new(sync.RWMutex),
				userMux: new(sync.RWMutex),
				companyBanchMux: new(sync.RWMutex),
				shiftMux: new(sync.RWMutex),
				shiftChangeMux: new(sync.RWMutex),
				shiftOverTimeMux: new(sync.RWMutex),
				dayOffMux: new(sync.RWMutex),
				lateExcusedMux: new(sync.RWMutex),
				forgetPunchMux: new(sync.RWMutex),
				banchStyleMux: new(sync.RWMutex),
				banchRuleMux: new(sync.RWMutex),
			}
		}
	}
	return dbInstance
}

func(dbObj *DB) Conn(path string) {
	defer panichandler.Recover()
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("error loading .env file")
	}
	databaseIP := os.Getenv("DATA_BASE_IP")
	databasePort := os.Getenv("DATA_BASE_PORT")
	databaseName := os.Getenv("DATA_BASE_NAME")
	databaseUser := os.Getenv("DATA_BASE_USER")
	databasePassword := os.Getenv("DATA_BASE_PASSWORD")
	// fmt.Println(databaseIP, databasePort, databaseUser, databasePassword)
	dsn := databaseUser + ":" + databasePassword + "@tcp(" + databaseIP + ":" + databasePort +")/" + databaseName + "?" + "parseTime=true"
	(*dbObj).MysqlDB, err = sql.Open("mysql", dsn)

	// fmt.Println(dsn)
	if err != nil {
		log.Fatal(err)
	} else {
		(*dbObj).MysqlDB.SetMaxIdleConns(100000)
		(*dbObj).MysqlDB.SetMaxOpenConns(100000)
		(*dbObj).MysqlDB.SetConnMaxLifetime(time.Second * 100)
	}
	DataBaseInit();
}

//---------------------------select---------------------------------

// 0 => 全部, value => nil
//  1 =>  userId, value => int64
//  2 => account, value => string
// 3 => companyCode, value => string
func(dbObj *DB) SelectUser(selectKey int, value... interface{}) *[]table.UserTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).User.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).User.SelectSingleByUserId
		break
	case 2:
		// value need string
		querys = (*query.MysqlSingleton()).User.SelectSingleByAccount
		break
	case 3:
		// value need string
		querys = (*query.MysqlSingleton()).User.SelectAllByCompanyCode
		break
	default:
		querys = (*query.MysqlSingleton()).User.SelectAll
		break
	}
	user := new(table.UserTable)
	carry := []table.UserTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	companyCode := new(sql.NullString)
	for res.Next() {
		err = res.Scan(
			&user.UserId,
			companyCode,
			&user.UserName,
			&user.EmployeeNumber,
			&user.Account,
			&user.Password,
			&user.OnWorkDay,
			&user.Banch,
			&user.Permession,
			&user.WorkState,
			&user.CreateTime,
			&user.LastModify,
			&user.MonthSalary,
			&user.PartTimeSalary,
		)
		(*dbObj).checkErr(err)
		if companyCode.String == "" {
			user.CompanyCode = ""
		} else {
			user.CompanyCode = companyCode.String
		}
		
		if err == nil {
			carry = append(carry, *user)
		}
	}
	return &carry
}

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

// 0 => 全部, value => nil
//  1 => 公司id, value => int64
//  2 => 公司碼, value => string
func(dbObj *DB) SelectCompany(selectKey int, value... interface{}) *[]table.CompanyTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).Company.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).Company.SelectSingleByCompanyId
		break
	case 2:
		// value need string
		querys = (*query.MysqlSingleton()).Company.SelectSingleByCompanyCode
		break
	default:
		querys = (*query.MysqlSingleton()).Company.SelectAll
		break
	}
	company := new(table.CompanyTable)
	carry := []table.CompanyTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&company.CompanyId,
			&company.CompanyCode,
			&company.CompanyName,
			&company.CompanyLocation,
			&company.CompanyPhoneNumber,
			&company.TermStart,
			&company.TermEnd,
			&company.CreateTime,
			&company.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *company)
		}
	}

	return &carry
}

// 0 => 全部, value => nil
//	1 => 公司Id, value => int64
// 	2 => id (banchId), value => int64
func(dbObj *DB) SelectCompanyBanch(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectSingleByCompanyId
		break
	case 2:
		// value need int
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectSingleById
		break;
	default:
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectAll
		break
	}
	companyBanch := new(table.CompanyBanchTable)
	carry := []table.CompanyBanchTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&companyBanch.Id,
			&companyBanch.CompanyId,
			&companyBanch.BanchName,
			&companyBanch.BanchShiftStyle,
			&companyBanch.CreateTime,
			&companyBanch.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *companyBanch)
		}
	}

	return &carry 
}

// 0 => all, value => nil
//  1 => 班表id, value => int64
func(dbObj *DB) SelectShift(selectKey int, value... interface{}) *[]table.ShiftTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).Shift.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).Shift.SelectSingleByShiftId
		break
	default:
		querys = (*query.MysqlSingleton()).Shift.SelectAll
		break
	}
	shift := new(table.ShiftTable)
	carry := []table.ShiftTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&shift.ShiftId,
			&shift.UserId,
			&shift.OnShiftTime,
			&shift.OffShiftTime,
			&shift.PunchIn,
			&shift.PunchOut,
			&shift.SpecifyTag,
			&shift.CreateTime,
			&shift.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *shift)
		}
	}
	return &carry
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectShiftChange(selectKey int, value... interface{}) *[]table.ShiftChangeTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftChange.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).ShiftChange.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).ShiftChange.SelectAll
		break
	}
	shiftChange := new(table.ShiftChangeTable)
	carry := []table.ShiftChangeTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&shiftChange.CaseId,
			&shiftChange.InitiatorShiftId,
			&shiftChange.RequestedShiftId,
			&shiftChange.Reason,
			&shiftChange.CaseProcess,
			&shiftChange.SpecifyTag,
			&shiftChange.CreateTime,
			&shiftChange.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *shiftChange)
		}
	}
	return &carry
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectShiftOverTime(selectKey int, value... interface{}) *[]table.ShiftOverTimeTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectAll
		break
	}
	shiftOverTime := new(table.ShiftOverTimeTable)
	carry := []table.ShiftOverTimeTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&shiftOverTime.CaseId,
			&shiftOverTime.ShiftId,
			&shiftOverTime.InitiatorOnOverTime,
			&shiftOverTime.InitiatorOffOverTime,
			&shiftOverTime.Reason,
			&shiftOverTime.CaseProcess,
			&shiftOverTime.SpecifyTag,
			&shiftOverTime.CreateTime,
			&shiftOverTime.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *shiftOverTime)
		}
	}
	return &carry
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectForgetPunch(selectKey int, value... interface{}) *[]table.ForgetPunchTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectAll
		break
	}
	forgetPunch := new(table.ForgetPunchTable)
	carry := []table.ForgetPunchTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&forgetPunch.CaseId,
			&forgetPunch.ShiftId,
			&forgetPunch.TargetPunch,
			&forgetPunch.Reason,
			&forgetPunch.CaseProcess,
			&forgetPunch.SpecifyTag,
			&forgetPunch.CreateTime,
			&forgetPunch.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *forgetPunch)
		}
	}
	return &carry
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectDayOff(selectKey int, value... interface{}) *[]table.DayOffTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).DayOff.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).DayOff.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).DayOff.SelectAll
		break
	}
	dayOff := new(table.DayOffTable)
	carry := []table.DayOffTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&dayOff.CaseId,
			&dayOff.ShiftId,
			&dayOff.DayOffType,
			&dayOff.Reason,
			&dayOff.CaseProcess,
			&dayOff.SpecifyTag,
			&dayOff.CreateTime,
			&dayOff.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *dayOff)
		}
	}
	return &carry
}

// 0 => all, value => nil
//  1 => caseId, value => int64
func(dbObj *DB) SelectLateExcused(selectKey int, value... interface{}) *[]table.LateExcusedTable {
	defer panichandler.Recover()
	querys := ""
	switch selectKey {
	case 0:
		querys = (*query.MysqlSingleton()).LateExcused.SelectAll
		break
	case 1:
		// value need int
		querys = (*query.MysqlSingleton()).LateExcused.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).LateExcused.SelectAll
		break
	}
	lateExcused := new(table.LateExcusedTable)
	carry := []table.LateExcusedTable{}
	res, err := (*dbObj).MysqlDB.Query(querys, value...)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
			&lateExcused.CaseId,
			&lateExcused.ShiftId,
			&lateExcused.LateExcusedType,
			&lateExcused.Reason,
			&lateExcused.CaseProcess,
			&lateExcused.SpecifyTag,
			&lateExcused.CreateTime,
			&lateExcused.LastModify,
		)
		(*dbObj).checkErr(err)
		if err == nil {
			carry = append(carry, *lateExcused)
		}
	}
	return &carry
}

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

// ---------------------------------delete------------------------------------

//使用者的唯一id (關聯資料表userpreference 也上鎖)
func(dbObj *DB) DeleteUser(deleteKey int, userId interface{}) bool {
	defer panichandler.Recover()

	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()

	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()

	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).User.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(userId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
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

//公司的唯一id (關聯資料表 companyBanch 也上鎖)
func(dbObj *DB) DeleteCompany(deleteKey int, companyId interface{}) bool {
	defer panichandler.Recover()

	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()

	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()

	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Company.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(companyId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 公司部門的id
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).CompanyBanch.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(id)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 班表的唯一id (關聯資料表	shiftovertime shiftchange forgetpunch lateexcused dayoff 都上鎖)
func(dbObj *DB) DeleteShift(deleteKey int, shiftId interface{}) bool {
	defer panichandler.Recover()

	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()

	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()

	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()

	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()

	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()

	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()

	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Shift.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(shiftId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftChange.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftOverTime.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ForgetPunch.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).LateExcused.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteDayOff(deleteKey int, caseId interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).DayOff.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
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

//  ---------------------------------update------------------------------------


// 0 => all
func(dbObj *DB) UpdateUser(updateKey int, data *table.UserTable, value... interface{}) bool {
		defer panichandler.Recover()
		(*dbObj).userMux.Lock()
		defer (*dbObj).userMux.Unlock()
		defer func ()  {
			(*dbObj).containers.user = nil
		}()
		querys := ""
		switch updateKey {
		case 0:
			querys = (*query.MysqlSingleton()).User.UpdateSingle
			(*dbObj).containers.user = append(
				(*dbObj).containers.user,
				(*data).CompanyCode,
				(*data).Password,
				(*data).UserName,
				(*data).OnWorkDay,
				(*data).Banch,
				(*data).Permession,
				(*data).WorkState,
				(*data).LastModify,
				(*data).MonthSalary,
				(*data).PartTimeSalary,
				(*data).UserId,
			)
			break
		case 1:

		}
		
		stmt, err := (*dbObj).MysqlDB.Prepare(querys)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		_, err = stmt.Exec((*dbObj).containers.user...)
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

// 0 => all
func(dbObj *DB) UpdateCompany(updateKey int, data *table.CompanyTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	defer func ()  {
		(*dbObj).containers.company = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).Company.UpdateSingle
		(*dbObj).containers.company = append(
			(*dbObj).containers.company,
			(*data).CompanyName,
			(*data).CompanyLocation,
			(*data).CompanyPhoneNumber,
			(*data).TermStart,
			(*data).TermEnd,
			(*data).LastModify,
			(*data).CompanyId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.company...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateCompanyBanch(updateKey int, data *table.CompanyBanchTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	defer func ()  {
		(*dbObj).containers.companyBanch = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).CompanyBanch.UpdateSingle
		(*dbObj).containers.companyBanch = append(
			(*dbObj).containers.companyBanch,
			(*data).BanchName,
			(*data).BanchShiftStyle,
			(*data).LastModify,
			(*data).Id,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.companyBanch...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateShift(updateKey int, data *table.ShiftTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	defer func ()  {
		(*dbObj).containers.shift = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).Shift.UpdateSingle
		(*dbObj).containers.shift = append(
			(*dbObj).containers.shift,
			(*data).OnShiftTime,
			(*data).OffShiftTime,
			(*data).PunchIn,
			(*data).PunchOut,
			(*data).SpecifyTag,
			(*data).LastModify,
			(*data).ShiftId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.shift...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateShiftChange(updateKey int, data *table.ShiftChangeTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	defer func ()  {
		(*dbObj).containers.shiftChange = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftChange.UpdateSingle
		(*dbObj).containers.shiftChange = append(
			(*dbObj).containers.shiftChange,
			(*data).InitiatorShiftId,
			(*data).RequestedShiftId,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
			(*data).LastModify,
			(*data).CaseId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.shiftChange...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateShiftOverTime(updateKey int, data *table.ShiftOverTimeTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	defer func ()  {
		(*dbObj).containers.shiftOverTime = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).ShiftOverTime.UpdateSingle
		(*dbObj).containers.shiftOverTime = append(
			(*dbObj).containers.shiftOverTime,
			(*data).InitiatorOnOverTime,
			(*data).InitiatorOffOverTime,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
			(*data).LastModify,
			(*data).CaseId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.shiftOverTime...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateForgetPunch(updateKey int, data *table.ForgetPunchTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	defer func ()  {
		(*dbObj).containers.forgetPunch = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).ForgetPunch.UpdateSingle
		(*dbObj).containers.forgetPunch = append(
			(*dbObj).containers.forgetPunch,
			(*data).TargetPunch,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
			(*data).LastModify,
			(*data).CaseId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.forgetPunch...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateDayOff(updateKey int, data *table.DayOffTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	defer func ()  {
		(*dbObj).containers.dayOff = nil	
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).DayOff.UpdateSingle
		(*dbObj).containers.dayOff = append(
			(*dbObj).containers.dayOff,
			(*data).DayOffType,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
			(*data).LastModify,
			(*data).CaseId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.dayOff...)
	if err != nil {
		(*dbObj).checkErr(err)
		return false
	}
	return true
}

// 0 => all
func(dbObj *DB) UpdateLateExcused(updateKey int, data *table.LateExcusedTable, value ...interface{}) bool {
	defer panichandler.Recover()
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	defer func ()  {
		(*dbObj).containers.lateExcused = nil
	}()
	querys := ""
	switch updateKey {
	case 0:
		querys = (*query.MysqlSingleton()).LateExcused.UpdateSingle
		(*dbObj).containers.lateExcused = append(
			(*dbObj).containers.lateExcused,
			(*data).LateExcusedType,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
			(*data).LastModify,
			(*data).CaseId,
		)
		break;
	}
	
	stmt, err := (*dbObj).MysqlDB.Prepare(querys)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec((*dbObj).containers.lateExcused...)
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


//  ---------------------------------insert------------------------------------
func(dbObj *DB) InsertUser(data *table.UserTable) (bool, int64) {
	///
		defer panichandler.Recover()
		(*dbObj).userMux.Lock()
		defer (*dbObj).userMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).User.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).CompanyCode,
			(*data).Account,
			(*data).Password,
			(*data).UserName,
			(*data).EmployeeNumber,
			(*data).OnWorkDay,
			(*data).Banch,
			(*data).Permession,
			(*data).WorkState,
			(*data).CreateTime,
			(*data).LastModify,
			(*data).MonthSalary,
			(*data).PartTimeSalary,
		)
		(*dbObj).checkErr(err)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id

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
func(dbObj *DB) InsertCompany(data *table.CompanyTable) (bool, int64) {

		defer panichandler.Recover()
		(*dbObj).companyMux.Lock()
		defer (*dbObj).companyMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Company.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)

		res, err := stmt.Exec(
			(*data).CompanyCode,
			(*data).CompanyName,
			(*data).CompanyLocation,
			(*data).CompanyPhoneNumber,
			(*data).TermStart,
			(*data).TermEnd,
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
func(dbObj *DB) InsertCompanyBanch(data *table.CompanyBanchTable) (bool, int64) {

		defer panichandler.Recover()
		(*dbObj).companyBanchMux.Lock()
		defer (*dbObj).companyBanchMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).CompanyBanch.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).CompanyId,
			(*data).BanchName,
			(*data).BanchShiftStyle,
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
func(dbObj *DB) InsertShift(data *table.ShiftTable) (bool, int64) {

		defer panichandler.Recover()
		(*dbObj).shiftMux.Lock()
		defer (*dbObj).shiftMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Shift.InsertAll)
		defer stmt.Close()

		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).UserId,
			(*data).OnShiftTime,
			(*data).OffShiftTime,
			(*data).PunchIn,
			(*data).PunchOut,
			(*data).SpecifyTag,
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
func(dbObj *DB) InsertShiftChange(data *table.ShiftChangeTable) (bool, int64) {

		defer panichandler.Recover()
		(*dbObj).shiftChangeMux.Lock()
		defer (*dbObj).shiftChangeMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftChange.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).InitiatorShiftId,
			(*data).RequestedShiftId,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
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

func(dbObj *DB) InsertShiftOverTime(data *table.ShiftOverTimeTable) (bool, int64) {

		defer panichandler.Recover()
		(*dbObj).shiftOverTimeMux.Lock()
		defer (*dbObj).shiftOverTimeMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftOverTime.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).ShiftId,
			(*data).InitiatorOnOverTime,
			(*data).InitiatorOffOverTime,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
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

func(dbObj *DB) InsertForgetPunch(data *table.ForgetPunchTable) (bool, int64) {

		defer panichandler.Recover()
		(*dbObj).forgetPunchMux.Lock()
		defer (*dbObj).forgetPunchMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ForgetPunch.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).ShiftId,
			(*data).TargetPunch,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
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
func(dbObj *DB) InsertDayOff(data *table.DayOffTable) (bool, int64) {

		defer panichandler.Recover()
		(*dbObj).dayOffMux.Lock()
		defer (*dbObj).dayOffMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).DayOff.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).ShiftId,
			(*data).DayOffType,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
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
func(dbObj *DB) InsertLateExcused(data *table.LateExcusedTable) (bool, int64) {

		defer panichandler.Recover()
		(*dbObj).lateExcusedMux.Lock()
		defer (*dbObj).lateExcusedMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).LateExcused.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			(*data).ShiftId,
			(*data).LateExcusedType,
			(*data).Reason,
			(*data).CaseProcess,
			(*data).SpecifyTag,
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
func(dbObj *DB) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}