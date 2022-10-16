package middleWare

import (
	"backend/handler"
	"backend/methods"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyCompanyAndBanch(props *gin.Context) {
	userId, existed := (*props).Get("UserId")
	// user id 是否存在
	if !existed {
		(*props).AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "找不到使用者id",
		})
		return
	}

	// 轉換 user id
	convertUserId, err := methods.AnyToInt64(userId)
	// fmt.Print(convertUserId, err)
	if err != nil {
		(*props).AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "找不到使用者id",
		})
		return
	}
	
	// 尋找自己資料
	myUserData := (*handler.Singleton()).SelectUser(1, convertUserId)
	if methods.IsNotExited(myUserData) {
		(*props).AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "找不到使用者資料",
		})
		return
	}

	// 尋找公司
	company := (*handler.Singleton()).SelectCompany(2, (*myUserData)[0].CompanyCode)
	if methods.IsNotExited(company) {
		(*props).AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "尚未有公司",
		})
		return
	}
	(*props).Set("MyUserData", (*myUserData)[0])
	(*props).Set("MyCompany", (*company)[0])
	(*props).Next()
}