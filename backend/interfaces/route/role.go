package route

import (
	"backend/interfaces/middleware"
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Role(props *gin.RouterGroup) {
	props.Use(Middleware.Permission)
	{
		props.GET("/", CTL_Role.Get)
		props.POST("/", CTL_Role.Update)
		props.PUT("/", CTL_Role.Add)
		props.DELETE("/", CTL_Role.Delete)

		props.GET("/single", CTL_Role.GetSingle)
		props.GET("/selector", CTL_Role.GetSelector)
	}
}
