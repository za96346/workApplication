package service

import (
	"backend/handler"
	"backend/methods"
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
	res := (*Mysql).SelectUser(2, (*reqBody).Account)
	if methods.IsNotExited(res) {
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

	company := *(new([]table.CompanyTable))
	findCompany := (*Mysql).SelectCompany(2, (*res)[0].CompanyCode)
	if methods.IsNotExited(findCompany) {
		company = append(company, *new(table.CompanyTable))
	} else {
		company = (*findCompany)
	}

	//登入成功
	tk := handler.Token {
		User: (*res)[0],
		Company: company[0],
	}

	(*Redis).InsertToken(tk.GetLoginToken())
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().LoginSuccess,
		"data": tk.GetLoginToken(),
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
	res := (*Mysql).SelectUser(2, (*registeForm).Account)
	if !methods.IsNotExited(res) {
		(*props).JSON(http.StatusConflict, gin.H{
			"message": StatusText().AccountHasBeenRegisted,
		})
		return
	}

	// 檢查公司碼 是否存在
	if (*registeForm).CompanyCode != "" {
		company := (*Mysql).SelectCompany(2, (*registeForm).CompanyCode)
		if !methods.IsNotExited(company) {
			(*props).JSON(http.StatusConflict, gin.H{
				"message": StatusText().CompanyCodeIsNotRight,
			})
			return
		}
	}
	// 檢查驗證碼是否正確
	rightCaptcha := (*Redis).SelectEmailCaptcha((*registeForm).Account)
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
	if len((*registeForm).Password) < 8 {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().PasswordNotSafe,
		})
		return
	}
	// 新增使用者

	user := &table.UserTable{
		Account: (*registeForm).Account,
		Password: (*registeForm).Password,
		CompanyCode: (*registeForm).CompanyCode,
		UserName: (*registeForm).UserName,
		EmployeeNumber: "",
		Permession: 2,
		Banch: -1,
		MonthSalary: 0,
		PartTimeSalary: 0,
		OnWorkDay: now,
		CreateTime: now,
		LastModify: now,
	}
	status, _ := (*Mysql).InsertUser(user)
	if !status {
		// 註冊失敗
		(*props).JSON(http.StatusForbidden, gin.H{
			"message": StatusText().RegisterFail,
		})
		return
	}


	// 註冊成功 把captcha 刪掉
	(*Redis).DeleteCaptcha((*registeForm).Account)
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