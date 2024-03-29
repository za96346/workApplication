package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"backend/interfaces/method"
	"backend/application/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PerformanceController struct {
	performanceApp application.PerformanceAppInterface
}

func NewPerformance(app application.PerformanceAppInterface) *PerformanceController {
	return &PerformanceController{
		performanceApp: app,
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
			BanchId: reqParams.BanchId,
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
		return
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
		return
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
		return
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "新增成功",
		},
	)
}

func (e *PerformanceController) CopyPerformance(Request *gin.Context) {
	reqBody := new(entities.Performance)

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := e.performanceApp.CopyPerformance(reqBody, session)

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
		return
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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := e.performanceApp.DeletePerformance(
		&entities.Performance{
			PerformanceId: reqBody.PerformanceId,
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

func (e *PerformanceController) ChangeBanch(Request *gin.Context) {
	reqBody := new(struct {
		PerformanceId   int         `json:"PerformanceId" binding:"required"`
		BanchId         int         `json:"BanchId" binding:"required"`
	})

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := e.performanceApp.ChangeBanch(
		&entities.Performance{
			PerformanceId: reqBody.PerformanceId,
			BanchId: reqBody.BanchId,
		},
		session,
	)

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