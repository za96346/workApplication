package route

import (
	// "strconv"

	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
)
func User(props *gin.RouterGroup) {
	props.POST("/forgetPassword", worker.AssignWorker(24))
	props.Use(middleWare.TokenPrase)
	{
		props.POST("/changePassword", worker.AssignWorker(19))
		props.GET("/my", worker.AssignWorker(1))
		props.POST("/my", worker.AssignWorker(22))
		props.GET("/single", middleWare.SameCompany, worker.AssignWorker(0))
		props.POST("/single",
			middleWare.Permession(100),
			middleWare.SameCompany,
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(2),
		)
		props.GET("/all", middleWare.Permession(100, 1), middleWare.MyCompanyAndBanch(true), worker.AssignWorker(7))
	}
}
