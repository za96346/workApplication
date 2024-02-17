package controller

import (
	"backend/application/services"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
	systemApp application.SystemAppInterface
}

func NewSystem() *SystemController {
	return &SystemController{
		systemApp: &application.SystemApp{},
	}
}

func (e *SystemController) GetAuth(c *gin.Context) {
}

func (e *SystemController) GetFunc(c *gin.Context) {
}

func (e *SystemController) GetRoleBanchList(c *gin.Context) {
}