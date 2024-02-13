package Middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
	權限檢查 middle
*/
func Permission(c *gin.Context)  {
	session := sessions.Default(c)
	if session.Get("isLogin") != "Y" {
		c.String(http.StatusNetworkAuthenticationRequired, "fail")
		c.Abort()
		return
	}
	c.Next()
}