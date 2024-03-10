package controller

import (
	"backend/application/services"
	"backend/interfaces/method"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
	systemApp application.SystemAppInterface
}

func NewSystem(app application.SystemAppInterface) *SystemController {
	return &SystemController{
		systemApp: app,
	}
}

func (e *SystemController) GetAuth(Request *gin.Context) {
	session, err := method.NewSession(
		Request,
		&method.ReqStruct{},
	)
	if err != nil {return}

	functionItem, permission, appErr := e.systemApp.GetAuth(session)
	permissionToJson, _ := json.Marshal(permission)
	session.SessionInstance.Set("permission", string(permissionToJson))

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "失敗",
				"data": nil,
			},
		)
		return
	}

	session.SessionInstance.Save()

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": map[string]interface{} {
				"menu": *functionItem,
				"permission": *permission,
			},
		},
	)
}

func (e *SystemController) GetFunc(Request *gin.Context) {
	functionItem, operationItem, functionRoleBanchRelationMap := e.systemApp.GetFunc()

	// 整理 return data
	result := make(map[string]interface{})

	result["functionItem"] = *functionItem
	result["operationItem"] = *operationItem
	result["functionRoleBanchRelation"] = *functionRoleBanchRelationMap

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": result,
		},
	)
}

func (e *SystemController) GetRoleBanchList(Request *gin.Context) {
	session, err := method.NewSession(
		Request,
		&method.ReqStruct{},
	)
	if err != nil {return}

	scopeRole, scopeBanch, scopeUser, availableBanch, availableRole, availableUser, appErr := e.systemApp.GetRoleBanchList(session)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "失敗",
				"data": nil,
			},
		)
	}

	data := make(map[string]interface{})
	data["scopeRole"] = scopeRole
	data["scopeBanch"] = scopeBanch
	data["scopeUser"] = scopeUser
	data["availableBanch"] = availableBanch
	data["availableRole"] = availableRole
	data["availableUser"] = availableUser

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": data,
		},
	)
}