package database

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
)

var DBSingletonMux = new(sync.Mutex)
var DBInstance *db

type db struct {
	InsertCompanyMux *sync.Mutex
	InsertUserPreferenceAllMux *sync.RWMutex
	InsertUserMux *sync.RWMutex
	InsertCompanyBanchMux *sync.RWMutex
	InsertShiftMux *sync.RWMutex
	InsertShiftChangeMux *sync.RWMutex
	MysqlDB *sql.DB
	resStatus bool
}

type dbInterface interface {
	Conn()
	SelectUserSingle() *UserTable
	SelectUserAll()
	SelectCompanySingle() *CompanyTable
	InsertUserAll() bool
	InsertUserPreferenceAll() bool
	InsertCompanyAll() bool
	InsertCompanyBanchAll() bool
	InsertShiftAll() bool
}

func MysqlSingleton() *db {
	if DBInstance == nil {
		DBSingletonMux.Lock()
		if DBInstance == nil {
			DBInstance = &db{
				InsertCompanyMux: new(sync.Mutex),
				InsertUserPreferenceAllMux: new(sync.RWMutex),
				InsertUserMux: new(sync.RWMutex),
				InsertCompanyBanchMux: new(sync.RWMutex),
				InsertShiftMux: new(sync.RWMutex),
				InsertShiftChangeMux: new(sync.RWMutex),
			}
			defer DBSingletonMux.Unlock()
		}
	}
	return DBInstance
}

func(dbObj *db) Conn() {
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
		// MysqlDB.SetMaxIdleConns(100000)
		// MysqlDB.SetMaxOpenConns(100000)
		(*dbObj).MysqlDB.SetConnMaxLifetime(time.Second * 100)
	}
	DataBaseInit();
	// SelectUserAll();
}

//---------------------------select---------------------------------

func(dbObj *db) SelectUserAll() *[]UserTable {
	user := new(UserTable)
	carry := []UserTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).User.SelectAll)
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
func(dbObj *db) SelectUserSingle(accountOrUserId interface{}) **UserTable {
	user := new(UserTable)
	// defer runtime.GC()
	err := (*dbObj).MysqlDB.QueryRow((*SqlQuerySingleton()).User.SelectSingle, accountOrUserId, accountOrUserId).Scan(
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
	return &user
}
func(dbObj *db) SelectUserPreferenceAll() *[]UserPreferenceTable {
	userPreference := new(UserPreferenceTable)
	carry := []UserPreferenceTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).UserPreference.SelectAll)
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


func(dbObj *db) SelectCompanyAll() *[]CompanyTable {
	company := new(CompanyTable)
	carry := []CompanyTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).Company.SelectAll)
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

func(dbObj *db) SelectCompanySingle(companyIdOrCompanyCode interface{}) **CompanyTable {
	company := new(CompanyTable)
	err := (*dbObj).MysqlDB.QueryRow((*SqlQuerySingleton()).Company.SelectSingle, companyIdOrCompanyCode, companyIdOrCompanyCode).Scan(
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
	fmt.Println("回傳SelectCompanySingle 記憶體位置 => ", &company)
	return &company
}

func(dbObj *db) SelectCompanyBanchAll() *[]CompanyBanchTable {
	companyBanch := new(CompanyBanchTable)
	carry := []CompanyBanchTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).CompanyBanch.SelectAll)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
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

func(dbObj *db) SelectShiftAll() *[]ShiftTable {
	shift := new(ShiftTable)
	carry := []ShiftTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).Shift.SelectAll)
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
func(dbObj *db) SelectShiftSingleByUserId(userId int64) **ShiftTable {
	shift := new(ShiftTable)
	err := (*dbObj).MysqlDB.QueryRow((*SqlQuerySingleton()).Shift.SelectSingleByUserId, userId).Scan(
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
	return &shift
}
func(dbObj *db) SelectShiftSingleByShiftId(shiftId int64) **ShiftTable {
	shift := new(ShiftTable)
	err := (*dbObj).MysqlDB.QueryRow((*SqlQuerySingleton()).Shift.SelectSingleByShiftId, shiftId).Scan(
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
	return &shift
}

func(dbObj *db) SelectShiftChangeAll() *[]ShiftChangeTable {
	shiftChange := new(ShiftChangeTable)
	carry := []ShiftChangeTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).ShiftChange.SelectAll)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
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

func(dbObj *db) SelectShiftOverTimeAll() *[]ShiftOverTimeTable {
	shiftOverTime := new(ShiftOverTimeTable)
	carry := []ShiftOverTimeTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).ShiftOverTime.SelectAll)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
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

func(dbObj *db) SelectForgetPunchAll() *[]ForgetPunchTable {
	forgetPunch := new(ForgetPunchTable)
	carry := []ForgetPunchTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).ForgetPunch.SelectAll)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
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

