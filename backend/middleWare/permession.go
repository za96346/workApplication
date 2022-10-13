package middleWare

import (
	"backend/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Permession(allowPermession ...int) gin.HandlerFunc {
	return func(props *gin.Context) {
		userId, existed := (*props).Get("UserId")
		// user id 尋找
		if !existed {
			(*props).AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "找不到使用者2",
			})
			return
		}

		// 尋找user
		user := (*handler.Singleton()).SelectUser(1, userId.(int64))
		if len(*user) == 0 {
			(*props).AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "找不到使用者3",
			})
			return
		}

		// 判斷權限
		count := 0
		for _, v := range allowPermession {
			if (*user)[0].Permession != v {
				count += 1
			}
		}
		if count == len(allowPermession) {
			(*props).AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "權限不足",
			})
			return
		}
		(*props).Set("CompanyCode", (*user)[0].CompanyCode)
		(*props).Next()
	}
}