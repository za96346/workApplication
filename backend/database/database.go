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

var MysqlDB *sql.DB
var err error

var selectUserSingleMux = new(sync.Mutex)
var InsertCompanyMux = new(sync.Mutex)
var InsertUserPreferenceAllMux = new(sync.RWMutex)
var InsertUserMux = new(sync.RWMutex)
var InsertCompanyBanchMux = new(sync.RWMutex)

var DBSingletonMux = new(sync.Mutex)
var resStatus bool

var DBInstance *db

type db struct {
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
}

func DBSingleton() *db {
	if DBInstance == nil {
		DBSingletonMux.Lock()
		if DBInstance == nil {
			DBInstance = &db{}
			defer DBSingletonMux.Unlock()
		}
	}
	return DBInstance
}

func(db *db)Conn() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	databaseIP := os.Getenv("DATA_BASE_IP")
	databasePort := os.Getenv("DATA_BASE_PORT")
	databaseName := os.Getenv("DATA_BASE_NAME")
	databaseUser := os.Getenv("DATA_BASE_USER")
	databasePassword := os.Getenv("DATA_BASE_PASSWORD")
	// fmt.Println(databaseIP, databasePort, databaseUser, databasePassword)
	dsn := databaseUser + ":" + databasePassword + "@tcp(" + databaseIP + ":" + databasePort +")/" + databaseName
	MysqlDB, err = sql.Open("mysql", dsn)

	// fmt.Println(dsn)
	if err != nil {
		log.Fatal(err)
	} else {
		// MysqlDB.SetMaxIdleConns(100000)
		// MysqlDB.SetMaxOpenConns(100000)
		MysqlDB.SetConnMaxLifetime(time.Second * 100)
	}
	DataBaseInit();
	// SelectUserAll();
}

func(db *db) SelectUserSingle(accountOrUserId interface{}) **UserTable {
	user := new(UserTable)
	// defer runtime.GC()
	err := MysqlDB.QueryRow((*SqlQuerySingleton()).User.selectSingle, accountOrUserId, accountOrUserId).Scan(
		&user.UserId,
		&user.CompanyCode,
		&user.Account,
		&user.Password,
		&user.OnWorkDay,
		&user.Banch,
		&user.Permession,
		&user.Work_state,
		&user.CreateTime,
		&user.LastModify,
		&user.MonthSalary,
		&user.PartTimeSalary,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("回傳SelectUserSingle 記憶體位置 => ", &user)
	return &user
}

func(db *db) SelectUserAll() *[]UserTable {
	user := new(UserTable)
	carry := []UserTable{}
	res, err := MysqlDB.Query((*SqlQuerySingleton()).User.selectAll)
	defer res.Close()
	// defer runtime.GC()
	if err != nil {
		fmt.Println(err)
		// return result
	}
	for res.Next() {
		err = res.Scan(
			&user.UserId,
			&user.CompanyCode,
			&user.Account,
			&user.Password,
			&user.OnWorkDay,
			&user.Banch,
			&user.Permession,
			&user.Work_state,
			&user.CreateTime,
			&user.LastModify,
			&user.MonthSalary,
			&user.PartTimeSalary,
		)
		if err != nil {
			log.Fatal(err)
		} else {
			carry = append(carry, *user)
		}
		log.Println("user -> ", *user)
	}
	// fmt.Println("==>", user, &user, *user)
	fmt.Println("回傳SelectUseAll 記憶體位置 => ", &carry)
	return &carry
}
func(db *db) SelectCompanySingle(companyIdOrCompanyCode interface{}) **CompanyTable {
	company := new(CompanyTable)
	err = MysqlDB.QueryRow((*SqlQuerySingleton()).Company.SelectSingle, companyIdOrCompanyCode, companyIdOrCompanyCode).Scan(
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
	// defer runtime.GC()
	if err != nil {
		fmt.Println(err)
		// return result
	}
	fmt.Println("回傳SelectCompanySingle 記憶體位置 => ", &company)
	return &company
}
func(db *db) InsertUserAll(
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
		InsertUserMux.Lock()
		defer InsertUserMux.Unlock()
		stmt, err := MysqlDB.Prepare((*SqlQuerySingleton()).User.InsertAll)
		defer stmt.Close()
		if err != nil {
			fmt.Println(err)
			return false
		}
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
		
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true

}
func(db *db) InsertUserPreferenceAll(
	userId float64,
	style string,
	fontSize string,
	selfPhoto string,
	createTime time.Time,
	lastModify time.Time,
	) bool {
		///
		InsertUserPreferenceAllMux.Lock()
		defer  InsertUserPreferenceAllMux.Unlock()
		stmt, err := MysqlDB.Prepare((*SqlQuerySingleton()).UserPreference.InsertAll)
		defer stmt.Close()
		if err != nil {
			fmt.Println(err)
			return false
		}
		_, err = stmt.Exec(
			userId,
			style,
			fontSize,
			selfPhoto,
			createTime,
			lastModify,
		)
		
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true;
}
func(db *db) InsertCompanyAll(
	companyCode string,
	companyName string,
	companyLocation string,
	companyPhoneNumber string,
	termStart time.Time,
	termEnd time.Time,
	createTime time.Time,
	lastModify time.Time) bool {
		InsertCompanyMux.Lock()
		defer InsertCompanyMux.Unlock()
		stmt, err := MysqlDB.Prepare((*SqlQuerySingleton()).Company.InsertAll)
		defer stmt.Close()
		if err != nil {
			fmt.Println(err)
			return false
		}

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
		
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
}
func(db *db) InsertCompanyBanchAll(
	companyId float64,
	banchName string,
	banchShiftStyle string,
	createTime time.Time,
	lastModify time.Time,
	) bool {
		InsertCompanyBanchMux.Lock()
		defer InsertCompanyBanchMux.Unlock()
		stmt, err := MysqlDB.Prepare((*SqlQuerySingleton()).CompanyBanch.InsertAll)
		defer stmt.Close()
		if err != nil {
			fmt.Println(err)
			return false
		}
		_, err = stmt.Exec(
			companyId,
			banchName,
			banchShiftStyle,
			createTime,
			lastModify,
		)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
}