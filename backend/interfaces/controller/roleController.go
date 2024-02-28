package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"backend/infrastructure/persistence"
	"backend/interfaces/enum"
	"backend/interfaces/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	repo *persistence.Repositories
	roleApp application.RoleAppInterface
}

func NewRole(repo *persistence.Repositories) *RoleController {
	return &RoleController{
		repo: repo,
		roleApp: &application.RoleApp{},
	}
}

func (e *RoleController) GetRoles(Request *gin.Context) {
	// 請求處理
	reqParams := new(struct {
		RoleName string `json:"RoleName"`
	})

	// 權限驗證
	session := method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		ReqParamsValidation: true,
		ReqParamsStruct: reqParams,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.RoleManage),
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler() != nil {return}

	data, _ := e.roleApp.GetRoles(&entities.Role{
		CompanyId: session.CompanyId,
		RoleName: reqParams.RoleName,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":  data,
		},
	)
}

func (e *RoleController) GetRole(Request *gin.Context) {
	// 請求處理
	reqParams := new(struct {
		RoleId int `json:"RoleId" binding:"required"`
	})

	// 權限驗證
	session := method.SessionStruct{
		Request: Request,
		ReqParamsValidation: true,
		ReqParamsStruct: reqParams,


		PermissionValidation: true,
		PermissionFuncCode: string(enum.RoleManage),
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler() != nil {return}

	role, rolePermissionMap := e.roleApp.GetRole(&entities.Role{
		CompanyId: session.CompanyId,
		RoleId: reqParams.RoleId,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":  map[string]interface{}{
				"Role": *role,
				"Permission": *rolePermissionMap,
			},
		},
	)
}

func (e *RoleController) GetRolesSelector(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		PermissionValidation: false,
	}
	if session.SessionHandler() != nil {return}

	data, _ := e.roleApp.GetRolesSelector(&entities.Role{
		CompanyId: session.CompanyId,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": *data,
		},
	)
}

func (e *RoleController) UpdateRole(Request *gin.Context) {
	// 請求處理
	reqBody := new(struct {
		RoleId int `json:"RoleId" binding:"required"`
		RoleName string `json:"RoleName" binding:"required"`
		Sort *int `json:"Sort"`
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

		PermissionValidation: true,
		PermissionFuncCode: string(enum.RoleManage),
		PermissionItemCode: "edit",
	}

	if session.SessionHandler() != nil {return}

	e.roleApp.UpdateRole(&entities.Role{
		CompanyId: session.CompanyId,
		RoleId: reqBody.RoleId,
		RoleName: reqBody.RoleName,
		Sort: reqBody.Sort,
	}, &reqBody.Data)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

func (e *RoleController) SaveRole(Request *gin.Context) {
	// 請求處理
	reqBody := new(struct {
		RoleName string `json:"RoleName" binding:"required"`
		Sort *int `json:"Sort"`
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

		PermissionValidation: true,
		PermissionFuncCode: string(enum.RoleManage),
		PermissionItemCode: "add",
	}

	e.roleApp.SaveRole(&entities.Role{
		CompanyId: session.CompanyId,
	}, &reqBody.Data)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "新增成功",
		},
	)
}

func (e *RoleController) DeleteRole(Request *gin.Context) {
	// 請求處理
	reqBody := new(struct {
		RoleId int `json:"RoleId" binding:"required"`
	})

	// 權限驗證
	session := method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.RoleManage),
		PermissionItemCode: "delete",
	}
	if session.SessionHandler() != nil {return}

	e.roleApp.DeleteRole(&entities.Role{
		CompanyId: session.CompanyId,
		RoleId: reqBody.RoleId,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}
