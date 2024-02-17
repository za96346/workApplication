package controller

import (
	"backend/application/services"

	"github.com/gin-gonic/gin"
)

type EntryController struct {
	entryApp application.EntryAppInterface
}

func NewEntry() *EntryController {
	return &EntryController{
		entryApp: &application.EntryApp{},
	}
}

func (e *EntryController) Login(c *gin.Context) {
}