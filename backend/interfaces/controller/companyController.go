package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	companyApp application.CompanyAppInterface
}

func NewCompany() *CompanyController {
	return &CompanyController{
		companyApp: &application.CompanyApp{},
	}
}

func (s *CompanyController) UpdateCompany(c *gin.Context) {
	var companyEntity entities.Company
	if err := c.ShouldBindJSON(&companyEntity); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

}

func (s *CompanyController) GetCompany(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Print(userId)
}
