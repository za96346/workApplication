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
}

func ConvertSliceToInt(in []any) (out []int) {
    out = make([]int, 0, len(in))
    for _, v := range in {
        out = append(out, int(v.(float64)))
    }
    return
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
				Request,
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