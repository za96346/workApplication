package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	// "runtime"
	"sync"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"backend/table"
	"backend/query"
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
	forgetPunchMux *sync.RWMutex
	MysqlDB *sql.DB // 要先使用連線方法後才能使用這個
	resStatus bool
}

type dbInterface interface {
	Conn()
	SelectUserSingle() *table.UserTable
	SelectUserAll()
	SelectCompanySingle() *table.CompanyTable
	InsertUserAll() bool
	InsertUserPreferenceAll() bool
	InsertCompanyAll() bool
	InsertCompanyBanchAll() bool
	InsertShiftAll() bool
}

func Singleton() *DB {
	if dbInstance == nil {
		dbSingletonMux.Lock()
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
			}
			defer dbSingletonMux.Unlock()
		}
	}
	return dbInstance
}

func(dbObj *DB) Conn() {
	err := godotenv.Load()
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

// 0 => all
func(dbObj *DB) SelectUserAll(selectKey int, value... interface{}) *[]table.UserTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).User.SelectAll
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
	for res.Next() {
		err = res.Scan(
			&user.UserId,
			&user.CompanyCode,
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
		if err == nil {
			carry = append(carry, *user)
		}
	}
	return &carry
}

// 0 => 使用者id int (default)
//  1 => 帳號 string
func(dbObj *DB) SelectUserSingle(selectKey int, value... interface{}) *table.UserTable {
	user := new(table.UserTable)
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).User.SelectSingleByUserId
		break
	case 1:
		// value need string
		querys = (*query.MysqlSingleton()).User.SelectSingleByAccount
		break
	default:
		querys = (*query.MysqlSingleton()).User.SelectSingleByUserId
		break;
	}
	err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
		&user.UserId,
		&user.CompanyCode,
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
	return user
}

// 0 => all
func(dbObj *DB) SelectUserPreferenceAll(selectKey int, value... interface{}) *[]table.UserPreferenceTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).UserPreference.SelectAll
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

// 0 => 使用者id (default)
func(dbObj *DB) SelectUserPreferenceSingle(selectKey int, value... interface{}) *table.UserPreferenceTable {
	userPreference := new(table.UserPreferenceTable)
	// defer runtime.GC()
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).UserPreference.SelectSingleByUserId
		break
	default:
		querys = (*query.MysqlSingleton()).UserPreference.SelectSingleByUserId
		break
	}
	err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
		&userPreference.UserId,
		&userPreference.Style,
		&userPreference.FontSize,
		&userPreference.SelfPhoto,
		&userPreference.CreateTime,
		&userPreference.LastModify,
	)
	(*dbObj).checkErr(err)
	return userPreference
}

// 0 => all
func(dbObj *DB) SelectCompanyAll(selectKey int, value... interface{}) *[]table.CompanyTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).Company.SelectAll
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

//0 => 公司id (default)
// 1 => 公司碼
func(dbObj *DB) SelectCompanySingle(selectKey int, value... interface{}) *table.CompanyTable {
		company := new(table.CompanyTable)
		querys := ""
		switch selectKey {
		case 0:
			// value need int
			querys = (*query.MysqlSingleton()).Company.SelectSingleByCompanyId
			break
		case 1:
			// value need string
			querys = (*query.MysqlSingleton()).Company.SelectSingleByCompanyCode
			break
		default:
			querys = (*query.MysqlSingleton()).Company.SelectSingleByCompanyId
			break
		}
		err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
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
		return company
}

// 0 => all
func(dbObj *DB) SelectCompanyBanchAll(selectKey int, value... interface{}) *[]table.CompanyBanchTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).CompanyBanch.SelectAll
		break
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

