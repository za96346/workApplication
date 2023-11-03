package CTL_Role

import (
	"backend/Model"
	"encoding/json"
	"net/http"
	"time"

	"backend/method"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_Role 角色]--",
}


func checkRequest() {
	
}

// 獲取公司角色
func Get(Request *gin.Context) {
	// session
	session := method.SessionStruct{}
	if session.SessionHandler(Request) != nil {return}

	data := new([]Model.Role)
	Model.DB.Where("companyId = ?", session.CompanyId).Where("deleteFlag = ?", "N").Find(data)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":  data,
		},
	)
}

// 獲取公司單一角色
func GetSingle(Request *gin.Context) {
	roleData := &Model.Role{}
	rolePermission := &[]Model.RoleStruct{}
	rolePermissionMap := map[string](map[string][]int){}

	// session
	session := method.SessionStruct{}
	if session.SessionHandler(Request) != nil {return}

	// 請求處理
	reqBody := new(struct {
		RoleId int
	})

	if Request.ShouldBindJSON(&reqBody) != nil {
		ErrorInstance.ErrorHandler(Request, "Request Data 格式不正確")
		return
	}

	// 查詢DB
	Model.DB.Where("companyId = ?", session.CompanyId).Where("roleId = ?", reqBody.RoleId).Where("deleteFlag = ?", "N").First(roleData)
	Model.DB.Where("companyId = ?", session.CompanyId).Where("roleId = ?", reqBody.RoleId).Find(rolePermission)

	for _, v := range *rolePermission {
		if rolePermissionMap[v.FuncCode] == nil {
			rolePermissionMap[v.FuncCode] = make(map[string][]int)
		}
		rolePermissionMap[v.FuncCode][v.ItemCode] = []int{}
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":  map[string]interface{}{
				"Role": *roleData,
				"Permission": rolePermissionMap,
			},
		},
	)
}

// 更新角色結構
func Update(Request *gin.Context) {
	TX := Model.DB.Begin()

	session := method.SessionStruct{}
	if session.SessionHandler(Request) != nil {return}


	// 請求處理
	reqBody := new(struct {
		RoleId int
		RoleName string
		StopFlag string
		Type string
		/**
			Data = {
				[funcCode]: {
					[itemCode]: {
						scopeBanch: []BanchId, 
						scopeRole: []RoleId
					}
				}
			}
		*/
		Data map[string](map[string](map[string][]int))

	})

	if Request.ShouldBindJSON(&reqBody) != nil {
		ErrorInstance.ErrorHandler(Request, "Request Data 格式不正確")
		TX.Rollback()
		return
	}

	// 要更新的欄位
	updateRoleQuery := map[string]interface{}{
		"roleName": reqBody.RoleName,
		"stopFlag": reqBody.StopFlag,
		"lastModify": time.Now(),
	}

	// 更新 或 新增 role table
	if reqBody.Type == "add" {
		var MaxCount int64
		TX.Model(&Model.Role{}).Where("companyId = ?", session.CompanyId).Count(&MaxCount)
		updateRoleQuery["companyId"] = session.CompanyId
		updateRoleQuery["roleId"] = MaxCount + 1

		TX.Model(&Model.Role{}).Create(&updateRoleQuery)
	} else {
		
		TX.Model(&Model.Role{}).Where(
			"companyId = ?",
			session.CompanyId,
		).Where(
			"roleId = ?",
			reqBody.RoleId,
		).Updates(&updateRoleQuery)
	}

	// 先把 此role structure 的資料 刪除
	TX.Where(
		"companyId = ?",
		session.CompanyId,
	).Where(
		"roleId = ?",
		reqBody.RoleId,
	).Delete(&Model.RoleStruct{})

	// 在 寫入 新的 進入 db
	for funcCode, itemObject := range reqBody.Data {
		for itemCode, scopeObject := range itemObject {
			scopeBanch, _:= json.Marshal(scopeObject["scopeBanch"])
			scopeRole, _:= json.Marshal(scopeObject["scopeRole"])
			updateData := &Model.RoleStruct{
				CompanyId: session.CompanyId,
				RoleId: reqBody.RoleId,
				FuncCode: funcCode,
				ItemCode: itemCode,
				ScopeBanch: string(scopeBanch),
				ScopeRole: string(scopeRole),
				CreateTime: time.Now(),
				LastModify: time.Now(),
			}
			if TX.Create(updateData).Error != nil {
				ErrorInstance.ErrorHandler(Request, "新增失敗")
				TX.Rollback()
				return
			}
			
		}
	}

	TX.Commit()
	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

// 刪除角色
func Delete(Request *gin.Context) {
	TX := Model.DB.Begin()

	session := method.SessionStruct{}
	if session.SessionHandler(Request) != nil {return}

	// 請求處理
	reqBody := new(struct {
		RoleId int
	})

	if Request.ShouldBindJSON(&reqBody) != nil {
		ErrorInstance.ErrorHandler(Request, "Request Data 格式不正確")
		TX.Rollback()
		return
	}

	// 要更新的欄位
	updateRoleQuery := map[string]interface{}{
		"deleteFlag": "Y",
		"deleteTime": time.Now(),
		"lastModify": time.Now(),
	}

	TX.Model(&Model.Role{}).Where(
		"companyId = ?",
		0,
	).Where(
		"roleId = ?",
		reqBody.RoleId,
	).Updates(&updateRoleQuery)

	TX.Commit()
	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}