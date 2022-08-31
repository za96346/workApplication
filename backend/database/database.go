package database

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (
	// "fmt"
	"log"
	"os"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var MysqlDB *sql.DB
var err error

func MysqlDBConn() {
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
		MysqlDB.SetMaxIdleConns(100000)
		MysqlDB.SetMaxOpenConns(100000)
		MysqlDB.SetConnMaxLifetime(time.Second * 100)
	}
	DataBaseInit();
}

func SelectSingleUser() {
	// err = DB.QueryRow("select * from user where account = ?", "a00002").Scan(&account, &name, &password, &banch, &emp_id, &on_work_day, &position, &work_state)
	rows, err := MysqlDB.Query("select * from user");
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	defer MysqlDB.Close()

	for rows.Next() {
		// err := rows.Scan(
		// 	&userTableInstance.Account,
		// 	&userTableInstance.Banch,
		// 	&userTableInstance.Emp_id,
		// 	&userTableInstance.Name,
		// 	&userTableInstance.On_work_day,
		// 	&userTableInstance.Password,
		// 	&userTableInstance.Position,
		// 	&userTableInstance.Work_state)
		if err != nil {
			log.Fatal(err)
		}
		log.Println()
	}
}