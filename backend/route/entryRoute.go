package route

import (
	"github.com/gin-gonic/gin"
	"backend/worker"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)



func EntryRoute(props *gin.RouterGroup) {
	props.POST("/login", worker.AssignWorker(3))
	props.PUT("/register", worker.AssignWorker(4))
}