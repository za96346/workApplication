package route

import (
	// "strconv"

	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
)
func User(props *gin.RouterGroup) {

	props.Use(middleWare.TokenPrase)
	{
		props.GET("/single/:userId", worker.AssignWorker(0))
		props.GET("/my", worker.AssignWorker(1))
		props.PUT("/update", worker.AssignWorker(2))
	}
}
