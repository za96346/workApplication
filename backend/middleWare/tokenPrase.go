package middleWare

import (
	"backend/handler"
	// "fmt"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenPrase(props *gin.Context) {
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
		props.AbortWithStatusJSON(http.StatusNotAcceptable, "token Expire")
		return
	}
	
	// 解析 token
	userInfo, err := handler.ParseToken(tokenParams)
	if err != nil {
		props.AbortWithStatusJSON(http.StatusNotAcceptable, "not Allow")
		return
	}
	
	props.Set("Account", userInfo["Account"].(string))
	props.Set("UserId", int64(userInfo["UserId"].(float64)))
	props.Next()
}