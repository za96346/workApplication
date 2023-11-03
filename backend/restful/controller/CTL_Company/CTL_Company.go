package CTL_Company

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(Request *gin.Context) {
	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    "company get",
		},
	)
}