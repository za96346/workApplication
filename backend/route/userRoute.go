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
		props.GET("/my", worker.AssignWorker(1))
		props.GET("/single", middleWare.SameCompany, worker.AssignWorker(0))
		props.POST("/update", middleWare.SameCompany, worker.AssignWorker(2))
		props.GET("/all", middleWare.Permession(100), worker.AssignWorker(7))
	}
}
