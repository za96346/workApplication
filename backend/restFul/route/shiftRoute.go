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
			worker.AssignWorker(32),
		)
		props.PUT("/workTime",
			middleWare.Permession(100),
			worker.AssignWorker(33),
		)
		props.POST("/workTime",
			middleWare.Permession(100),
			worker.AssignWorker(34),
		)
		props.DELETE("/workTime",
			middleWare.Permession(100),
			worker.AssignWorker(35),
		)
	
		// paidVocation
		props.GET("/paidVocation",
			middleWare.Permession(100, 1, 2),
			worker.AssignWorker(36),
		)
		props.PUT("/paidVocation",
			middleWare.Permession(100),
			worker.AssignWorker(37),
		)
		props.POST("/paidVocation",
			middleWare.Permession(100),
			worker.AssignWorker(38),
		)
		props.DELETE("/paidVocation",
			middleWare.Permession(100),
			worker.AssignWorker(39),
		)

		// 班表 查詢
		props.GET("/month",
			middleWare.Permession(100, 1, 2),
			worker.AssignWorker(49),
		)
	}
}
