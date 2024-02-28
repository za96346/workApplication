package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"backend/infrastructure/persistence"
	"backend/interfaces/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	repo *persistence.Repositories
	companyApp application.CompanyAppInterface
}

func NewCompany(repo *persistence.Repositories) *CompanyController {
	return &CompanyController{
		repo: repo,
		companyApp: &application.CompanyApp{},
	}
}

func (s *CompanyController) GetCompany(Request *gin.Context) {
	session, err := method.NewSession(Request, nil)
	if err != nil {return}

	responseData, appErr := s.companyApp.GetCompany(
		&entities.Company{
			CompanyId: session.User.CompanyId,
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
	}

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := s.companyApp.UpdateCompany(reqBody, session)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "更新失敗",
			},
		)
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}
