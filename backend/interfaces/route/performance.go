package route

import (
	"backend/interfaces/middleware"
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Performance(props *gin.RouterGroup, performanceController *controller.PerformanceController) {
	props.Use(middleware.Permission)
	{
		props.GET("/", performanceController.GetPerformances)
		props.PUT("/", performanceController.SavePerformance)

		props.POST("/", performanceController.UpdatePerformance)
		props.DELETE("/", performanceController.DeletePerformance)
		props.PUT("/copy", performanceController.SavePerformance)
		props.POST("/banch", performanceController.ChangeBanch)
		props.GET("/year", performanceController.GetYearPerformances)
	}
}
