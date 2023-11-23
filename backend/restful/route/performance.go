package Route

import (
	// "strconv"

	"backend/middleware"
	"backend/restful/controller/CTL_Performance"

	"github.com/gin-gonic/gin"
)

func Performance(props *gin.RouterGroup) {
	props.Use(middleware.Permission)
	{
		props.GET("/", CTL_Performance.Get)
		props.PUT("/", CTL_Performance.Add)
		props.POST("/", CTL_Performance.Edit)
		props.DELETE("/", CTL_Performance.Delete)
		props.GET("/year", CTL_Performance.GetYear)
		props.GET("/searchBar", CTL_Performance.SearchBar)
	}
}
