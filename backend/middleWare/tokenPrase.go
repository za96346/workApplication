package middleWare

import (
	"backend/handler"
	panichandler "backend/panicHandler"
	"backend/table"

	// "fmt"
	// "log"
	"net/http"

	"backend/redis"

	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"

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
	if !(*redis.Singleton()).IsTokenExited(tokenParams) {
		(*props).AbortWithStatusJSON(http.StatusNotExtended, gin.H{
			"message": "登入憑證過期 ， 請重新登入",
		})
		return
	}
	// 解析 token
	userInfo, err := handler.ParseToken(tokenParams)
	if err != nil {
		(*props).AbortWithStatusJSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "登入憑證錯誤",
		})
		return
	}

	(*redis.Singleton()).ResetExpireTime(tokenParams)
	user := new(table.UserTable)
	company := new(table.CompanyTable)

	mapstructure.Decode(userInfo["User"], user)
	mapstructure.Decode(userInfo["Company"], company)
	
	(*props).Set("Account", (*user).Account)
	(*props).Set("UserId", (*user).UserId)
	(*props).Set("CompanyCode", (*user).CompanyCode)
	(*props).Set("UserName", (*user).UserName)
	(*props).Set("EmployeeNumber", (*user).EmployeeNumber)
	(*props).Set("OnWorkDay", (*user).OnWorkDay)
	(*props).Set("Banch", (*user).Banch)
	(*props).Set("Permession", (*user).Permession)
	(*props).Set("CompanyId", (*company).CompanyId)
	(*props).Set("BossId", (*company).BossId)

	(*props).Next()
}
