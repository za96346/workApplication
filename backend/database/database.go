package database

//$go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
import (
	// "fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	account string
	password string
	name string
	banch string
	emp_id string
	on_work_day string
	position string
	work_state string
)
var DB *sql.DB
var err error

func DBInit() {
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
	DB, err = sql.Open("mysql", dsn)

	// fmt.Println(dsn)
	if err != nil {
		log.Fatal(err)
	} else {
		DB.SetMaxIdleConns(100000)
		DB.SetMaxOpenConns(100000)
	}
	SelectSingleUser();
}

func SelectSingleUser() {
	// err = DB.QueryRow("select * from user where account = ?", "a00002").Scan(&account, &name, &password, &banch, &emp_id, &on_work_day, &position, &work_state)
	rows, err := DB.Query("select * from user");
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	defer DB.Close()
	for rows.Next() {
		err := rows.Scan(&account, &name, &password, &banch, &emp_id, &on_work_day, &position, &work_state)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(account, name, password, banch, emp_id, on_work_day, position, work_state)
	}
}