//0 => companyId (default)
// 1 => id (banchId)
func(dbObj *DB) SelectCompanyBanchSingle(selectKey int, value... interface{}) *table.CompanyBanchTable {

		companyBanch := new(table.CompanyBanchTable)
		querys := ""
		switch selectKey {
		case 0:
			// value need int
			querys = (*query.MysqlSingleton()).CompanyBanch.SelectSingleByCompanyId
			break
		case 1:
			querys = (*query.MysqlSingleton()).CompanyBanch.SelectSingleById
			break;
		default:
			querys = (*query.MysqlSingleton()).CompanyBanch.SelectSingleByCompanyId
			break
		}
		err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
			&companyBanch.Id,
			&companyBanch.CompanyId,
			&companyBanch.BanchName,
			&companyBanch.BanchShiftStyle,
			&companyBanch.CreateTime,
			&companyBanch.LastModify,
		)
		(*dbObj).checkErr(err)
		return companyBanch
}

// 0 => all
func(dbObj *DB) SelectShiftAll(selectKey int, value... interface{}) *[]table.ShiftTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).Shift.SelectAll
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

// 0 => 班表id
func(dbObj *DB) SelectShiftSingle(selectKey int, value... interface{}) *table.ShiftTable {
	shift := new(table.ShiftTable)
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).Shift.SelectSingleByShiftId
		break
	default:
		querys = (*query.MysqlSingleton()).Shift.SelectSingleByShiftId
		break
	}
	err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
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
	return shift
}

// 0 => all
func(dbObj *DB) SelectShiftChangeAll(selectKey int, value... interface{}) *[]table.ShiftChangeTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).ShiftChange.SelectAll
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

// 0 => caseId (default)
func(dbObj *DB) SelectShiftChangeSingle(selectKey int, value... interface{}) *table.ShiftChangeTable {
	shiftChange := new(table.ShiftChangeTable)
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).ShiftChange.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).ShiftChange.SelectSingleByCaseId
		break
	}
	err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
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
	return shiftChange
}

// 0 => all
func(dbObj *DB) SelectShiftOverTimeAll(selectKey int, value... interface{}) *[]table.ShiftOverTimeTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectAll
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

// 0 => caseId (default)
func(dbObj *DB) SelectShiftOverTimeSingle(selectKey int, value... interface{}) *table.ShiftOverTimeTable {
	shiftOverTime:= new(table.ShiftOverTimeTable)
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).ShiftOverTime.SelectSingleByCaseId
		break
	}
	err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
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
	return shiftOverTime
}

// 0 => all
func(dbObj *DB) SelectForgetPunchAll(selectKey int, value... interface{}) *[]table.ForgetPunchTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectAll
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

// 0 => caseId (default)
func(dbObj *DB) SelectForgetPunchSingle(selectKey int, value... interface{}) *table.ForgetPunchTable {
	forgetPunch := new(table.ForgetPunchTable)
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).ForgetPunch.SelectSingleByCaseId
		break
	}
	err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
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
	return forgetPunch
}

// 0 => all
func(dbObj *DB) SelectDayOffAll(selectKey int, value... interface{}) *[]table.DayOffTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).DayOff.SelectAll
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

// 0 => caseId (default)
func(dbObj *DB) SelectDayOffSingle(selectKey int, value... interface{}) *table.DayOffTable {
	dayOff := new(table.DayOffTable)
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).DayOff.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).DayOff.SelectSingleByCaseId
		break
	}
	err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
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
	return dayOff
}

// 0 => all
func(dbObj *DB) SelectLateExcusedAll(selectKey int, value... interface{}) *[]table.LateExcusedTable {
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).LateExcused.SelectAll
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

// 0 => caseId (default)
func(dbObj *DB) SelectLateExcusedSingle(selectKey int, value... interface{}) *table.LateExcusedTable {
	lateExcused := new(table.LateExcusedTable)
	querys := ""
	switch selectKey {
	case 0:
		// value need int
		querys = (*query.MysqlSingleton()).LateExcused.SelectSingleByCaseId
		break
	default:
		querys = (*query.MysqlSingleton()).LateExcused.SelectSingleByCaseId
		break
	}
	err := (*dbObj).MysqlDB.QueryRow(querys, value...).Scan(
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
	return lateExcused
}


// ---------------------------------delete------------------------------------

//使用者的唯一id
func(dbObj *DB) DeleteUser(deleteKey int, userId interface{}) bool {
	(*dbObj).userMux.Lock()
	defer (*dbObj).userMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).User.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(userId)
	if err != nil {
		return false
	}
	return true
}

