package route

import (
	"backend/interfaces/middleware"
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Performance(props *gin.RouterGroup) {
	props.Use(Middleware.Permission)
	{
		props.GET("/", CTL_Performance.Get)
		props.PUT("/", CTL_Performance.Add)

		props.POST("/", CTL_Performance.Edit)
		props.DELETE("/", CTL_Performance.Delete)
		props.PUT("/copy", CTL_Performance.Add)
		props.POST("/banch", CTL_Performance.ChangeBanch)
		props.GET("/year", CTL_Performance.GetYear)
	}
}
