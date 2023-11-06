package method

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"backend/Model"
)

var ErrorInstance = &ErrorStruct{
	MessageTitle: "[session fail]--",
}

type SessionStruct struct {
	Instance *sessions.Session
	Request *gin.Context
	CompanyId int // 公司id
	IsLogin bool // 是否成功登入 "Y" | "N"
	UserId int
	RoleId int
	BanchId int
	UserName string
	EmployeeNumber string

	/**
		Permission = {
			[funcCode]: {
				[itemCode]: {
					scopeBanch: []BanchId | all | self, 
					scopeRole: []RoleId | all | self,
				}
			}
		}
	*/
	Permission map[string](map[string](map[string]interface{}))
	PermissionFuncCode string
	PermissionItemCode string
	CurrentPermission map[string]interface{} // 當前的權限
	CurrentPermissionScopeBanch []int // 當前的 scope banch
	CurrentPermissionScopeRole []int // 當前的 scope role
	PermissionValidation bool // 是否開啟 權限驗證

	IsCurrentScopeBanchAll bool // scope banch 是否是 all (給add 看的)
	IsCurrentScopeRoleAll bool // scope role 是否是 all (給add 看的)

	/*
		請求
	*/
	ReqBodyValidation bool // 是否開啟 請求 json binding 驗證
	ReqBodyStruct interface{} // 請求結構 ( please give it as a pointer. )
}



/*
	session 基本處理
*/
func(instance *SessionStruct) SessionHandler() error {
	session := sessions.Default((*instance).Request)
	session.Set("companyId", "0")
	(*instance).IsCurrentScopeBanchAll = false
	(*instance).IsCurrentScopeRoleAll = false

	// 公司 id
	companyId, err := strconv.Atoi(session.Get("companyId").(string))
	if err != nil {
		ErrorInstance.ErrorHandler((*instance).Request, "公司id Error")
		return err
	}

	// 是否登入
	isLogin := false
	if session.Get("isLogin") == "Y" {isLogin = true}

	// 使用者id
	userId, err := strconv.Atoi(session.Get("userId").(string))
	if err != nil {
		ErrorInstance.ErrorHandler((*instance).Request, "使用者id Error")
		return err
	}

	//角色id
	roleId, err := strconv.Atoi(session.Get("roleId").(string))
	if err != nil {
		ErrorInstance.ErrorHandler((*instance).Request, "使用者角色id Error")
		return err
	}

	// 部門id
	banchId, err := strconv.Atoi(session.Get("banchId").(string))
	if err != nil {
		ErrorInstance.ErrorHandler((*instance).Request, "使用者部門id Error")
		return err
	}

	// 使用者姓名
	userName := session.Get("userName").(string)

	// 使用者 員工編號
	employeeNumber := session.Get("employeeNumber").(string)

	// 權限json decode
	permission := session.Get("permission")
	if permission != nil {
		json.Unmarshal([]byte(permission.(string)), &(*instance).Permission)

		funcCode := (*instance).PermissionFuncCode
		itemCode := (*instance).PermissionItemCode

		permission, OK := (*instance).Permission[funcCode][itemCode]
		if !OK && (*instance).PermissionValidation {
			errMSG := fmt.Sprintf("權限驗證失敗--[funcCode: '%s'][itemCode: '%s']", funcCode, itemCode)
			ErrorInstance.ErrorHandler(
				(*instance).Request,
				errMSG,
			)
			return errors.New(errMSG)
		}

		// 可編輯角色範圍 的資料 搜尋 ( 分為自己，所有，自訂 )
		var scopeRole []int
		if permission["scopeRole"] == "all" {
			Model.DB.Model(&Model.Role{}).
				Select("roleId").
				Where("companyId = ?", companyId).
				Where("deleteFlag = ?", "N").
				Find(&Model.Role{}).
				Pluck("roleId", &scopeRole)

			// 設定 is current scope role all
			(*instance).IsCurrentScopeRoleAll = true
		} else if permission["scopeRole"] == "self" {
			scopeRole = append(scopeRole, roleId)
		} else if permission["scopeRole"] != nil {
			scopeRoleSlice := ConvertSliceToInt(
				permission["scopeRole"].([]any),
			)

			// 要把　自訂義裡面　可能被刪除的　roleId 過濾掉
			Model.DB.Model(&Model.Role{}).
				Select("roleId").
				Where("companyId = ?", companyId).
				Where("deleteFlag = ?", "N").
				Where("roleId in (?)", scopeRoleSlice).
				Find(&Model.Role{}).
				Pluck("roleId", &scopeRole)
		}

		// 可編輯部門範圍 的資料 搜尋 ( 分為自己，所有，自訂 )
		var scopeBanch  []int
		if permission["scopeBanch"] == "all" {
			Model.DB.Model(&Model.CompanyBanch{}).
				Select("banchId").
				Where("companyId = ?", companyId).
				Where("deleteFlag = ?", "N").
				Find(&Model.CompanyBanch{}).
				Pluck("banchId", &scopeBanch)
			
			// 設定 is current scope banch all
			(*instance).IsCurrentScopeBanchAll = true
		} else if permission["scopeBanch"] == "self" {
			scopeBanch = append(scopeBanch, banchId)
		} else if permission["scopeBanch"] != nil {
			scopeBanchSlice := ConvertSliceToInt(
				permission["scopeBanch"].([]any),
			)

			// 要把 自訂義裡面 可能被刪除的 banchId過濾掉
			Model.DB.Model(&Model.CompanyBanch{}).
				Select("banchId").
				Where("companyId = ?", companyId).
				Where("deleteFlag = ?", "N").
				Where("banchId in (?)", scopeBanchSlice).
				Find(&Model.CompanyBanch{}).
				Pluck("banchId", &scopeBanch)
		}

		// 綁定物件
		(*instance).CurrentPermission = permission
		(*instance).CurrentPermissionScopeBanch = scopeBanch
		(*instance).CurrentPermissionScopeRole = scopeRole
	}

	// 請求資料驗證
	if (*instance).ReqBodyValidation {
		bindError := (*instance).Request.ShouldBindJSON((*instance).ReqBodyStruct)
	
		if bindError != nil {
			ErrorInstance.ErrorHandler(
				(*instance).Request,
				fmt.Sprintf("Request Data 格式不正確 %s", bindError),
			)
			return bindError
		}	
	}

	// 綁定物件
	(*instance).CompanyId = companyId
	(*instance).IsLogin = isLogin
	(*instance).UserId = userId
	(*instance).RoleId = roleId
	(*instance).BanchId = banchId
	(*instance).UserName = userName
	(*instance).EmployeeNumber = employeeNumber
	(*instance).Instance = &session

	return nil
}

// 檢查可編輯的 部門範圍
func(instance *SessionStruct) CheckScopeBanchValidation(banchId int) error {
	// 檢查是否可以加入此部門
	if exists, _ := InArray(
		(*instance).CurrentPermissionScopeBanch,
		banchId);
		!exists {
		ErrorInstance.ErrorHandler((*instance).Request, "無法插入此部門，尚無權限")
		return errors.New("CheckScopeBanchValidation error.")
	}
	return nil
}

// 檢查可編輯的 角色範圍
func(instance *SessionStruct) CheckScopeRoleValidation(roleId int) error {
	// 檢查是否可以加入此角色
	if exists, _ := InArray(
		(*instance).CurrentPermissionScopeRole,
		roleId);
		!exists {
		ErrorInstance.ErrorHandler((*instance).Request, "無法插入此角色，尚無權限")
		return errors.New("CheckScopeRoleValidation error.")
	}

	return nil
}