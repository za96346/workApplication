package route

import (
	// "strconv"

	"github.com/gin-gonic/gin"
	"backend/worker"
)
func User(props *gin.RouterGroup) {

	props.GET("/fetch/single/:userId", worker.AssignWorker(0))
	props.POST("/create", worker.AssignWorker(1))
	props.PUT("/update", worker.AssignWorker(2))
}
