package service

import (
	"backend/handler"
	"backend/table"
	"time"

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
		(*props).JSON(http.StatusExpectationFailed, gin.H{
			"message": StatusText().LoginFail,
		})
		return
	}

	// 檢查帳號是否存在
	res := (*dbHandle).SelectUser(2, (*reqBody).Account)
	if !IsExited(res) {
		(*props).JSON(http.StatusUnauthorized, gin.H{
			"message": StatusText().NoUser,
		})
		return
	}

	// 檢查帳號密碼是否正確
	if (*res)[0].Account != (reqBody).Account || (*res)[0].Password != (reqBody).Password {
		//登入失敗
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().AccountOrPasswordError,
		})
		return
	}

	//登入成功
	tk := handler.Token {
		UserId: (*res)[0].UserId,
		Account: (*res)[0].Account,
	}
	(*dbHandle).Redis.InsertToken(tk.GetLoginToken())
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().LoginSuccess,
		"token": tk.GetLoginToken(),
	})
	
}
func Register(props *gin.Context, waitJob *sync.WaitGroup){
	defer panicHandle()
	defer (*waitJob).Done()
	now := time.Now()
	user := &table.UserTable{
		CreateTime: now,
		LastModify: now,
	}

	// 檢查格式
	if (*props).ShouldBindJSON(&user) != nil {
		(*props).JSON(http.StatusExpectationFailed, gin.H{
			"message": StatusText().RegisterFailNotAcceptDataFormat,
		})
		return
	}

	// 檢查帳號是否被註冊
	res := (*dbHandle).SelectUser(2, user.UserId)
	if IsExited(res) {
		(*props).JSON(http.StatusConflict, gin.H{
			"message": StatusText().AccountHasBeenRegisted,
		})
		return
	}

	// 新增
	status, _ := (*dbHandle).InsertUser(user)
	if !status {
		// 註冊失敗
		(*props).JSON(http.StatusForbidden, gin.H{
			"message": StatusText().RegisterFail,
		})
		return
	}

	// 註冊成功
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().RegisterSuccess,
	})
}