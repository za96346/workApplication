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
	"runtime"
	"time"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// . "./middleWare/permessionMiddleWare"
	"backend/handler"
	_ "backend/handler"
	"backend/middleWare"
	. "backend/route"
	"backend/worker"
)
func init() {
	handler.Init()
}
func main() {
	runtime.SetMutexProfileFraction(-1)
	worker.WorkerSingleton().CreateWorker(runtime.NumCPU() * 2)
	fmt.Println("開啟的worker數量", runtime.NumCPU() * 2)

	go func ()  {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	port := os.Getenv("PORT")
	apiServer := gin.Default()
	apiServer.Use(middleWare.RateLimit(time.Second, 100, 100))

	userApi := apiServer.Group("/workApp/user")
	UserRouter(userApi)
	// apiServer.Use(permessionMiddleWare("a1234"))
	apiServer.Run(":" + port)
}
