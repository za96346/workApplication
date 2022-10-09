package service

import (
	"backend/handler"
	"backend/response"
	"backend/table"
	"fmt"
	"time"

	// "fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	var userCaptcha map[string]interface{}
	err1 := (*props).ShouldBindBodyWith(&userCaptcha, binding.JSON)
	err2 := (*props).ShouldBindBodyWith(&user, binding.JSON)
	// 檢查格式
	if err1 != nil || err2 != nil  {
		(*props).JSON(http.StatusExpectationFailed, gin.H{
			"message": StatusText().RegisterFailNotAcceptDataFormat,
		})
		fmt.Println(err1, err2)
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

	// 檢查驗證碼是否正確
	rightCaptcha := (*dbHandle).Redis.SelectEmailCaptcha((*user).Account)
	v := 0
	switch userCaptcha["Captcha"].(type) {
	case int:
		v = userCaptcha["Captcha"].(int)
		break
	case float64:
		v = int(userCaptcha["Captcha"].(float64))
		break
	default:
		fmt.Println("cap => ", userCaptcha["Captcha"], rightCaptcha)
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().EmailCaptchaIsNotRight,
		})
		return
	}
	if v != rightCaptcha || rightCaptcha == -1 {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().EmailCaptchaIsNotRight,
		})
		return
	}

	// 新增使用者
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

func EmailCaptcha(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	email := &response.OnlyEmail{}
	if (*props).ShouldBind(email) != nil {
		props.JSON(http.StatusOK, gin.H{
			"message": StatusText().RegisterFailNotAcceptDataFormat,
		})
		return
	}
	if !handler.SendEmail(email.Email) {
		props.JSON(http.StatusNotFound, gin.H{
			"message": StatusText().EmailCaptchaSendFail,
		})
		return
	}
	props.JSON(http.StatusOK, gin.H{
		"message": StatusText().EmailCaptchaSendSuccess,
	})
}

func CheckAccess(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	props.JSON(http.StatusAccepted, gin.H{
		"message": "身份認證成功",
	})
}