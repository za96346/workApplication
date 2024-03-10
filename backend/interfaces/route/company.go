package route

import (
	"backend/interfaces/middleware"
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Company(props *gin.RouterGroup, companyController *controller.CompanyController) {
	props.Use(middleware.Permission)
	{
		props.GET("/", companyController.GetCompany)
		props.POST("/", companyController.UpdateCompany)
	}
}
