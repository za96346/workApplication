package Route

import (
	// "strconv"

	"backend/restful/controller/CTL_System"

	"github.com/gin-gonic/gin"
)

func System(props *gin.RouterGroup) {
	props.GET("/auth", CTL_System.GetAuth)
}
