package route

import (
	"github.com/gin-gonic/gin"
	"backend/worker"
)

func EntryRoute(props *gin.RouterGroup) {
	props.POST("/login", worker.AssignWorker(3))
	props.POST("/register", worker.AssignWorker(4))
}