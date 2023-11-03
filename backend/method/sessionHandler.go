package method

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var ErrorInstance = &ErrorStruct{
	MessageTitle: "[session fail]--",
}

type SessionStruct struct {
	CompanyId int // 公司id
	IsLogin bool // 是否成功登入 "Y" | "N"
	UserId int
	RoleId int
	BanchId int
	UserName string
	EmployeeNumber string
}

func(instance *SessionStruct) SessionHandler(Request *gin.Context) error {
	session := sessions.Default(Request)
	session.Set("companyId", "0")

	// 公司 id
	companyId, err := strconv.Atoi(session.Get("companyId").(string))
	if err != nil {
		ErrorInstance.ErrorHandler(Request, "公司id Error")
		return err
	}

	// 是否登入
	isLogin := false
	if session.Get("isLogin") == "Y" {isLogin = true}

	// 使用者id
	userId, err := strconv.Atoi(session.Get("userId").(string))
	if err != nil {
		ErrorInstance.ErrorHandler(Request, "使用者id Error")
		return err
	}

	//角色id
	roleId, err := strconv.Atoi(session.Get("roleId").(string))
	if err != nil {
		ErrorInstance.ErrorHandler(Request, "使用者角色id Error")
		return err
	}

	// 部門id
	banchId, err := strconv.Atoi(session.Get("banchId").(string))
	if err != nil {
		ErrorInstance.ErrorHandler(Request, "使用者部門id Error")
		return err
	}

	// 使用者姓名
	userName := session.Get("userName").(string)

	// 使用者 員工編號
	employeeNumber := session.Get("employeeNumber").(string)
	
	// 綁定物件
	(*instance).CompanyId = companyId
	(*instance).IsLogin = isLogin
	(*instance).UserId = userId
	(*instance).RoleId = roleId
	(*instance).BanchId = banchId
	(*instance).UserName = userName
	(*instance).EmployeeNumber = employeeNumber

	return nil
}