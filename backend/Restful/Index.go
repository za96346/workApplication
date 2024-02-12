package restFul

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"

	"backend/middleware"
	"backend/RestFul/Route"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// 設定 http server
func SetUp() *gin.Engine {
	port := os.Getenv("PORT")
	apiServer := gin.Default()

	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore(
		[]byte("secret11111"),
	)

	// store.Options(sessions.Options{
	// 	HttpOnly: true,
	// })

	apiServer.Use(
		middleware.CORS(),
		sessions.Sessions("workapp_session", store),
		middleware.RateLimit(time.Second, 100, 100),
	)

	// route group
	userApi := apiServer.Group("/workApp/user")
	roleApi := apiServer.Group("/workApp/role")
	companyApi := apiServer.Group("/workApp/company")
	entryApi := apiServer.Group("/workApp/entry")
	systemApi := apiServer.Group("/workApp/system")
	banchApi := apiServer.Group("/workApp/banch")
	performanceApi := apiServer.Group("/workApp/performance")

	Route.User(userApi)
	Route.Role(roleApi)
	Route.Company(companyApi)
	Route.Entry(entryApi)
	Route.System(systemApi)
	Route.Banch(banchApi)
	Route.Performance(performanceApi)

	// start
	apiServer.Run(":" + port)

	fmt.Print("api server started successfully.")

	return apiServer
}
