package route

import (
	"backend/interfaces/middleware"
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Role(props *gin.RouterGroup, roleController *controller.RoleController) {
	props.Use(middleware.Permission)
	{
		props.GET("/", roleController.GetRoles)
		props.POST("/", roleController.UpdateRole)
		props.PUT("/", roleController.SaveRole)
		props.DELETE("/", roleController.DeleteRole)

		props.GET("/single", roleController.GetRole)
		props.GET("/selector", roleController.GetRolesSelector)
	}
}
