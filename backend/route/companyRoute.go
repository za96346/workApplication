package route

import (
	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
)

func Company(props *gin.RouterGroup) {
	props.Use(middleWare.TokenPrase)
	{
		props.GET("/banch/all", middleWare.MyCompanyAndBanch, worker.AssignWorker(8))

		props.GET("/banch/style", middleWare.MyCompanyAndBanch, worker.AssignWorker(9))
		props.POST("/banch/style", middleWare.Permession(100, 1), middleWare.MyCompanyAndBanch, worker.AssignWorker(11))
		props.PUT("/banch/style", middleWare.Permession(100, 1), middleWare.MyCompanyAndBanch)

		props.GET("/banch/rule", middleWare.MyCompanyAndBanch, worker.AssignWorker(10))
		props.POST("/banch/rule", middleWare.Permession(100, 1), middleWare.MyCompanyAndBanch)
		props.PUT("/banch/rule", middleWare.Permession(100, 1), middleWare.MyCompanyAndBanch)
	}
}