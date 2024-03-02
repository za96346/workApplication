package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"backend/infrastructure/persistence"
	"backend/interfaces/method"
	"backend/application/dtos"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PerformanceController struct {
	repo *persistence.Repositories
	performanceApp application.PerformanceAppInterface
}

func NewPerformance(repo *persistence.Repositories) *PerformanceController {
	return &PerformanceController{
		repo: repo,
		performanceApp: &application.PerformanceApp{},
	}
}

func (e *PerformanceController) GetPerformances(Request *gin.Context) {
	reqParams := new(dtos.PerformanceQueryParams)

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqParamsValidation: true,
			ReqParamsStruct: reqParams,
		},
	)
	if err != nil {return}

	data, appErr := e.performanceApp.GetPerformances(
		&entities.Performance{
			BanchId: *reqParams.BanchId,
		},
		reqParams,
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
			"data":    *data,
		},
	)
}

func (e *PerformanceController) GetYearPerformances(Request *gin.Context) {
	reqParams := new(dtos.PerformanceQueryParams)

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqParamsValidation: true,
			ReqParamsStruct: reqParams,
		},
	)
	if err != nil {return}

	data, appErr := e.performanceApp.GetYearPerformances(
		&entities.Performance{},
		reqParams,
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
			"data":    *data,
		},
	)
}

func (e *PerformanceController) SavePerformance(Request *gin.Context) {
	reqBody := new(entities.Performance)

	// 不想多寫一個 copy
	// PermissionItemCode := "add"
	// if strings.Contains(Request.Request.URL.Path, "copy") {
	// 	PermissionItemCode = "copy"
	// }

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := e.performanceApp.SavePerformance(reqBody, session)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "新增失敗",
			},
		)
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "新增成功",
		},
	)
}

func (e *PerformanceController) UpdatePerformance(Request *gin.Context) {
	reqBody := new(entities.Performance)

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := e.performanceApp.UpdatePerformance(reqBody, session)

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

func (e *PerformanceController) DeletePerformance(Request *gin.Context) {
	reqBody := new(struct {
		PerformanceId int `gorm:"column:performanceId;primaryKey" json:"PerformanceId" binding:"required"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.Performance),
		PermissionItemCode: "delete",
	}
	if session.SessionHandler() != nil {return}

	e.performanceApp.DeletePerformance(&entities.Performance{
		CompanyId: session.CompanyId,
		PerformanceId: reqBody.PerformanceId,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}

func (e *PerformanceController) ChangeBanch(Request *gin.Context) {
	reqBody := new(struct {
		PerformanceId   int         `json:"PerformanceId" binding:"required"`
		BanchId         int         `json:"BanchId" binding:"required"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode:  string(enum.Performance),
		PermissionItemCode: "edit",
	}
	if session.SessionHandler() != nil {return}

	e.performanceApp.ChangeBanch(&entities.Performance{
		CompanyId: session.CompanyId,
		PerformanceId: reqBody.PerformanceId,
		BanchId: reqBody.BanchId,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}