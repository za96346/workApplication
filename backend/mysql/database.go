package mysql

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (
	"database/sql"
	"fmt"
	"log"
	"os"
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
	quitWorkUserMux *sync.RWMutex
	waitCompanyReply *sync.RWMutex
	weekendSetting *sync.RWMutex
	workTime *sync.RWMutex
	paidVocation *sync.RWMutex
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
	quitWorkUser []interface{}
	waitCompanyReply []interface{}
	weekendSetting []interface{}
	worktime []interface{}
	paidVocation []interface{}
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
				quitWorkUserMux: new(sync.RWMutex),
				waitCompanyReply: new(sync.RWMutex),
				weekendSetting: new(sync.RWMutex),
				workTime: new(sync.RWMutex),
				paidVocation: new(sync.RWMutex),
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
	// DataBaseInit();
}

func(dbObj *DB) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}