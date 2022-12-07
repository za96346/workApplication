package route

import (
	// "strconv"

	"backend/middleWare"
	"backend/worker"

	"github.com/gin-gonic/gin"
)
func Shift (props *gin.RouterGroup) {
	props.Use(middleWare.TokenPrase)
	{
		// workTime
		props.GET("/workTime",
			middleWare.Permession(100, 1, 2),
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(32),
		)
		props.PUT("/workTime",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(33),
		)
		props.POST("/workTime",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(34),
		)
		props.DELETE("/workTime",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(35),
		)
	
		// paidVocation
		props.GET("/paidVocation",
			middleWare.Permession(100, 1, 2),
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(36),
		)
		props.PUT("/paidVocation",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(37),
		)
		props.POST("/paidVocation",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(38),
		)
		props.DELETE("/paidVocation",
			middleWare.Permession(100),
			middleWare.MyCompanyAndBanch(true),
			worker.AssignWorker(39),
		)
	}
}
