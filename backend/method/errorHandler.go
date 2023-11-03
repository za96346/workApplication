package method

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorStruct struct {
	MessageTitle string
}

// 錯誤處理
func (instance *ErrorStruct) ErrorHandler(Request *gin.Context, MSG string) {
	Request.JSON(
		http.StatusForbidden,
		gin.H {
			"message": instance.MessageTitle + MSG,
		},
	)
}