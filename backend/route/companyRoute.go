package route

import (
	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
)

func Company(props *gin.RouterGroup) {
	props.Use(middleWare.TokenPrase)
	{
		props.GET("/banch/all", worker.AssignWorker(8))
		props.GET("/banch/style", worker.AssignWorker(9))
	}
}