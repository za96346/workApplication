package interfaces

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"

	"backend/infrastructure/persistence"
	"backend/interfaces/controller"
	"backend/interfaces/middleware"
	"backend/interfaces/route"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// 設定 http server
func SetUp(repo *persistence.Repositories) *gin.Engine {
	port := os.Getenv("PORT")
	apiServer := gin.Default()

	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore(
		[]byte("secret11111"),
	)

	// 添加全域middleware
	apiServer.Use(
		middleware.CORS(),
		sessions.Sessions("workapp_session", store),
		middleware.RateLimit(time.Second, 100, 100),
	)

	// 新增 route group
	userApi := apiServer.Group("/workApp/user")
	roleApi := apiServer.Group("/workApp/role")
	companyApi := apiServer.Group("/workApp/company")
	entryApi := apiServer.Group("/workApp/entry")
	systemApi := apiServer.Group("/workApp/system")
	banchApi := apiServer.Group("/workApp/banch")
	performanceApi := apiServer.Group("/workApp/performance")

	// 實例 app
	companyController := controller.NewCompany(repo)
	companyBanchController := controller.NewCompanyBanch(repo)
	entryController := controller.NewEntry(repo)
	performanceController := controller.NewPerformance(repo)
	roleController := controller.NewRole(repo)
	systemController := controller.NewSystem(repo)
	userController := controller.NewUser(repo)

	// 嵌入 route group
	route.User(userApi, userController)
	route.Role(roleApi, roleController)
	route.Company(companyApi, companyController)
	route.Entry(entryApi, entryController)
	route.System(systemApi, systemController)
	route.Banch(banchApi, companyBanchController)
	route.Performance(performanceApi, performanceController)

	// start
	apiServer.Run(":" + port)

	fmt.Print("api server started successfully.")

	return apiServer
}
