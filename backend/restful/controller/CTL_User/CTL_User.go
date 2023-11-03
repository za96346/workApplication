package CTL_User

import (
	"backend/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 尋找自己的 使用者資料
func GetMine(Request *gin.Context) {
	var data *Model.User
	Model.DB.First(&data)
	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    *data,
		},
	)
}

// 獲取