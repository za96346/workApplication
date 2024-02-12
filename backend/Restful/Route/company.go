package Route

import (
	// "strconv"

	"backend/Middleware"
	"backend/Restful/Controller/CTL_Company"

	"github.com/gin-gonic/gin"
)

func Company(props *gin.RouterGroup) {
	props.Use(Middleware.Permission)
	{
		props.GET("/", CTL_Company.Get)
		props.POST("/", CTL_Company.Edit)
	}
}