func(dbObj *db) SelectDayOffAll() *[]DayOffTable {
	dayOff := new(DayOffTable)
	carry := []DayOffTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).DayOff.SelectAll)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
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
func(dbObj *db) SelectLateExcusedAll() *[]LateExcusedTable {
	lateExcused := new(LateExcusedTable)
	carry := []LateExcusedTable{}
	res, err := (*dbObj).MysqlDB.Query((*SqlQuerySingleton()).LateExcused.SelectAll)
	defer res.Close()
	(*dbObj).checkErr(err)
	for res.Next() {
		err = res.Scan(
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



//  ---------------------------------insert------------------------------------
func(dbObj *db) InsertUserAll(
	companyCode string,
	account string,
	password string,
	onWorkDay string,
	banch string,
	permession string,
	workState string,
	createTime time.Time,
	lastModify time.Time,
	monthSalary int,
	partTimeSalary int)bool {
	///
		(*dbObj).InsertUserMux.Lock()
		defer (*dbObj).InsertUserMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*SqlQuerySingleton()).User.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		_, err = stmt.Exec(
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
		
		(*dbObj).checkErr(err)
		return true

}
func(dbObj *db) InsertUserPreferenceAll(
	userId int64,
	style string,
	fontSize string,
	selfPhoto string,
	createTime time.Time,
	lastModify time.Time,
	) bool {
		///
		(*dbObj).InsertUserPreferenceAllMux.Lock()
		defer  (*dbObj).InsertUserPreferenceAllMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*SqlQuerySingleton()).UserPreference.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		_, err = stmt.Exec(
			userId,
			style,
			fontSize,
			selfPhoto,
			createTime,
			lastModify,
		)
		
		(*dbObj).checkErr(err)
		return true;
}
func(dbObj *db) InsertCompanyAll(
	companyCode string,
	companyName string,
	companyLocation string,
	companyPhoneNumber string,
	termStart time.Time,
	termEnd time.Time,
	createTime time.Time,
	lastModify time.Time) bool {

		(*dbObj).InsertCompanyMux.Lock()
		defer (*dbObj).InsertCompanyMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*SqlQuerySingleton()).Company.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)

		_, err = stmt.Exec(
			companyCode,
			companyName,
			companyLocation,
			companyPhoneNumber,
			termStart,
			termEnd,
			createTime,
			lastModify,
		)
		
		(*dbObj).checkErr(err)
		return true
}
func(dbObj *db) InsertCompanyBanchAll(
	companyId int64,
	banchName string,
	banchShiftStyle string,
	createTime time.Time,
	lastModify time.Time,
	) bool {

		(*dbObj).InsertCompanyBanchMux.Lock()
		defer (*dbObj).InsertCompanyBanchMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*SqlQuerySingleton()).CompanyBanch.InsertAll)
		defer stmt.Close()
		(*dbObj).checkErr(err)
		_, err = stmt.Exec(
			companyId,
			banchName,
			banchShiftStyle,
			createTime,
			lastModify,
		)
		(*dbObj).checkErr(err)
		return true
}
func(dbObj *db) InsertShiftAll(
	userId int64,
	onShiftTime time.Time,
	offShiftTime time.Time,
	punchIn time.Time,
	punchOut time.Time,
	createTime time.Time,
	lastModify time.Time,
	specifyTag string,
	) bool {

		(*dbObj).InsertShiftMux.Lock()
		defer (*dbObj).InsertShiftMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*SqlQuerySingleton()).Shift.InsertAll)
		defer stmt.Close()

		(*dbObj).checkErr(err)
		_, err = stmt.Exec(
			userId,
			onShiftTime,
			offShiftTime,
			punchIn,
			punchOut,
			specifyTag,
			createTime,
			lastModify,
		)
		(*dbObj).checkErr(err)
		return true
}
func(dbObj *db) InsertShiftChangeAll(
	initiatorShiftId int64,
	requestedShiftId int64,
	reson string,
	caseProcess string,
	specifyTag string,
	createTime time.Time,
	lastModify time.Time,
	) bool {

		(*dbObj).InsertShiftChangeMux.Lock()
		defer (*dbObj).InsertShiftChangeMux.Unlock()
		stmt, err := (*dbObj).MysqlDB.Prepare((*SqlQuerySingleton()).ShiftChange.InsertAll)
		(*dbObj).checkErr(err)
		_, err = stmt.Exec(
			initiatorShiftId,
			requestedShiftId,
			reson,
			caseProcess,
			specifyTag,
			createTime,
			lastModify,
		)
		(*dbObj).checkErr(err)

		return true
}
func(dbObj *db) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}