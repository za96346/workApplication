package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"backend/infrastructure/persistence"
	"backend/interfaces/enum"
	"backend/interfaces/method"
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
	reqParams := new(struct{
		BanchId *int `json:"BanchId"`
		RoleId *int `json:"RoleId"`
		UserName *string `json:"UserName"`
		StartDate *string `json:"StartDate"`
		EndDate *string `json:"EndDate"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		ReqParamsValidation: true,
		ReqParamsStruct: reqParams,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.Performance),
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler() != nil {return}

	data, _ := e.performanceApp.GetPerformances(&entities.Performance{
		CompanyId: session.CompanyId,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    *data,
		},
	)
}

func (e *PerformanceController) GetYearPerformances(Request *gin.Context) {
	reqParams := new(struct{
		BanchId *int `json:"BanchId"`
		RoleId *int `json:"RoleId"`
		UserName *string `json:"UserName"`
		StartYear *string `json:"StartYear"`
		EndYear *string `json:"EndYear"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		ReqParamsValidation: true,
		ReqParamsStruct: reqParams,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.Performance),
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler() != nil {return}

	data, _ := e.performanceApp.GetYearPerformances(&entities.Performance{
		CompanyId: session.CompanyId,
	})

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
	PermissionItemCode := "add"
	if strings.Contains(Request.Request.URL.Path, "copy") {
		PermissionItemCode = "copy"
	}

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.Performance),
		PermissionItemCode: PermissionItemCode,
	}
	if session.SessionHandler() != nil {return}

	reqBody.CompanyId = session.CompanyId
	e.performanceApp.SavePerformance(reqBody)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "新增成功",
		},
	)
}

func (e *PerformanceController) UpdatePerformance(Request *gin.Context) {
	reqBody := new(entities.Performance)

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.Performance),
		PermissionItemCode: "edit",
	}
	if session.SessionHandler() != nil {return}

	reqBody.CompanyId = session.CompanyId
	e.performanceApp.UpdatePerformance(reqBody)

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