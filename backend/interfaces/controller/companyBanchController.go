package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"backend/interfaces/enum"
	"backend/interfaces/method"
	"net/http"

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

// 拿取
func (s *CompanyBanchController) GetCompanyBanches(Request *gin.Context) {
	reqParams := new(struct{
		BanchName   *string       `gorm:"column:banchName" json:"BanchName"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.BanchManage),
		PermissionItemCode: "inquire",

		ReqBodyValidation: false,
		ReqParamsStruct: reqParams,
		ReqParamsValidation: true,
	}
	if session.SessionHandler() != nil {return}

	responseData, _ := s.companyBanchApp.GetCompanyBanches(
		session.CompanyId,
		&session.CurrentPermissionScopeBanch,
		reqParams.BanchName,
	)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    responseData,
		},
	)
}

func (s *CompanyBanchController) GetCompanyBanchesSelector(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		PermissionValidation: false,
	}
	if session.SessionHandler() != nil {return}

	responseData, _ := s.companyBanchApp.GetCompanyBanchesSelector(session.CompanyId)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": responseData,
		},
	)
}

func (s *CompanyBanchController) UpdateCompanyBanch(Request *gin.Context) {
	reqBody := new(entities.CompanyBanch)
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.BanchManage),
		PermissionItemCode: "edit",
	}
	if session.SessionHandler() != nil {return}
	if session.CheckScopeBanchValidation((*reqBody).BanchId) != nil {return}

	s.companyBanchApp.UpdateCompanyBanch(reqBody)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

func (s *CompanyBanchController) SaveCompanyBanch(Request *gin.Context) {
	reqBody := new(entities.CompanyBanch)
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.BanchManage),
		PermissionItemCode: "add",
	}
	if session.SessionHandler() != nil {return}

	s.companyBanchApp.SaveCompanyBanch(reqBody)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "新增成功",
		},
	)
}

func (s *CompanyBanchController) DeleteCompanyBanch(Request *gin.Context) {
	reqBody := new(struct {
		BanchId int `json:"BanchId" binding:"required"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.BanchManage),
		PermissionItemCode: "delete",
	}
	if session.SessionHandler() != nil {return}
	if session.CheckScopeBanchValidation((*reqBody).BanchId) != nil {return}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}
