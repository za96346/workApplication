package controller

import (
	application "backend/application/services"
	"backend/domain/entities"
	"backend/infrastructure/persistence"
	"backend/interfaces/enum"
	"backend/interfaces/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyBanchController struct {
	repo *persistence.Repositories
	companyBanchApp application.CompanyBanchAppInterface
}

func NewCompanyBanch(repo *persistence.Repositories) *CompanyBanchController {
	return &CompanyBanchController{
		repo: repo,
		companyBanchApp: &application.CompanyBanchApp{},
	}
}

// 拿取
func (s *CompanyBanchController) GetCompanyBanches(Request *gin.Context) {
	reqParams := new(struct{
		BanchName   *string       `gorm:"column:banchName" json:"BanchName"`
	})

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqParamsStruct: reqParams,
			ReqParamsValidation: true,
		},
	)
	if err != nil {return}

	responseData, _ := s.companyBanchApp.GetCompanyBanches(
		&entities.CompanyBanch{
			BanchName: *reqParams.BanchName,
		},
		session,
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
	session, err := method.NewSession(
		Request,
		nil,
	)
	if err != nil {return}

	responseData, _ := s.companyBanchApp.GetCompanyBanchesSelector(session)

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
	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

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
