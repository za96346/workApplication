package Route

import (
	// "strconv"

	"backend/middleware"
	"backend/restful/controller/CTL_Role"

	"github.com/gin-gonic/gin"
)

func Role(props *gin.RouterGroup) {
	props.Use(middleware.Permission)
	{
		props.GET("/", CTL_Role.Get)
		props.POST("/", CTL_Role.Update)
		props.GET("/single", CTL_Role.GetSingle)
		props.DELETE("", CTL_Role.Delete)
	}
}
