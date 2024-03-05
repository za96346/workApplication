package controller

import (
	"backend/application/dtos"
	"backend/application/services"
	"backend/domain/entities"
	"backend/infrastructure/persistence"
	"backend/interfaces/method"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	repo *persistence.Repositories
	userApp application.UserAppInterface
}

func NewUser(app application.UserAppInterface) *UserController {
	return &UserController{
		userApp: app,
	}
}

func (u *UserController) GetMine(Request *gin.Context) {
	session, err := method.NewSession(
		Request,
		&method.ReqStruct{},
	)
	if err != nil {return}

	data, appErr := u.userApp.GetMine(
		&entities.User{},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "失敗",
				"data":    nil,
			},
		)
		return
	}

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqParamsValidation: true,
			ReqParamsStruct: reqParams,
		},
	)
	if err != nil {return}

	data, appErr := u.userApp.GetUsers(
		&entities.User{
			BanchId: reqParams.BanchId,
			RoleId: *reqParams.RoleId,
			UserName: *reqParams.UserName,
			EmployeeNumber: *reqParams.EmployeeNumber,
			QuitFlag: *reqParams.QuitFlag,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "失敗",
				"data":    nil,
			},
		)
		return
	}

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqParamsValidation: true,
			ReqParamsStruct: reqParams,
		},
	)
	if err != nil {return}

	data, appErr := u.userApp.GetUsersSelector(
		&entities.User{
			BanchId: reqParams.BanchId,
			RoleId: *reqParams.RoleId,
			UserName: *reqParams.UserName,
			EmployeeNumber:  *reqParams.EmployeeNumber,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
			gin.H {
				"message": "失敗",
				"data":    nil,
			},
		)
		return
	}

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

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := u.userApp.UpdateUser(reqBody, session)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
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

func (u *UserController) UpdatePassword(Request *gin.Context) {
	reqBody := new(dtos.UserPasswordUpdateQueryParams)

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := u.userApp.UpdatePassword(reqBody, session)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
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

func (u *UserController) UpdateMine(Request *gin.Context) {
	reqBody := new(struct{
		UserName string `json:"UserName" binding:"required"`
	})

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := u.userApp.UpdateMine(
		&entities.User{
			UserName: reqBody.UserName,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusBadRequest,
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

func (u *UserController) SaveUser(Request *gin.Context) {
	reqBody := new(entities.User)

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := u.userApp.SaveUser(reqBody, session)

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

func (u *UserController) DeleteUser(Request *gin.Context) {
	reqBody := new(struct {
		UserId int `json:"UserId" binding:"required"`
	})

	session, err := method.NewSession(
		Request,
		&method.ReqStruct{
			ReqBodyValidation: true,
			ReqBodyStruct: reqBody,
		},
	)
	if err != nil {return}

	_, appErr := u.userApp.DeleteUser(
		&entities.User{
			UserId: reqBody.UserId,
		},
		session,
	)

	if appErr != nil {
		Request.JSON(
			http.StatusOK,
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