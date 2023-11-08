package CTL_Performance

import (
	"backend/Model"
	"backend/method"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_Performance 績效]--",
}

const FuncCode = "performance"

// 尋找
func Get(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,

		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler() != nil {return}

	// 獲取資料
	var data []struct{
		Model.Performance
		BanchName string  `gorm:"column:banchName" json:"BanchName"`
		UserName string  `gorm:"column:userName" json:"UserName"`
	}
	Model.DB.
		Model(&Model.Performance{}).
		Where("performance.companyId = ?", session.CompanyId).
		Where("performance.banchId = ?", session.CurrentPermissionScopeBanch).
		Where("user.roleId in (?)", session.CurrentPermissionScopeRole).
		Where("performance.deleteFlag = ?", "N").
		Joins(`
			left join user
			on user.userId = performance.userId
			and user.companyId = performance.companyId
		`).
		Joins(`
			left join company_banch
			on company_banch.companyId = performance.companyId
			and company_banch.banchId = performance.banchId
		`).
		Select(`
			performance.*,
			user.userName,
			company_banch.banchName
		`).
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
	reqBody := new(Model.Performance)

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "add",
	}
	if session.SessionHandler() != nil {return}

	// 查詢此 user 資料
	userData := Model.User{}
	var count int64

	userQuery := Model.DB.
		Model(&Model.User{}).
		Where("userId = ?", reqBody.UserId).
		Where("companyId = ?", session.CompanyId)

	userQuery.Count(&count)
	userQuery.First(&userData)
	if count == int64(0) {
		ErrorInstance.ErrorHandler(Request, "找不到此使用者")
		return
	}

	// 檢查是否有此部門以及角色的權限
	if session.CheckScopeBanchValidation(*userData.BanchId) != nil {return}
	if session.CheckScopeRoleValidation(userData.RoleId) != nil {return}

	// 新增固定欄位
	now := time.Now()
	(*reqBody).GetNewPerformanceID(session.CompanyId)
	(*reqBody).DeleteFlag = "N"
	(*reqBody).DeleteTime = nil
	(*reqBody).CreateTime = &now
	(*reqBody).LastModify = &now

	// 新增資料
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

// 編輯
func Edit(Request *gin.Context) {
	reqBody := new(Model.Performance)

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "edit",
	}
	if session.SessionHandler() != nil {return}

	// 查詢此 user 資料
	userData := Model.User{}
	var count int64

	userQuery := Model.DB.
		Model(&Model.User{}).
		Where("userId = ?", reqBody.UserId).
		Where("companyId = ?", session.CompanyId)

	userQuery.Count(&count)
	userQuery.First(&userData)
	if count == int64(0) {
		ErrorInstance.ErrorHandler(Request, "找不到此使用者")
		return
	}

	// 檢查是否有此部門以及角色的權限
	if session.CheckScopeBanchValidation(*userData.BanchId) != nil {return}
	if session.CheckScopeRoleValidation(userData.RoleId) != nil {return}

	// 新增固定欄位
	now := time.Now()
	(*reqBody).DeleteFlag = "N"
	(*reqBody).DeleteTime = nil
	(*reqBody).LastModify = &now

	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("performanceId = ?", reqBody.PerformanceId).
		Updates(reqBody)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

// 刪除
func Delete(Request *gin.Context) {
	reqBody := new(struct {
		PerformanceId int `gorm:"column:performanceId;primaryKey" json:"PerformanceId"`
	})

	targetData := new(Model.Performance)
	targetUserData := new(Model.User)

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "delete",
	}
	if session.SessionHandler() != nil {return}

	// 獲取此筆績效
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("performanceId = ?", reqBody.PerformanceId).
		First(targetData)

	// 獲取此筆績效 使用者
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("userId = ?", targetData.UserId).
		First(targetUserData)
	
	// 檢查是否有此部門以及角色的權限
	if session.CheckScopeBanchValidation(*targetUserData.BanchId) != nil {return}
	if session.CheckScopeRoleValidation(targetUserData.RoleId) != nil {return}


	// 加入固定欄位
	now := time.Now()

	(*targetData).DeleteFlag = "Y"
	(*targetData).DeleteTime = &now
	(*targetData).LastModify = &now
	
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("performanceId = ?", reqBody.PerformanceId).
		Updates(targetData)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}

// 搜尋列 的 值
func SearchBar(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		PermissionValidation: false,
	}
	if session.SessionHandler() != nil {return}


}