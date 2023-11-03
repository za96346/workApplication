package Route

import (
	// "strconv"

	"backend/middleware"
	"backend/restful/controller/CTL_Company"

	"github.com/gin-gonic/gin"
)

func Company(props *gin.RouterGroup) {
	props.Use(middleware.Permission)
	{
		props.GET("/my", CTL_Company.Get)
	}
}
