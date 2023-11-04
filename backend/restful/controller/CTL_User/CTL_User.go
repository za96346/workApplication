package CTL_User

import (
	"backend/Model"
	"backend/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_User 使用者]--",
}

var FuncCode = "employeeManage"

// 尋找自己的 使用者資料
func GetMine(Request *gin.Context) {
	session := &method.SessionStruct{}
	if session.SessionHandler(Request) != nil {return}

	var data *Model.User
	Model.DB.
		Where("userId = ?", session.UserId).
		Where("companyId = ?", session.CompanyId).
		First(&data)

	// 清空密碼
	(*data).Password = ""

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    *data,
		},
	)
}

// 獲取公司 全部
func Get(Request *gin.Context) {
	session := &method.SessionStruct{
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler(Request) != nil {return}

	var data []Model.User
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("roleId in (?)", session.CurrentPermissionScopeRole).
		Where("banchId in (?)", session.CurrentPermissionScopeBanch).
		Find(&data)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    data,
		},
	)
}