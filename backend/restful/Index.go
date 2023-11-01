package restFul

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"

	"backend/middleware"
	"backend/restFul/route"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// 設定 http server
func SetApiServer() *gin.Engine {
	port := os.Getenv("PORT")
	apiServer := gin.Default()

	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore(
		[]byte("secret11111"),
	)

	apiServer.Use(
		sessions.Sessions("workapp_session", store),
		middleware.RateLimit(time.Second, 100, 100),
		middleware.CORS,
	)

	// route group
	userApi := apiServer.Group("/workApp/user")
	

	route.User(userApi)

	// start
	apiServer.Run(":" + port)

	fmt.Print("api server started successfully.")

	return apiServer
}
