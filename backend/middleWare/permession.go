package middleWare

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Permession(allowPermession ...int) gin.HandlerFunc {
	return func(props *gin.Context) {
		permession, _ := (*props).Get("Permession")

		// 判斷權限
		count := 0
		for _, v := range allowPermession {
			if permession != v {
				count += 1
			}
		}
		if count == len(allowPermession) {
			(*props).AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "權限不足",
			})
			return
		}
		(*props).Next()
	}
}