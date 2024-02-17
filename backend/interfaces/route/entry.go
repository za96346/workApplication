package route


import (
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Entry(props *gin.RouterGroup, entryController *controller.EntryController) {
	props.POST("/login", entryController.Login)
}
