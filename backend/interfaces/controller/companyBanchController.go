package controller

import (
	"backend/application/services"

	"github.com/gin-gonic/gin"
)

type CompanyBanchController struct {
	companyBanchApp application.CompanyBanchAppInterface
}

func NewCompanyBanch() *CompanyBanchController {
	return &CompanyBanchController{
		companyBanchApp: &application.CompanyBanchApp{},
	}
}

func (s *CompanyBanchController) GetCompanyBanches(c *gin.Context) {
}

func (s *CompanyBanchController) GetCompanyBanchesSelector(c *gin.Context) {
}

func (s *CompanyBanchController) UpdateCompanyBanch(c *gin.Context) {
}

func (s *CompanyBanchController) SaveCompanyBanch(c *gin.Context) {
}

func (s *CompanyBanchController) DeleteCompanyBanch(c *gin.Context) {
}
