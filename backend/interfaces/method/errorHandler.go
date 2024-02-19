package method

import (
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
)

type ErrorStruct struct {
	MessageTitle string
}

// 錯誤處理
func (instance *ErrorStruct) ErrorHandler(Request *gin.Context, MSG string) {
	session := SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		PermissionValidation: false,
	}
	session.SessionHandler()

	// now := time.Now()

	// Log := Model.Log{
	// 	UserId: session.UserId,
	// 	Routes: "[" + Request.Request.Method + "]" + Request.FullPath(),
	// 	Ip: Request.ClientIP(),
	// 	Msg: &MSG,
	// 	CreateTime: &now,
	// 	LastModify: &now,
	// }
	// Log.GetNewLogId(session.CompanyId)
	// Model.DB.Create(Log)
	Request.JSON(
		http.StatusForbidden,
		gin.H {
			"message": instance.MessageTitle + MSG,
		},
	)
}