//使用者的唯一id
func(dbObj *DB) DeleteUserPreference(deleteKey int, userId interface{}) bool {
	(*dbObj).userPreferenceMux.Lock()
	defer (*dbObj).userPreferenceMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).UserPreference.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(userId)
	if err != nil {
		return false
	}
	return true
}

//公司的唯一id
func(dbObj *DB) DeleteCompany(deleteKey int, companyId interface{}) bool {
	(*dbObj).companyMux.Lock()
	defer (*dbObj).companyMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Company.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(companyId)
	if err != nil {
		return false
	}
	return true
}

// 公司部門的id
func(dbObj *DB) DeleteCompanyBanch(deleteKey int, id interface{}) bool {
	(*dbObj).companyBanchMux.Lock()
	defer (*dbObj).companyBanchMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).CompanyBanch.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(id)
	if err != nil {
		return false
	}
	return true
}

// 班表的唯一id
func(dbObj *DB) DeleteShift(deleteKey int, shiftId interface{}) bool {
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Shift.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(shiftId)
	if err != nil {
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftChange(deleteKey int, caseId interface{}) bool {
	(*dbObj).shiftChangeMux.Lock()
	defer (*dbObj).shiftChangeMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftChange.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteShiftOverTime(deleteKey int, caseId interface{}) bool {
	(*dbObj).shiftOverTimeMux.Lock()
	defer (*dbObj).shiftOverTimeMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftOverTime.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteForgetPunch(deleteKey int, caseId interface{}) bool {
	(*dbObj).forgetPunchMux.Lock()
	defer (*dbObj).forgetPunchMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ForgetPunch.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteLateExcused(deleteKey int, caseId interface{}) bool {
	(*dbObj).lateExcusedMux.Lock()
	defer (*dbObj).lateExcusedMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).LateExcused.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		return false
	}
	return true
}

// 案件的唯一id
func(dbObj *DB) DeleteDayOff(deleteKey int, caseId interface{}) bool {
	(*dbObj).dayOffMux.Lock()
	defer (*dbObj).dayOffMux.Unlock()
	stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).DayOff.Delete)
	defer stmt.Close()
	(*dbObj).checkErr(err)
	_, err = stmt.Exec(caseId)
	if err != nil {
		return false
	}
	return true
}


//  ---------------------------------insert------------------------------------
func(dbObj *DB) InsertUser(
	companyCode string,
	account string,
	password string,
	onWorkDay time.Time,
	banch string,
	permession string,
	workState string,
	createTime time.Time,
	lastModify time.Time,
	monthSalary int,
	partTimeSalary int) (bool, int64) {
	///
		(*dbObj).userMux.Lock()
		defer (*dbObj).userMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).User.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			companyCode,
			account,
			password,
			onWorkDay,
			banch,
			permession,
			workState,
			createTime,
			lastModify,
			monthSalary,
			partTimeSalary,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id

}
func(dbObj *DB) InsertUserPreference(
	userId int64,
	style string,
	fontSize string,
	selfPhoto string,
	createTime time.Time,
	lastModify time.Time,
	) (bool, int64) {
		///
		(*dbObj).userPreferenceMux.Lock()
		defer  (*dbObj).userPreferenceMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).UserPreference.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			userId,
			style,
			fontSize,
			selfPhoto,
			createTime,
			lastModify,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id
}
func(dbObj *DB) InsertCompany(
	companyCode string,
	companyName string,
	companyLocation string,
	companyPhoneNumber string,
	termStart time.Time,
	termEnd time.Time,
	createTime time.Time,
	lastModify time.Time) (bool, int64) {

		(*dbObj).companyMux.Lock()
		defer (*dbObj).companyMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Company.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)

		res, err := stmt.Exec(
			companyCode,
			companyName,
			companyLocation,
			companyPhoneNumber,
			termStart,
			termEnd,
			createTime,
			lastModify,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id
}
func(dbObj *DB) InsertCompanyBanch(
	companyId int64,
	banchName string,
	banchShiftStyle string,
	createTime time.Time,
	lastModify time.Time,
	) (bool, int64) {

		(*dbObj).companyBanchMux.Lock()
		defer (*dbObj).companyBanchMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).CompanyBanch.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			companyId,
			banchName,
			banchShiftStyle,
			createTime,
			lastModify,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id
}
func(dbObj *DB) InsertShift(
	userId int64,
	onShiftTime time.Time,
	offShiftTime time.Time,
	punchIn time.Time,
	punchOut time.Time,
	createTime time.Time,
	lastModify time.Time,
	specifyTag string,
	) (bool, int64) {

		(*dbObj).shiftMux.Lock()
		defer (*dbObj).shiftMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).Shift.InsertAll)
		defer stmt.Close()

		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			userId,
			onShiftTime,
			offShiftTime,
			punchIn,
			punchOut,
			specifyTag,
			createTime,
			lastModify,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id
}
func(dbObj *DB) InsertShiftChange(
	initiatorShiftId int64,
	requestedShiftId int64,
	reson string,
	caseProcess string,
	specifyTag string,
	createTime time.Time,
	lastModify time.Time,
	) (bool, int64) {

		(*dbObj).shiftChangeMux.Lock()
		defer (*dbObj).shiftChangeMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftChange.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			initiatorShiftId,
			requestedShiftId,
			reson,
			caseProcess,
			specifyTag,
			createTime,
			lastModify,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id
}

