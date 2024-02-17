package route

import (
	"backend/interfaces/middleware"
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)
func User(props *gin.RouterGroup, userController *controller.UserController) {
	props.Use(middleware.Permission)
	{
		props.GET("/my", userController.GetMine)
		props.GET("/", userController.GetUsers)
		props.PUT("/", userController.SaveUser)
		props.POST("/", userController.UpdateUser)
		props.POST("/my", userController.UpdateMine)
		props.DELETE("/", userController.DeleteUser)
		props.POST("/password", userController.UpdatePassword)
		props.GET("/selector", userController.GetUsersSelector)
	}
}
