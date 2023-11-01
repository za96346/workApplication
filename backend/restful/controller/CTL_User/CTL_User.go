package CTL_User

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 尋找自己的 使用者資料
func GetMine(Request *gin.Context) {
	Request.JSON(
		http.StatusAccepted,
		gin.H {
			"message": "成功",
			"data":    "hi",
		},
	)
}
