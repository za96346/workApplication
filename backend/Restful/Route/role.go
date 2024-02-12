package Route

import (
	// "strconv"

	"backend/middleware"
	"backend/Restful/Controller/CTL_Role"

	"github.com/gin-gonic/gin"
)

func Role(props *gin.RouterGroup) {
	props.Use(middleware.Permission)
	{
		props.GET("/", CTL_Role.Get)
		props.POST("/", CTL_Role.Update)
		props.PUT("/", CTL_Role.Add)
		props.DELETE("/", CTL_Role.Delete)

		props.GET("/single", CTL_Role.GetSingle)
		props.GET("/selector", CTL_Role.GetSelector)
	}
}
