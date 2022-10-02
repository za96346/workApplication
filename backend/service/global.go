package service

import (
	"backend/handler"
	"backend/panicHandler"
)

var dbHandle = handler.Singleton()
var panicHandle = panichandler.Recover

func IsExited[T any](value *[]T) bool {
	if len(*value) > 0 {
		return true
	}
	return false
}