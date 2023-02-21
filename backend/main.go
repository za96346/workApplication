//go get github.com/gin-gonic/gin
//go get github.com/joho/godotenv

// http status code reference => https://go.dev/src/net/http/status.go
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"backend/handler"
	"backend/restFul"
	"backend/taskTimer"

	"github.com/joho/godotenv"

	"backend/socket"
	"backend/worker"
	"path/filepath"
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
	go socket.Conn()
	go tasktimer.AddDailyTask(
		func() {
			handler.SendDailyInfo("za96346@gmail.com")
			handler.SendDailyInfo("a00001@dajiama.org.tw")
		},
	)
	restFul.SetRouter()
}


