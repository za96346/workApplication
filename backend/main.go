//go get github.com/gin-gonic/gin
//go get github.com/joho/godotenv

// http status code reference => https://go.dev/src/net/http/status.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	// "time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// . "./middleWare/permessionMiddleWare"
	"backend/middleWare"
	"backend/route"
	"backend/worker"
	"path/filepath"
	"backend/socket"
)
func init() {
	runtime.SetMutexProfileFraction(-1)
	worker.WorkerSingleton().CreateWorker(runtime.NumCPU() * 2)
	fmt.Println("開啟的worker數量", runtime.NumCPU() * 2)
	if godotenv.Load(filepath.Join("./", ".env")) != nil {
		log.Fatal("error loading .env file")
	}
	go func ()  {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
}

func main() {
	SetRouter()
}

func SetRouter() *gin.Engine {
	port := os.Getenv("PORT")
	loggers()
	apiServer := gin.Default()
	apiServer.Use(middleWare.RateLimit(time.Second, 100, 100), middleWare.CORS)

	// route group
	userApi := apiServer.Group("/workApp/user")
	entryApi := apiServer.Group("/workApp/entry")
	companyApi := apiServer.Group("/workApp/company")
	shiftApi := apiServer.Group("/workApp/shift")
	route.User(userApi)
	route.EntryRoute(entryApi)
	route.Company(companyApi)
	route.Shift(shiftApi)

	go socket.Conn()

	// start
	apiServer.Run(":" + port)
	return apiServer
}

func loggers() {
	file, _ := os.Create("gin.log")                     // create log file
    gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}