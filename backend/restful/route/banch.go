package Route

import (
	// "strconv"

	"backend/middleware"
	"backend/restful/controller/CTL_Banch"

	"github.com/gin-gonic/gin"
)

func Banch(props *gin.RouterGroup) {
	props.Use(middleware.Permission)
	{
		props.GET("/", CTL_Banch.Get)
		props.POST("/", CTL_Banch.Edit)
		props.PUT("/", CTL_Banch.Add)
		props.DELETE("/", CTL_Banch.Delete)
	}
}
