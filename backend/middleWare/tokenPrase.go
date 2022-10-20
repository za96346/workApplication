package middleWare

import (
	"backend/handler"
	panichandler "backend/panicHandler"
	// "fmt"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenPrase(props *gin.Context) {
	defer panichandler.Recover()
	method := props.Request.Method
	tokenParams := ""
	if method == "GET" {
		tokenParams = (*props).Query("token")
		
		// log.Println("token => ", tokenParams)
	} else {
		tokenParams =  (*props).GetHeader("token")
		
		// log.Println("token => ", tokenParams)
	}

	// 判斷 token 是否過期
	if !handler.Singleton().Redis.IsTokenExited(tokenParams) {
		(*props).AbortWithStatusJSON(http.StatusNotExtended, gin.H{
			"message": "token Expire",
		})
		return
	}
	
	// 解析 token
	userInfo, err := handler.ParseToken(tokenParams)
	if err != nil {
		(*props).AbortWithStatusJSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "not Allow",
		})
		return
	}

	(*handler.Singleton()).Redis.ResetExpireTime(tokenParams)
	
	(*props).Set("Account", userInfo["Account"].(string))
	(*props).Set("UserId", int64(userInfo["UserId"].(float64)))
	(*props).Next()
}