package CTL_User

import (
	"backend/Model"
	"backend/method"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_User 使用者]--",
}

const FuncCode = "employeeManage"

// 尋找自己的 使用者資料
func GetMine(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
	}
	if session.SessionHandler() != nil {return}

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

// 獲取使用者 全部
func Get(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "inquire",
		ReqBodyValidation: false,
	}
	if session.SessionHandler() != nil {return}

	var data []Model.User
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("roleId in (?)", session.CurrentPermissionScopeRole).
		Where("banchId in (?)", session.CurrentPermissionScopeBanch).
		Where("deleteFlag", "N").
		Find(&data)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    data,
		},
	)
}

// 新增
func Add(Request *gin.Context) {
	reqBody := new(Model.User)
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "add",
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}
	if session.CheckScopeBanchValidation(*(*reqBody).BanchId) != nil {return}
	if session.CheckScopeRoleValidation((*reqBody).RoleId) != nil {return}

	// 加入一些固定欄位
	now := time.Now()

	(*reqBody).GetNewUserID(session.CompanyId)

	(*reqBody).CreateTime = &now
	(*reqBody).LastModify = &now
	(*reqBody).DeleteTime = nil
	(*reqBody).DeleteFlag = "N"

	// 插入 Recorder
	if Model.DB.Create(reqBody).Error != nil {
		ErrorInstance.ErrorHandler(Request, "新增失敗")
		return
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "新增成功",
		},
	)
}

/*
	編輯
	這邊可能要 再加上 原本此使用者的 部門驗證以及 角色驗證
*/
func Edit(Request *gin.Context) {
	reqBody := new(Model.User)
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "edit",
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}
	if session.CheckScopeBanchValidation(*(*reqBody).BanchId) != nil {return}
	if session.CheckScopeRoleValidation((*reqBody).RoleId) != nil {return}

	// 檢驗欄位
	if reqBody.UserId == 0 {
		ErrorInstance.ErrorHandler(Request, "更新失敗，UserId is nil.")
		return
	}

	//共同 語句
	commonQuery := Model.DB.
		Model(&Model.User{}).
		Where("companyId = ?", session.CompanyId).
		Where("userId = ?", reqBody.UserId)

	// 找到舊的值 ( 不讓請求 的時候 userId 有任何的串改可能． )
	var oldData Model.User 
	commonQuery.First(&oldData)

	// 加入一些固定欄位
	now := time.Now()

	(*reqBody).Account = oldData.Account
	(*reqBody).CompanyId = session.CompanyId
	(*reqBody).LastModify = &now
	(*reqBody).DeleteTime = nil
	(*reqBody).DeleteFlag = "N"

	// 更新
	err := commonQuery.Updates(&reqBody).Error
	if err != nil {
		ErrorInstance.ErrorHandler(Request, "更新失敗")
		return
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

/*
	刪除
*/
func Delete(Request *gin.Context) {
	reqBody := new(struct {
		UserId int `json:"UserId" binding:"required"`
	})

	targetData := new(Model.User)

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "delete",
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}

	// 獲取此比資料
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("userId = ?", reqBody.UserId).
		First(targetData)

	// 此資料scope 驗證
	if session.CheckScopeBanchValidation(*(*targetData).BanchId) != nil {return}
	if session.CheckScopeRoleValidation((*targetData).RoleId) != nil {return}

	// 加入一些固定欄位
	now := time.Now()

	(*targetData).LastModify = &now
	(*targetData).DeleteTime = &now
	(*targetData).DeleteFlag = "Y"

	err := Model.DB.
		Model(&Model.User{}).
		Where("companyId = ?", session.CompanyId).
		Where("userId = ?", reqBody.UserId).
		Updates(&targetData).
		Error

	if err != nil {
		ErrorInstance.ErrorHandler(Request, "刪除失敗")
		return
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}