//go get github.com/gin-gonic/gin
//go get github.com/joho/godotenv

// http status code reference => https://go.dev/src/net/http/status.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	_ "net/http/pprof"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// . "./middleWare/permessionMiddleWare"
	"backend/database"
	"backend/redis"
	. "backend/route"

)
var mysqlDB = database.DBSingleton()
func main() {
	runtime.SetMutexProfileFraction(-1)
	go func ()  {
		Gorace()		
	}()

	go func() {
		redis.RedisDBConn()
		mysqlDB.Conn()
		
	}()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	port := os.Getenv("PORT")
	apiServer := gin.Default()

	userApi := apiServer.Group("/workApp/user")
	UserRouter(userApi)
	// apiServer.Use(permessionMiddleWare("a1234"))
	apiServer.Run(":" + port)
}
func Gorace() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}

	http.ListenAndServe("0.0.0.0:6060", nil)
}
