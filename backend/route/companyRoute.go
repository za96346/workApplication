package route

import (
	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
)

func Company(props *gin.RouterGroup) {
	props.Use(middleWare.TokenPrase)
	{
		props.GET("/info", middleWare.MyCompanyAndBanch, worker.AssignWorker(17))
		props.POST("/info",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(18),
		)

		props.GET("/banch/all", middleWare.MyCompanyAndBanch, worker.AssignWorker(8))
		props.PUT("/banch/all",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(15),
		)
		props.POST("/banch/all",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(16),
		)

		props.GET("/banch/style", middleWare.MyCompanyAndBanch, worker.AssignWorker(9))
		props.POST("/banch/style",
			middleWare.Permession(100, 1),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(11),
		)
		props.PUT("/banch/style",
			middleWare.Permession(100, 1),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(13),
		)

		props.GET("/banch/rule", middleWare.MyCompanyAndBanch, worker.AssignWorker(10))
		props.POST("/banch/rule",
			middleWare.Permession(100, 1),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(12),
		)
		props.PUT("/banch/rule",
			middleWare.Permession(100, 1),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(14),
		)
	}
}