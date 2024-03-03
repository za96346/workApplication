package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"backend/infrastructure/persistence"
	"backend/interfaces/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	repo *persistence.Repositories
	roleApp application.RoleAppInterface
}

func NewRole(repo *persistence.Repositories, app application.RoleAppInterface) *RoleController {
	return &RoleController{
		repo: repo,
		roleApp: app,
	}
}

func (e *RoleController) GetRoles(Request *gin.Context) {
	// 請求處理
	reqParams := new(struct {
		RoleName string `json:"RoleName"`
	})

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqParamsValidation: true,
			ReqParamsStruct: reqParams,
		},
	)
	if err != nil {return}

	data, appErr := e.roleApp.GetRoles(
		&entities.Role{
			RoleName: reqParams.RoleName,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusOK,
			gin.H {
				"message": "失敗",
				"data":  data,
			},
		)
		return
	}

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqParamsValidation: true,
			ReqParamsStruct: reqParams,
		},
	)
	if err != nil {return}

	role, rolePermissionMap, appErr := e.roleApp.GetRole(
		&entities.Role{
			RoleId: reqParams.RoleId,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "失敗",
				"data":  nil,
			},
		)
		return
	}

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
	session, err := method.NewSession(
		Request,
		&method.ReqStruct{},
	)
	if err != nil {return}

	data, appErr := e.roleApp.GetRolesSelector(
		&entities.Role{},
		session,
	)

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := e.roleApp.UpdateRole(
		&entities.Role{
			RoleId: reqBody.RoleId,
			RoleName: reqBody.RoleName,
			Sort: reqBody.Sort,
		},
		&reqBody.Data,
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusOK,
			gin.H {
				"message": "更新失敗",
			},
		)
		return
	}

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := e.roleApp.SaveRole(
		&entities.Role{},
		&reqBody.Data,
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "新增失敗",
			},
		)
		return
	}

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := e.roleApp.DeleteRole(
		&entities.Role{
			RoleId: reqBody.RoleId,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "刪除失敗",
			},
		)
		return
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}