func(dbObj *DB) InsertShiftOverTime(
	shiftId int64,
	initiatorOnOverTime time.Time,
	initiatorOffOverTime time.Time,
	reson string,
	caseProcess string,
	specifyTag string,
	createTime time.Time,
	lastModify time.Time,
	) (bool, int64) {
		(*dbObj).shiftOverTimeMux.Lock()
		defer (*dbObj).shiftOverTimeMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ShiftOverTime.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			shiftId,
			initiatorOnOverTime,
			initiatorOffOverTime,
			reson,
			caseProcess,
			specifyTag,
			createTime,
			lastModify,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id
}

func(dbObj *DB) InsertForgetPunch(
	shiftId int64,
	targetPunch string,
	reson string,
	caseProcess string,
	specifyTag string,
	createTime time.Time,
	lastModify time.Time,
	) (bool, int64) {
		(*dbObj).forgetPunchMux.Lock()
		defer (*dbObj).forgetPunchMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).ForgetPunch.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			shiftId,
			targetPunch,
			reson,
			caseProcess,
			specifyTag,
			createTime,
			lastModify,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id
}
func(dbObj *DB) InsertDayOff(
	shiftId int64,
	dayOffType string,
	reson string,
	caseProcess string,
	specifyTag string,
	createTime time.Time,
	lastModify time.Time,
	) (bool, int64) {
		(*dbObj).dayOffMux.Lock()
		defer (*dbObj).dayOffMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).DayOff.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			shiftId,
			dayOffType,
			reson,
			caseProcess,
			specifyTag,
			createTime,
			lastModify,
		)
		id, _:= res.LastInsertId()
		if err != nil {
			return false, id
		}
		return true, id
}
func(dbObj *DB) InsertLateExcused(
	shiftId int64,
	lateExcusedType string,
	reson string,
	caseProcess string,
	specifyTag string,
	createTime time.Time,
	lastModify time.Time,
	) (bool, int64) {
		(*dbObj).lateExcusedMux.Lock()
		defer (*dbObj).lateExcusedMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*query.MysqlSingleton()).LateExcused.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		res, err := stmt.Exec(
			shiftId,
			lateExcusedType,
			reson,
			caseProcess,
			specifyTag,
			createTime,
			lastModify,
		)
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