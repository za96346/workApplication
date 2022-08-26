package middleWare
import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func permessionMiddleWare(account string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx)

        ctx.Next()
	}
}