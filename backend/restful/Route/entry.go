package Route


import (
	// "strconv"

	"backend/restful/controller/CTL_Entry"

	"github.com/gin-gonic/gin"
)

func Entry(props *gin.RouterGroup) {
	props.POST("/login", CTL_Entry.Login)
}
