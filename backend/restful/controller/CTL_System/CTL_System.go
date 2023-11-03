package CTL_System

import (
	"net/http"

	"backend/method"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_System 系統]--",
}

// 獲取權限
func GetAuth(Request *gin.Context) {
	session := &method.SessionStruct{}
	if session.SessionHandler(Request) != nil {return}

	// 回傳資料
	response := map[string]interface{} {
		"session": *session,
		"menu": "",
		"permission": "",
	}

	// Model.DB.Model(&Model.)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": response,
		},
	)
}