package route

import (
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func System(props *gin.RouterGroup, systemController *controller.SystemController) {
	props.GET("/auth", systemController.GetAuth)
	props.GET("/func", systemController.GetFunc)
	props.GET("/roleBanchList", systemController.GetRoleBanchList)
}
