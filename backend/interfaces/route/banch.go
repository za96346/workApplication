package route

import (
	"backend/interfaces/middleware"
	"backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Banch(props *gin.RouterGroup, companyBanchController *controller.CompanyBanchController) {
	props.Use(middleware.Permission)
	{
		props.GET("/", companyBanchController.GetCompanyBanches)
		props.POST("/", companyBanchController.UpdateCompanyBanch)
		props.PUT("/", companyBanchController.SaveCompanyBanch)
		props.DELETE("/", companyBanchController.DeleteCompanyBanch)

		props.GET("/selector", companyBanchController.GetCompanyBanchesSelector)
	}
}
