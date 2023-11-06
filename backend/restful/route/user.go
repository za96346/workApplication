package Route

import (
	// "strconv"

	"backend/middleware"
	"backend/restful/controller/CTL_User"

	"github.com/gin-gonic/gin"
)
func User(props *gin.RouterGroup) {
	props.Use(middleware.Permission)
	{
		props.GET("/my", CTL_User.GetMine)
		props.GET("/", CTL_User.Get)
		props.PUT("/", CTL_User.Add)
		props.POST("/", CTL_User.Edit)
		props.DELETE("/", CTL_User.Delete)
	}
}
