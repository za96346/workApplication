package route


import (
	"backend/interfaces/middleware"
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Entry(props *gin.RouterGroup) {
	props.POST("/login", CTL_Entry.Login)
}
