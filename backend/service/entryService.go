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
	if IsNotExited(res) {
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

	registeForm := &response.Register{}
	err1 := (*props).ShouldBindBodyWith(&registeForm, binding.JSON)
	// 檢查格式
	if err1 != nil  {
		(*props).JSON(http.StatusExpectationFailed, gin.H{
			"message": StatusText().RegisterFailNotAcceptDataFormat,
		})
		fmt.Println(err1)
		return
	}

	// 是否是電子信箱的格式
	if !handler.VerifyEmailFormat((*registeForm).Account) {
		(*props).JSON(http.StatusUnavailableForLegalReasons, gin.H{
			"message": StatusText().EmailIsNotRight,
		})
		return
	}

	// 檢查帳號是否被註冊
	res := (*dbHandle).SelectUser(2, (*registeForm).Account)
	if !IsNotExited(res) {
		(*props).JSON(http.StatusConflict, gin.H{
			"message": StatusText().AccountHasBeenRegisted,
		})
		return
	}

	// 檢查驗證碼是否正確
	rightCaptcha := (*dbHandle).Redis.SelectEmailCaptcha((*registeForm).Account)
	if (*registeForm).Captcha != rightCaptcha || rightCaptcha == -1 {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().EmailCaptchaIsNotRight,
		})
		return
	}

	// 密碼是否相等
	if (*registeForm).Password != (*registeForm).PasswordConfirm {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().PasswordIsNotSame,
		})
		return
	}

	// 新增使用者

	user := &table.UserTable{
		Account: (*registeForm).Account,
		Password: (*registeForm).Password,
		CompanyCode: (*registeForm).CompanyCode,
		UserName: (*registeForm).UserName,
		Permession: 2,
		WorkState: "on",
		Banch: -1,
		MonthSalary: 0,
		PartTimeSalary: 0,
		OnWorkDay: now,
		CreateTime: now,
		LastModify: now,
	}
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