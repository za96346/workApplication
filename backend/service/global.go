package service

import (
	"backend/handler"
	"backend/panicHandler"
)

var dbHandle = handler.Singleton()
var panicHandle = panichandler.Recover

func StatusText(num int) string {
	switch num {
	case 0:
		return "登入失敗 請輸入有效的資料"
	case 1:
		return "沒有此使用者"
	case 2:
		return "帳號或密碼錯誤"
	case 3:
		return "註冊失敗， 資料不正確"
	case 4:
		return "此帳號已經被註冊了"
	case 5:
		return "註冊失敗"
	case 6:
		return "註冊成功"
	default:
		return ""
	}
}

func IsExited[T any](value *[]T) bool {
	if len(*value) > 0 {
		return true
	}
	return false
}