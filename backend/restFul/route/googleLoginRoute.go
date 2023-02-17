package route

import (
	"backend/worker"
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)


func GoogleLoginRoute(props *gin.RouterGroup) {
	props.GET("/login", worker.AssignWorker(44))
	props.POST("/login", worker.AssignWorker(45))
}