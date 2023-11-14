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
	// 權限驗證
	session := method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
	}
	if session.SessionHandler() != nil {return}

	data := new([]Model.Role)
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("deleteFlag = ?", "N").
		Find(data)

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
	rolePermissionMap := map[string](map[string]map[string]interface{}){}

	// 請求處理
	reqBody := new(struct {
		RoleId int `json:"RoleId" binding:"required"`
	})

	// 權限驗證
	session := method.SessionStruct{
		Request: Request,
		ReqParamsValidation: true,
		ReqParamsStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}

	// 查詢DB
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("roleId = ?", reqBody.RoleId).
		Where("deleteFlag = ?", "N").
		First(roleData)

	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("roleId = ?", reqBody.RoleId).
		Find(rolePermission)

	for _, v := range *rolePermission {
		if rolePermissionMap[v.FuncCode] == nil {
			rolePermissionMap[v.FuncCode] = make(map[string]map[string]interface{})
		}
		if rolePermissionMap[v.FuncCode][v.ItemCode] == nil {
			rolePermissionMap[v.FuncCode][v.ItemCode] = make(map[string]interface{})
		}

		// handle scope banch
		if v.ScopeBanch != "all" && v.ScopeBanch != "self" {
			var scopeBanch []int
			json.Unmarshal([]byte(v.ScopeBanch), &scopeBanch)
			rolePermissionMap[v.FuncCode][v.ItemCode]["scopeBanch"] = scopeBanch
		} else {
			rolePermissionMap[v.FuncCode][v.ItemCode]["scopeBanch"] = v.ScopeBanch
		}

		// handle scope role
		if v.ScopeRole != "all" && v.ScopeRole != "self" {
			var scopeRole []int
			json.Unmarshal([]byte(v.ScopeRole), &scopeRole)
			rolePermissionMap[v.FuncCode][v.ItemCode]["scopeRole"] = scopeRole
		} else {
			rolePermissionMap[v.FuncCode][v.ItemCode]["scopeRole"] = v.ScopeRole
		}
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

	// 請求處理
	reqBody := new(struct {
		RoleId int `json:"RoleId" binding:"required"`
		RoleName string `json:"RoleName" binding:"required"`
		StopFlag string `json:"StopFlag" binding:"required"`
		Type string `json:"Type" binding:"required"`
		/**
			Data = {
				[funcCode]: {
					[itemCode]: {
						scopeBanch: []BanchId | all | self, 
						scopeRole: []RoleId | all | self,
					}
				}
			}
		*/
		Data map[string](map[string](map[string]interface{})) `json:"Data" binding:"required"`

	})

	// 權限驗證
	session := method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {
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
		TX.Model(&Model.Role{}).
			Where("companyId = ?", session.CompanyId).
			Count(&MaxCount)
		updateRoleQuery["companyId"] = session.CompanyId
		updateRoleQuery["roleId"] = MaxCount + 1

		TX.Model(&Model.Role{}).Create(&updateRoleQuery)
	} else {
		
		TX.Model(&Model.Role{}).
			Where("companyId = ?", session.CompanyId).
			Where("roleId = ?", reqBody.RoleId).
			Updates(&updateRoleQuery)
	}

	// 先把 此role structure 的資料 刪除
	TX.
		Where("companyId = ?", session.CompanyId).
		Where("roleId = ?", reqBody.RoleId).
		Delete(&Model.RoleStruct{})

	now := time.Now()

	// 在 寫入 新的 進入 db
	for funcCode, itemObject := range reqBody.Data {
		for itemCode, scopeObject := range itemObject {

			// 可編輯部門範圍
			scopeBanch := ""
			if scopeObject["scopeBanch"] == "all" || scopeObject["scopeBanch"] == "self" {
				scopeBanch = scopeObject["scopeBanch"].(string)
			} else {
				scopeBanchByte, _ := json.Marshal(scopeObject["scopeBanch"])
				scopeBanch = string(scopeBanchByte)
			}

			// 可編輯角色範圍
			scopeRole := ""
			if scopeObject["scopeRole"] == "all" || scopeObject["scopeRole"] == "self" {
				scopeRole = scopeObject["scopeRole"].(string)
			} else {
				scopeRoleByte, _ :=json.Marshal(scopeObject["scopeRole"])
				scopeRole = string(scopeRoleByte)
			}



			updateData := &Model.RoleStruct{
				CompanyId: session.CompanyId,
				RoleId: reqBody.RoleId,
				FuncCode: funcCode,
				ItemCode: itemCode,
				ScopeBanch: scopeBanch,
				ScopeRole: scopeRole,
				CreateTime: &now,
				LastModify: &now,
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

	// 請求處理
	reqBody := new(struct {
		RoleId int `json:"UserId" binding:"required"`
	})

	// 權限驗證
	session := method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {
		TX.Rollback()
		return
	}

	// 要更新的欄位
	updateRoleQuery := map[string]interface{}{
		"deleteFlag": "Y",
		"deleteTime": time.Now(),
		"lastModify": time.Now(),
	}

	TX.Model(&Model.Role{}).
		Where("companyId = ?", 0).
		Where("roleId = ?", reqBody.RoleId).
		Updates(&updateRoleQuery)

	TX.Commit()
	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}

func GetSelector(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		PermissionValidation: false,
	}
	if session.SessionHandler() != nil {return}

	// 獲取部門
	var targetData []Model.Role
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Find(&targetData)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": targetData,
		},
	)
}