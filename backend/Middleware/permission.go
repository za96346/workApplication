package Middleware

import (
	// "net/http"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
	權限檢查 middle
*/
func Permission(c *gin.Context)  {
	// session := sessions.Default(c)
	// if 1 == 2 {
	// 	c.String(http.StatusForbidden, "rate limit...")
	// 	c.Abort()
	// 	return
	// }
	c.Next()
}