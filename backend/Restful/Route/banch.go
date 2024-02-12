package Route

import (
	// "strconv"

	"backend/Middleware"
	"backend/Restful/Controller/CTL_Banch"

	"github.com/gin-gonic/gin"
)

func Banch(props *gin.RouterGroup) {
	props.Use(Middleware.Permission)
	{
		props.GET("/", CTL_Banch.Get)
		props.POST("/", CTL_Banch.Edit)
		props.PUT("/", CTL_Banch.Add)
		props.DELETE("/", CTL_Banch.Delete)

		props.GET("/selector", CTL_Banch.GetSelector)
	}
}
