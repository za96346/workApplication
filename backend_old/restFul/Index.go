package restFul

import (
	_ "net/http/pprof"
	"os"
	"time"

	"backend/logger"

	"backend/middleWare"
	"backend/restFul/route"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	port := os.Getenv("PORT")
	apiServer := gin.Default()
	apiServer.Use(middleWare.RateLimit(time.Second, 100, 100), middleWare.CORS, logger.LoggerToFile())

	// route group
	userApi := apiServer.Group("/workApp/user")
	entryApi := apiServer.Group("/workApp/entry")
	companyApi := apiServer.Group("/workApp/company")
	shiftApi := apiServer.Group("/workApp/shift")
	performanceApi := apiServer.Group("/workApp/pr")
	google := apiServer.Group("/workApp/google")
	

	route.User(userApi)
	route.EntryRoute(entryApi)
	route.Company(companyApi)
	route.Shift(shiftApi)
	route.Performance(performanceApi)
	route.GoogleLoginRoute(google)

	// start
	apiServer.Run(":" + port)
	return apiServer
}
