package middleWare

import (
	panichandler "backend/panicHandler"
	"strconv"

	// "fmt"
	"backend/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SameCompany(props *gin.Context) {
	defer panichandler.Recover()

	// 判斷 有沒有userid
	userId, existed := (*props).Get("UserId")
	if !existed {
		(*props).AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "找不到使用者, user id is illegal",
		})
		return
	}

	// 判斷id 是否有效
	targetUserId := (*props).Query("userId")
	n, err := strconv.ParseInt(targetUserId, 10, 64)
	if err != nil || userId == nil {
		(*props).AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "無效id",
		})
		return
	}

	// 選擇兩人的資料
	me := (*handler.Singleton()).SelectUser(1, userId.(int64))
 	iWantGetPeople := (*handler.Singleton()).SelectUser(1, n)

	// 是否有資料
	if len(*me) == 0 || len(*iWantGetPeople) == 0 {
		(*props).AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "找不到使用者",
		})
		return
	}

	// 比對公司
	if (*me)[0].CompanyCode != (*iWantGetPeople)[0].CompanyCode && (*iWantGetPeople)[0].CompanyCode != "" {
		(*props).AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "公司不同",
		})
		return
	}
	(*props).Set("targetUser", (*iWantGetPeople)[0])
	(*props).Next()
}