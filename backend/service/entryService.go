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
			"message": StatusText(0),
		})
		return
	}

	// 檢查帳號是否存在
	res := (*dbHandle).SelectUser(2, (*reqBody).Account)
	if !IsExited(res) {
		(*props).JSON(http.StatusUnauthorized, gin.H{
			"message": StatusText(1),
		})
		return
	}

	// 檢查帳號密碼是否正確
	if (*res)[0].Account == (reqBody).Account {
		if (*res)[0].Password == (reqBody).Password {
			tk := handler.Token {
				UserId: (*res)[0].UserId,
				Account: (*res)[0].Account,
			}
			(*dbHandle).Redis.InsertToken(tk.GetLoginToken())
			(*props).JSON(http.StatusOK, tk.GetLoginToken())
			return
		}
	}
	(*props).JSON(http.StatusBadRequest, gin.H{
		"message": StatusText(2),
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
			"message": StatusText(3),
		})
		return
	}

	// 檢查帳號是否被註冊
	res := (*dbHandle).SelectUser(2, user.UserId)
	if IsExited(res) {
		(*props).JSON(http.StatusConflict, gin.H{
			"message": StatusText(4),
		})
		return
	}

	// 新增
	status, _ := (*dbHandle).InsertUser(user)
	if !status {
		(*props).JSON(http.StatusForbidden, gin.H{
			"message": StatusText(5),
		})
		return
	}
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText(6),
	})
}