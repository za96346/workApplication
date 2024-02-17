package controller

import (
	"backend/application/services"

	"github.com/gin-gonic/gin"
)

type PerformanceController struct {
	performanceApp application.PerformanceAppInterface
}

func NewPerformance() *PerformanceController {
	return &PerformanceController{
		performanceApp: &application.PerformanceApp{},
	}
}

func (e *PerformanceController) GetPerformances(c *gin.Context) {
}

func (e *PerformanceController) GetYearPerformances(c *gin.Context) {
}

func (e *PerformanceController) UpdatePerformance(c *gin.Context) {
}

func (e *PerformanceController) SavePerformance(c *gin.Context) {
}

func (e *PerformanceController) DeletePerformance(c *gin.Context) {
}

func (e *PerformanceController) ChangeBanch(c *gin.Context) {
}