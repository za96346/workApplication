package controller

import (
	"backend/application/services"
	"backend/infrastructure/persistence"
	"backend/interfaces/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
	repo *persistence.Repositories
	systemApp application.SystemAppInterface
}

func NewSystem(repo *persistence.Repositories) *SystemController {
	return &SystemController{
		repo: repo,
		systemApp: &application.SystemApp{},
	}
}

func (e *SystemController) GetAuth(Request *gin.Context) {
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
	}
	if session.SessionHandler() != nil {return}

	functionItem, permission := e.systemApp.GetAuth(
		session.CompanyId,
		session.RoleId,
	)

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
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
	}
	if session.SessionHandler() != nil {return}

	scopeRole, scopeBanch, scopeUser, availableBanch, availableRole, availableUser := e.systemApp.GetRoleBanchList(
		session.CompanyId,
		&session.CurrentPermissionScopeBanch,
		&session.CurrentPermissionScopeRole,
	)

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