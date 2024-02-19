package controller

import (
	"backend/interfaces/method"
	"backend/application/services"
	"backend/domain/entities"
	"backend/interfaces/enum"
	"net/http"

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

func (s *CompanyController) GetCompany(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.CompanyData),
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler() != nil {return}

	responseData, _ := s.companyApp.GetCompany(
		&entities.Company{
			CompanyId: session.CompanyId,
		},
	)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    responseData,
		},
	)
}

func (s *CompanyController) UpdateCompany(Request *gin.Context) {
	reqBody := new(entities.Company)
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.CompanyData),
		PermissionItemCode: "edit",
	}
	if session.SessionHandler() != nil {return}

	s.companyApp.UpdateCompany(reqBody)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}
