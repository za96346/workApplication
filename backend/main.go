//go get github.com/gin-gonic/gin
//go get github.com/joho/godotenv


//http status code reference => https://go.dev/src/net/http/status.go
package main

import (
	// "fmt"
	"log"
	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// . "./middleWare/permessionMiddleWare"
	. "backend/route"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	port := os.Getenv("PORT")
	apiServer := gin.Default()

	userApi := apiServer.Group("/workApp/user")
	UserRouter(userApi)
	// apiServer.Use(permessionMiddleWare("a1234"))
	apiServer.Run(port)
}

