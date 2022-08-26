package route

import (
	. "backend/service"
	// "fmt"

	"github.com/gin-gonic/gin"
)
func UserRouter(props *gin.RouterGroup) {
	props.GET("/fetch/single/:userId", FindSingleUser)
	props.POST("/create", CreateUser)
	props.PUT("/update", UpdateUser)
}