package CTL_System

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/Model"
	"backend/method"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_System 系統]--",
}

// 獲取權限
func GetAuth(Request *gin.Context) {
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
	}
	if session.SessionHandler() != nil {return}

	RoleStruct := &[]Model.RoleStruct{}
	functionItem := &[]Model.FunctionItem{}

	// menu := &map[string]Model.FunctionItem{}
	permission := &map[string](map[string](map[string]interface{})){}

	// 拿取角色結構表
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("roleId = ?", session.RoleId).
		Find(RoleStruct)

	// 拿取功能項目表
	Model.DB.
		Where(
			"funcCode in (?)",
			Model.DB.
				Model(&Model.RoleStruct{}).
				Distinct().
				Select("funcCode").
				Where("companyId = ?", session.CompanyId).
				Where("roleId = ?", session.RoleId),
		).
		Find(functionItem)

    // 整理權限的資料結構
	for _, value := range *RoleStruct {
		// empty map init
		if (*permission)[value.FuncCode] == nil {
			(*permission)[value.FuncCode] = make(map[string]map[string]interface{})
		}

		if (*permission)[value.FuncCode][value.ItemCode] == nil {
			(*permission)[value.FuncCode][value.ItemCode] = make(map[string]interface{})
		}

		// 可編輯角色範圍 json decode
		if value.ScopeRole != "all" && value.ScopeRole != "self" {
			var scopeRole []int
			json.Unmarshal([]byte(value.ScopeRole), &scopeRole)
			(*permission)[value.FuncCode][value.ItemCode]["scopeRole"] = scopeRole
		} else {
			(*permission)[value.FuncCode][value.ItemCode]["scopeRole"] = value.ScopeRole
		}

		// 可編輯部門範圍 json decode
		if value.ScopeBanch != "all" && value.ScopeBanch != "self" {
			var scopeBanch []int
			json.Unmarshal([]byte(value.ScopeBanch), &scopeBanch)
			(*permission)[value.FuncCode][value.ItemCode]["scopeBanch"] = scopeBanch
		} else {
			(*permission)[value.FuncCode][value.ItemCode]["scopeBanch"] = value.ScopeBanch
		}
	}
	fmt.Printf("error1")
	// 權限寫入 session
	permissionToJson, _ := json.Marshal(permission)
	(*session.Instance).Set("permission", string(permissionToJson))

	(*session.Instance).Save()

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": map[string]interface{} {
				"menu": *functionItem,
				"permission": *permission,
				
				"session-BanchId": session.BanchId,
				"session-RoleId": session.RoleId,
				"session-UserId": session.UserId,
				"session-CompanyId": session.CompanyId,
				"session-IsLogin": session.IsLogin,
				"session-UserName": session.UserName,
				"session-EmployeeNumber": session.EmployeeNumber,
				"session-Permission": session.Permission,
			},
		},
	)

}