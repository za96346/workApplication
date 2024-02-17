package controller

import (
	"backend/application/services"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleApp application.RoleAppInterface
}

func NewRole() *RoleController {
	return &RoleController{
		roleApp: &application.RoleApp{},
	}
}

func (e *RoleController) GetRole(c *gin.Context) {
}

func (e *RoleController) GetRoles(c *gin.Context) {
}

func (e *RoleController) GetRolesSelector(c *gin.Context) {
}

func (e *RoleController) UpdateRole(c *gin.Context) {
}

func (e *RoleController) SaveRole(c *gin.Context) {
}

func (e *RoleController) DeleteRole(c *gin.Context) {
}
