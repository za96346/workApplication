package Route

import (
	// "strconv"

	"backend/Middleware"
	"backend/Restful/Controller/CTL_User"

	"github.com/gin-gonic/gin"
)
func User(props *gin.RouterGroup) {
	props.Use(Middleware.Permission)
	{
		props.GET("/my", CTL_User.GetMine)
		props.GET("/", CTL_User.Get)
		props.PUT("/", CTL_User.Add)
		props.POST("/", CTL_User.Edit)
		props.POST("/my", CTL_User.EditMine)
		props.DELETE("/", CTL_User.Delete)
		props.POST("/password", CTL_User.UpdatePassword)
		props.GET("/selector", CTL_User.GetSelector)
	}
}
