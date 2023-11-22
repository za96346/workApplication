package CTL_Company

import (
	"net/http"
	"time"

	"backend/Model"
	"backend/method"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_Company 公司]--",
}

const FuncCode = "companyData"

// 拿取
func Get(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "inquire",
	}
	if session.SessionHandler() != nil {return}

	var responseData Model.Company

	Model.DB.
		Where("companyId = ?", session.CompanyId).
		First(&responseData)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":    responseData,
		},
	)
}

// 編輯
func Edit(Request *gin.Context) {
	reqBody := new(Model.Company)
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "edit",
	}
	if session.SessionHandler() != nil {return}

	now := time.Now()
	(*reqBody).LastModify = &now
	(*reqBody).CompanyId = session.CompanyId

	err := Model.DB.
		Where("companyId = ?", session.CompanyId).
		Updates(reqBody).
		Error

	if err != nil {
		
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}