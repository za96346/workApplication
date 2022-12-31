package route

import (
	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
)

func Performance(props *gin.RouterGroup) {
	props.Use(middleWare.TokenPrase)
	{
		props.GET("/performance", middleWare.Permession(100, 1, 2), worker.AssignWorker(40))
		props.POST("/performance", middleWare.Permession(100, 1, 2), worker.AssignWorker(41))
		props.PUT("/performance", middleWare.Permession(100, 1), worker.AssignWorker(42))
		props.DELETE("/performance", middleWare.Permession(100, 1), worker.AssignWorker(43))

	}
}