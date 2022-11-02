package route

import (
	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
)

func Company(props *gin.RouterGroup) {
	props.Use(middleWare.TokenPrase)
	{
		// info
		props.GET("/info", middleWare.MyCompanyAndBanch, worker.AssignWorker(17))
		props.POST("/info",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(18),
		)
		props.PUT("/info",
			worker.AssignWorker(23),
		)

		// banch all
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
		props.DELETE("/banch/all",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(25),
		)

		// banch style
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
		props.DELETE("/banch/style",
			middleWare.Permession(100, 1),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(20),
		)
		
		
		// banch rule
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
		props.DELETE("/banch/rule",
			middleWare.Permession(100, 1),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(21),
		)

		// wait company reply
		props.GET("/wait/reply",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(26),
		)
		props.POST("/wait/reply",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(27),
		)
		props.PUT("/wait/reply",
			middleWare.Permession(2),
			middleWare.MyCompanyAndBanch,
			worker.AssignWorker(28),
		)
	}
}