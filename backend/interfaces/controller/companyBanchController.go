package controller

import (
	application "backend/application/services"
	"backend/domain/entities"
	"backend/interfaces/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyBanchController struct {
	companyBanchApp application.CompanyBanchAppInterface
}

func NewCompanyBanch(app application.CompanyBanchAppInterface) *CompanyBanchController {
	return &CompanyBanchController{
		companyBanchApp: app,
	}
}

// 拿取
func (s *CompanyBanchController) GetCompanyBanches(Request *gin.Context) {
	reqParams := new(struct{
		BanchName   string       `gorm:"column:banchName" json:"BanchName"`
	})

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqParamsStruct: reqParams,
			ReqParamsValidation: true,
		},
	)
	if err != nil {return}

	responseData, appErr := s.companyBanchApp.GetCompanyBanches(
		&entities.CompanyBanch{
			BanchName: (*reqParams).BanchName,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "失敗",
				"data":    []int{},
			},
		)
		return
	}

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

	responseData, appErr := s.companyBanchApp.GetCompanyBanchesSelector(session)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "失敗",
				"data":    []int{},
			},
		)
		return
	}

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

	_, appErr := s.companyBanchApp.UpdateCompanyBanch(reqBody, session)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "更新失敗",
			},
		)
		return
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

func (s *CompanyBanchController) SaveCompanyBanch(Request *gin.Context) {
	reqBody := new(entities.CompanyBanch)
	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := s.companyBanchApp.SaveCompanyBanch(reqBody, session)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "新增失敗",
			},
		)
		return
	}

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := s.companyBanchApp.DeleteCompanyBanch(
		&entities.CompanyBanch{
			BanchId: reqBody.BanchId,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "刪除失敗",
			},
		)
		return
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}
