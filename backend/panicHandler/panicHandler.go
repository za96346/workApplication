package panichandler
import (
	"fmt"
)

func Recover() {
	if err := recover(); err != nil {
	   fmt.Println(err) // 这里的err其实就是panic传入的内容
	}
 }