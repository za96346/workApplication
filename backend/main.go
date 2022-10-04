//go get github.com/gin-gonic/gin
//go get github.com/joho/godotenv

// http status code reference => https://go.dev/src/net/http/status.go
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"io"
	"runtime"
	"time"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// . "./middleWare/permessionMiddleWare"
	"backend/handler"
	"backend/middleWare"
	"backend/route"
	"backend/worker"
)
func init() {
	handler.Init()
	runtime.SetMutexProfileFraction(-1)
	worker.WorkerSingleton().CreateWorker(runtime.NumCPU() * 2)
	fmt.Println("開啟的worker數量", runtime.NumCPU() * 2)
	if godotenv.Load() != nil {
		log.Fatal("error loading .env file")
	}
	go func ()  {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
}

func main() {

	port := os.Getenv("PORT")
	loggers()
	apiServer := gin.Default()
	apiServer.Use(middleWare.RateLimit(time.Second, 100, 100))

	userApi := apiServer.Group("/workApp/user")
	entryApi := apiServer.Group("/workApp/entry")
	route.User(userApi)
	route.EntryRoute(entryApi)
	apiServer.Run(":" + port)
}

func loggers() {
	file, _ := os.Create("gin.log")                     // create log file
    gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}