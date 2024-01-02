package CTL_Performance

import (
	"backend/Model"
	"backend/method"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_Performance 績效]--",
}

const FuncCode = "performance"

// 尋找
func Get(Request *gin.Context) {
	reqParams := new(struct{
		BanchId *int `json:"BanchId"`
		RoleId *int `json:"RoleId"`
		UserName *string `json:"UserName"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		ReqParamsValidation: true,
		ReqParamsStruct: reqParams,

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
	searchQuery := Model.DB.
		Model(&Model.Performance{}).
		Where("performance.companyId = ?", session.CompanyId).
		Where("user.banchId in (?)", *session.GetScopeBanchWithCustomize(reqParams.BanchId)).
		Where("user.roleId in (?)", *session.GetScopeRolehWithCustomize(reqParams.RoleId)).
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
		Select(
			"performance.*",
			"user.userName as userName",
			"company_banch.banchName as banchName",
			"user.roleId",
		)

	// 使用者名稱
	if reqParams.UserName != nil {
		searchQuery.Where("userName like ?", "%" + *reqParams.UserName + "%")
	}

	searchQuery.Find(&data)
	

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

	PermissionItemCode := "add"
	if strings.Contains(Request.Request.URL.Path, "copy") {
		PermissionItemCode = "copy"
	}

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: PermissionItemCode,
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
	(*reqBody).BanchId = *userData.BanchId
	(*reqBody).DeleteFlag = "N"
	(*reqBody).DeleteTime = nil
	(*reqBody).CreateTime = &now
	(*reqBody).LastModify = &now

	if (*reqBody).IsYearMonthDuplicated() {
		ErrorInstance.ErrorHandler(Request, "新增失敗-檢查到重複資料")
		return
	}

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

	//共同 語句
	commonQuery := Model.DB.
		Model(&Model.Performance{}).
		Where("companyId = ?", session.CompanyId).
		Where("performanceId = ?", reqBody.PerformanceId)

	// 找到舊的值 ( 不讓請求 的時候 userId 有任何的串改可能． )
	var oldData Model.Performance 
	commonQuery.First(&oldData)

	// 新增固定欄位
	now := time.Now()

	(*reqBody).CompanyId = session.CompanyId
	(*reqBody).UserId = oldData.UserId
	(*reqBody).BanchId = oldData.BanchId
	(*reqBody).DeleteFlag = "N"
	(*reqBody).DeleteTime = nil
	(*reqBody).LastModify = &now

	if (*reqBody).IsYearMonthDuplicated() {
		ErrorInstance.ErrorHandler(Request, "新增失敗-檢查到重複資料")
		return
	}

	// 更新
	err := commonQuery.Updates(reqBody).Error

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
	
	err := Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("performanceId = ?", reqBody.PerformanceId).
		Updates(targetData).
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

// 尋找年度績效
func GetYear(Request *gin.Context) {
	reqParams := new(struct{
		BanchId *int `json:"BanchId"`
		RoleId *int `json:"RoleId"`
		UserName *string `json:"UserName"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		ReqParamsValidation: true,
		ReqParamsStruct: reqParams,

		PermissionValidation: true,
		PermissionFuncCode: "yearPerformance",
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler() != nil {return}

	// 獲取資料
	var data []struct{
		Year int `gorm:"column:year" json:"Year"`
		UserName string  `gorm:"column:userName" json:"UserName"`
		Score float32 `gorm:"column:score" json:"Score"`
	}
	searchQuery := Model.DB.
		Model(&Model.Performance{}).
		Where("performance.companyId = ?", session.CompanyId).
		Where("performance.banchId in (?)", *session.GetScopeBanchWithCustomize(reqParams.BanchId)).
		Where("user.roleId in (?)", *session.GetScopeRolehWithCustomize(reqParams.RoleId)).
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
		Group("performance.userId").
		Group("performance.year").
		Group("user.userName").
		Order("score desc").
		Select(
			"performance.year as year",
			"user.userName as userName",
			`
				round(
					(
						sum(performance.attitude)
						+ sum(performance.efficiency)
						+ sum(performance.professional)
					) / 36, 2
				) as score
			`,
		)

	// 使用者名稱
	if reqParams.UserName != nil {
		searchQuery.Where("user.userName like ?", "%" + *reqParams.UserName + "%")
	}

	searchQuery.Find(&data)
	

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    data,
		},
	)
}

// 更換部門
// func ChangeBanch(Request *gin.Context) {
// 	reqBody := new(struct {
// 		PerformanceId   int         `json:"PerformanceId"`
// 		BanchId         int         `json:"BanchId"`
// 	})

// 	// 權限驗證
// 	session := &method.SessionStruct{
// 		Request: Request,
// 		ReqBodyValidation: true,
// 		ReqBodyStruct: reqBody,

// 		PermissionValidation: true,
// 		PermissionFuncCode: FuncCode,
// 		PermissionItemCode: "edit",
// 	}
// 	if session.SessionHandler() != nil {return}

// 	// 查詢此 user 資料
// 	userData := Model.User{}
// 	var count int64

// 	userQuery := Model.DB.
// 		Model(&Model.User{}).
// 		Where("userId = ?", reqBody.UserId).
// 		Where("companyId = ?", session.CompanyId)

// 	userQuery.Count(&count)
// 	userQuery.First(&userData)
// 	if count == int64(0) {
// 		ErrorInstance.ErrorHandler(Request, "找不到此使用者")
// 		return
// 	}

// 	// 檢查是否有此部門以及角色的權限
// 	if session.CheckScopeBanchValidation(*userData.BanchId) != nil {return}
// 	if session.CheckScopeRoleValidation(userData.RoleId) != nil {return}

// 	//共同 語句
// 	commonQuery := Model.DB.
// 		Model(&Model.Performance{}).
// 		Where("companyId = ?", session.CompanyId).
// 		Where("performanceId = ?", reqBody.PerformanceId)

// 	// 找到舊的值 ( 不讓請求 的時候 userId 有任何的串改可能． )
// 	var oldData Model.Performance 
// 	commonQuery.First(&oldData)

// 	// 新增固定欄位
// 	now := time.Now()

// 	(*reqBody).CompanyId = session.CompanyId
// 	(*reqBody).UserId = oldData.UserId
// 	(*reqBody).BanchId = oldData.BanchId
// 	(*reqBody).DeleteFlag = "N"
// 	(*reqBody).DeleteTime = nil
// 	(*reqBody).LastModify = &now

// 	if (*reqBody).IsYearMonthDuplicated() {
// 		ErrorInstance.ErrorHandler(Request, "新增失敗-檢查到重複資料")
// 		return
// 	}

// 	// 更新
// 	err := commonQuery.Updates(reqBody).Error

// 	if err != nil {
// 		ErrorInstance.ErrorHandler(Request, "更新失敗")
// 		return
// 	}

// 	Request.JSON(
// 		http.StatusOK,
// 		gin.H {
// 			"message": "更新成功",
// 		},
// 	)
// }