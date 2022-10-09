package route

import (
	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)



func EntryRoute(props *gin.RouterGroup) {
	props.POST("/login", worker.AssignWorker(3))
	props.PUT("/register", worker.AssignWorker(4))
	props.POST("/email/captcha", worker.AssignWorker(6))
	props.Use(middleWare.TokenPrase)
	{
		props.GET("/checkAccess", worker.AssignWorker(5))
	}
}