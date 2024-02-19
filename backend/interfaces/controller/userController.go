package controller

import (
	"backend/application/services"
	"backend/domain/entities"
	"backend/interfaces/enum"
	"backend/interfaces/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userApp application.UserAppInterface
}

func NewUser() *UserController {
	return &UserController{
		userApp: &application.UserApp{},
	}
}

func (u *UserController) GetMine(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
	}
	if session.SessionHandler() != nil {return}

	data, _ := u.userApp.GetMine(&entities.User{
		UserId: session.UserId,
		CompanyId: session.CompanyId,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    *data,
		},
	)
}

func (u *UserController) GetUsers(Request *gin.Context) {
	// 請求處理
	reqParams := new(struct {
		BanchId *int `json:"BanchId"`
		RoleId *int `json:"RoleId"`
		UserName *string `json:"UserName"`
		EmployeeNumber *string `json:"EmployeeNumber"`
		QuitFlag *string `json:"QuitFlag"`
	})
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,

		PermissionValidation: true,
		PermissionFuncCode: string(enum.EmployeeManage),
		PermissionItemCode: "inquire",

		ReqBodyValidation: false,
		ReqParamsValidation: true,
		ReqParamsStruct: reqParams,
	}
	if session.SessionHandler() != nil {return}

	data, _ := u.userApp.GetUsers(
		&entities.User{
			CompanyId: session.CompanyId,
			BanchId: reqParams.BanchId,
			RoleId: *reqParams.RoleId,
			UserName: *reqParams.UserName,
			EmployeeNumber: *reqParams.EmployeeNumber,
			QuitFlag: *reqParams.QuitFlag,
		},
		session.GetScopeBanchWithCustomize(reqParams.BanchId),
		session.GetScopeRolehWithCustomize(reqParams.RoleId),
	)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    *data,
		},
	)
}

func (u *UserController) GetUsersSelector(Request *gin.Context) {
	// 請求處理
	reqParams := new(struct {
		BanchId *int `json:"BanchId"`
		RoleId *int `json:"RoleId"`
		UserName *string `json:"UserName"`
		EmployeeNumber *string `json:"EmployeeNumber"`
	})
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqParamsValidation: true,
		ReqParamsStruct: reqParams,
	}
	if session.SessionHandler() != nil {return}

	data, _ := u.userApp.GetUsersSelector(&entities.User{
		CompanyId: session.CompanyId,
		BanchId: reqParams.BanchId,
		RoleId: *reqParams.RoleId,
		UserName: *reqParams.UserName,
		EmployeeNumber:  *reqParams.EmployeeNumber,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    data,
		},
	)
}

func (u *UserController) UpdateUser(Request *gin.Context) {
	reqBody := new(entities.User)

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.EmployeeManage),
		PermissionItemCode: "edit",
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}
	if session.CheckScopeBanchValidation(*(*reqBody).BanchId) != nil {return}
	if session.CheckScopeRoleValidation((*reqBody).RoleId) != nil {return}

	// 檢驗欄位
	if reqBody.UserId == 0 {
		// ErrorInstance.ErrorHandler(Request, "更新失敗，UserId is nil.")
		return
	}

	(*reqBody).CompanyId = session.CompanyId
	u.userApp.UpdateUser(reqBody)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

func (u *UserController) UpdatePassword(Request *gin.Context) {
	reqBody := new(struct{
		OldPassword string `json:"OldPassword" binding:"required"`
		NewPassword string `json:"NewPassword" binding:"required"`
		NewPasswordAgain string `json:"NewPasswordAgain" binding:"required"`
		UserId int `gorm:"column:userId;primaryKey" json:"UserId" binding:"required"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.EmployeeManage),
		PermissionItemCode: "edit",
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

func (u *UserController) UpdateMine(Request *gin.Context) {
	reqBody := new(struct{
		UserName string `json:"UserName" binding:"required"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.SelfData),
		PermissionItemCode: "edit",
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}

	u.userApp.UpdateMine(&entities.User{
		CompanyId: session.CompanyId,
		UserId: session.UserId,
		UserName: reqBody.UserName,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

func (u *UserController) SaveUser(Request *gin.Context) {
	reqBody := new(entities.User)
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.EmployeeManage),
		PermissionItemCode: "add",
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}
	if session.CheckScopeBanchValidation(*(*reqBody).BanchId) != nil {return}
	if session.CheckScopeRoleValidation((*reqBody).RoleId) != nil {return}

	reqBody.CompanyId = session.CompanyId
	u.userApp.SaveUser(reqBody)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "新增成功",
		},
	)
}

func (u *UserController) DeleteUser(Request *gin.Context) {
	reqBody := new(struct {
		UserId int `json:"UserId" binding:"required"`
	})

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		PermissionValidation: true,
		PermissionFuncCode: string(enum.EmployeeManage),
		PermissionItemCode: "delete",
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
	}
	if session.SessionHandler() != nil {return}

	u.userApp.DeleteUser(&entities.User{
		CompanyId: session.CompanyId,
		UserId: reqBody.UserId,
	})

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}