package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type EntryController struct {
	entryApp application.EntryAppInterface
}

func NewEntry() *EntryController {
	return &EntryController{
		entryApp: &application.EntryApp{},
	}
}

func (e *EntryController) Login(Request *gin.Context) {
	session := sessions.Default(Request)

	// 請求處理
	reqBody := new(struct {
		Password string
		Account string
	})

	if Request.ShouldBindJSON(&reqBody) != nil {
		// ErrorInstance.ErrorHandler(Request, "Request Data 格式不正確")
		return
	}

	user, _ := e.entryApp.Login(&entities.User{
		Account: reqBody.Account,
		Password: reqBody.Password,
	})

	// 加入 banch id 為 nil 的轉換
	if (*user).BanchId == nil {
		a := -1
		(*user).BanchId = &a
	}

	// 登入成功後 ， 寫入session
	session.Set("isLogin", "Y")
	session.Set("userId", strconv.Itoa((*user).UserId))
	session.Set("companyId", strconv.Itoa((*user).CompanyId))
	session.Set("roleId", strconv.Itoa((*user).RoleId))
	session.Set("banchId", strconv.Itoa(*(*user).BanchId))
	session.Set("userName", (*user).UserName)
	session.Set("employeeNumber", (*user).EmployeeNumber)

	session.Save()

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "登入成功",
		},
	)
}