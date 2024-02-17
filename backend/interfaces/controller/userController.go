package controller

import (
	"backend/application/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userApp application.UserAppInterface
}

func NewUser() *UserController {
	return &UserController{
		userApp: &application.UserApp{},
	}
}

func (u *UserController) GetMine(c *gin.Context) {
}

func (u *UserController) GetUsers(c *gin.Context) {
}

func (u *UserController) GetUsersSelector(c *gin.Context) {
}

func (u *UserController) UpdateUser(c *gin.Context) {
}

func (u *UserController) UpdatePassword(c *gin.Context) {
}

func (u *UserController) UpdateMine(c *gin.Context) {
}

func (u *UserController) SaveUser(c *gin.Context) {
}

func (u *UserController) DeleteUser(c *gin.Context) {
}