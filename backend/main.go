package main

import (
	"backend/infrastructure/persistence"
	"backend/interfaces"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

func init() {
	if godotenv.Load(filepath.Join("./", ".env")) != nil {
		panic("error loading .env file")
	}
}

func main() {
	userName := os.Getenv("DATA_BASE_USER")  //账号
	password := os.Getenv("DATA_BASE_PASSWORD") //密码
	host := os.Getenv("DATA_BASE_IP") //数据库地址，可以是Ip或者域名
	port := os.Getenv("DATA_BASE_PORT") //数据库端口
	dbName := os.Getenv("DATA_BASE_NAME") //数据库名

	repo, err := persistence.NewRepositories(
		"mysql",
		userName,
		password,
		port,
		host,
		dbName,
	)
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	interfaces.SetUp(repo)
}