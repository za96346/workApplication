//go get github.com/gin-gonic/gin
//go get github.com/joho/godotenv

// http status code reference => https://go.dev/src/net/http/status.go
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// . "./middleWare/permessionMiddleWare"
	"backend/database"
	"backend/redis"
	. "backend/route"
	"backend/worker"
)
var mysqlDB = database.DBSingleton()
func main() {
	runtime.SetMutexProfileFraction(-1)
	worker.WorkerSingleton().CreateWorker(200)

	go func ()  {
		Gorace()		
	}()

	go func() {
		redis.RedisDBConn()
		(*mysqlDB).Conn()
		
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

	http.ListenAndServe("0.0.0.0:6060", nil)
}
