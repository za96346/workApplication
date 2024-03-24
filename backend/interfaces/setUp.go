package interfaces

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"

	application "backend/application/services"
	"backend/infrastructure/persistence"
	"backend/interfaces/controller"
	"backend/interfaces/middleware"
	"backend/interfaces/route"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
)

// 設定 http server
func SetUp(repo *persistence.Repositories) *gin.Engine {
	port := os.Getenv("PORT")
	apiServer := gin.Default()

	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := memstore.NewStore(
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
	companyController := controller.NewCompany(
		&application.CompanyApp{
			CompanyRepo: repo.Company,
			CompanyBanchRepo: repo.CompanyBanch,
			RoleRepo: repo.Role,
		},
	)
	companyBanchController := controller.NewCompanyBanch(
		&application.CompanyBanchApp{
			CompanyBanchRepo: repo.CompanyBanch,
			RoleRepo: repo.Role,
		},
	)
	entryController := controller.NewEntry(
		&application.EntryApp{
			UserRepo: repo.User,
		},
	)
	performanceController := controller.NewPerformance(
		&application.PerformanceApp{
			PerformanceRepo: repo.Performance,
			UserRepo: repo.User,
			CompanyBanchRepo: repo.CompanyBanch,
			RoleRepo: repo.Role,
		},
	)
	roleController := controller.NewRole(
		&application.RoleApp{
			RoleRepo: repo.Role,
			RoleStructRepo: repo.RoleStruct,
			CompanyBanchRepo: repo.CompanyBanch,
		},
	)
	systemController := controller.NewSystem(
		&application.SystemApp{
			RoleRepo: repo.Role,
			RoleStructRepo: repo.RoleStruct,
			FunctionItemRepo: repo.FunctionItem,
			FunctionRoleBanchRelationRepo: repo.FunctionRoleBanchRelation,
			OperationItemRepo: repo.OperationItem,
			CompanyBanchRepo: repo.CompanyBanch,
			UserRepo: repo.User,
		},
	)
	userController := controller.NewUser(
		&application.UserApp{
			UserRepo: repo.User,
			RoleRepo: repo.Role,
			CompanyBanchRepo: repo.CompanyBanch,
		},
	)

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
