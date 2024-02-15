package Model

import (
	"fmt"
	"strconv"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 設定 model
func SetUp() {
	//配置MySQL连接参数
	username := os.Getenv("DATA_BASE_USER")  //账号
	password := os.Getenv("DATA_BASE_PASSWORD") //密码
	host := os.Getenv("DATA_BASE_IP") //数据库地址，可以是Ip或者域名
	port, _ := strconv.Atoi(os.Getenv("DATA_BASE_PORT")) //数据库端口
	Dbname := os.Getenv("DATA_BASE_NAME") //数据库名

	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		Dbname,
	)

	//连接MYSQL
	db, err := gorm.Open(
		mysql.Open(DSN),
		&gorm.Config{},
	)

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	DB = db
	fmt.Print("Model set up successfully.")
}