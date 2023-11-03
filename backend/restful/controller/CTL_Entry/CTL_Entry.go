package CTL_Entry

import (
	"net/http"
	"strconv"

	"backend/Model"
	"backend/method"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_Entry 入口]--",
}

// 登入
func Login(Request *gin.Context) {
	session := sessions.Default(Request)

	user := &Model.User{}

	// 請求處理
	reqBody := new(struct {
		Password string
		Account string
	})

	if Request.ShouldBindJSON(&reqBody) != nil {
		ErrorInstance.ErrorHandler(Request, "Request Data 格式不正確")
		return
	}

	// 驗證帳號密碼

	Model.DB.Where("account", reqBody.Account).First(user)

	if !(
		(*user).Password == reqBody.Password && 
		(*user).Account == reqBody.Account &&
		(*user).Account != "" &&
		(*user).Password != "") {

		ErrorInstance.ErrorHandler(
			Request,
			"帳號或密碼錯誤",
		)
		return
	}

	// 登入成功後 ， 寫入session
	session.Set("isLogin", "Y")
	session.Set("userId", strconv.Itoa((*user).UserId))
	session.Set("companyId", strconv.Itoa((*user).CompanyId))
	session.Set("roleId", strconv.Itoa((*user).RoleId))
	session.Set("banchId", strconv.Itoa((*user).BanchId))
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