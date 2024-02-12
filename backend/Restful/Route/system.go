package Route

import (
	// "strconv"

	"backend/Restful/controller/CTL_System"

	"github.com/gin-gonic/gin"
)

func System(props *gin.RouterGroup) {
	props.GET("/auth", CTL_System.GetAuth)
	props.GET("/func", CTL_System.GetFunctionItem)
	props.GET("/roleBanchList", CTL_System.GetRoleBanchList)
}
