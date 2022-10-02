package service

import (
	"backend/handler"
	"backend/table"
	// "fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func Login(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	reqBody := new(struct {
		Account string
		Password string
	})

	// 檢查格式
	if (*props).ShouldBindJSON(&reqBody) != nil {
		(*props).JSON(http.StatusBadRequest, "登入失敗 請輸入有效的資料")
		return
	}

	// 檢查帳號是否存在
	res := (*dbHandle).SelectUser(2, (*reqBody).Account)
	if !IsExited(res) {
		(*props).JSON(http.StatusOK, "沒有此使用者")
		return
	}

	// 檢查帳號密碼是否正確
	if (*res)[0].Account == (reqBody).Account {
		if (*res)[0].Password == (reqBody).Password {
			tk := handler.Token {
				UserId: (*res)[0].UserId,
				Account: (*res)[0].Account,
			}
			(*props).JSON(http.StatusOK, "登入成功" + tk.GetLoginToken())
			return
		}
	}
	(*props).JSON(http.StatusBadRequest, "帳號或密碼錯誤")
	
}
func Register(props *gin.Context, waitJob *sync.WaitGroup){
	defer panicHandle()
	defer (*waitJob).Done()
	user := new(table.UserTable)

	// 檢查格式
	if (*props).ShouldBindJSON(&user) != nil {
		(*props).JSON(http.StatusBadRequest, "註冊失敗， 資料不正確")
		return
	}

	// 檢查帳號是否被註冊
	res := (*dbHandle).SelectUser(2, user.UserId)
	if IsExited(res) {
		(*props).JSON(http.StatusConflict, "此帳號已經被註冊了")
		return
	}

	// 新增
	status, _ := (*dbHandle).InsertUser(user)
	if !status {
		(*props).JSON(http.StatusForbidden, "註冊失敗")
		return
	}
	(*props).JSON(http.StatusOK, "註冊成功")
}