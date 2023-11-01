package route

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
	}
}
