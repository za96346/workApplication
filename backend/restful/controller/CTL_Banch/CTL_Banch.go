package CTL_Banch

import (
	"backend/Model"
	"backend/method"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ErrorInstance = &method.ErrorStruct{
	MessageTitle: "[CTL_Banch 公司部門]--",
}

const FuncCode = "banchManage"

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

	var responseData []Model.CompanyBanch

	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("banchId in (?)", session.CurrentPermissionScopeBanch).
		Where("deleteFlag = ?", "N").
		Find(&responseData)

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
	reqBody := new(Model.CompanyBanch)
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
	if session.CheckScopeBanchValidation((*reqBody).BanchId) != nil {return}

	// 添加固定欄位
	now := time.Now()

	(*reqBody).LastModify = &now
	(*reqBody).DeleteTime = nil
	(*reqBody).DeleteFlag = "N"

	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("banchId = ?", reqBody.BanchId).
		Updates(reqBody)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

// 新增
func Add(Request *gin.Context) {
	reqBody := new(Model.CompanyBanch)
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "add",
	}
	if session.SessionHandler() != nil {return}

	// 添加固定欄位
	now := time.Now()
	(*reqBody).GetNewBanchID(session.CompanyId)

	(*reqBody).LastModify = &now
	(*reqBody).CreateTime = &now
	(*reqBody).DeleteTime = nil
	(*reqBody).DeleteFlag = "N"

	// 新增
	if 	Model.DB.Create(reqBody).Error != nil {
		ErrorInstance.ErrorHandler(Request, "新增失敗")
		return
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "新增成功",
		},
	)
}

// 新增
func Delete(Request *gin.Context) {
	reqBody := new(struct {
		BanchId int `json:"BanchId" binding:"required"`
	})

	targetData := new(Model.CompanyBanch)

	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: true,
		ReqBodyStruct: reqBody,
		PermissionValidation: true,
		PermissionFuncCode: FuncCode,
		PermissionItemCode: "delete",
	}
	if session.SessionHandler() != nil {return}
	if session.CheckScopeBanchValidation((*reqBody).BanchId) != nil {return}

	// 拿取此筆資料
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("banchId = ?", reqBody.BanchId).
		First(targetData)

	// 加入固定欄位
	now := time.Now()
	(*targetData).DeleteFlag = "Y"
	(*targetData).DeleteTime = &now
	(*targetData).LastModify = &now
	
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Where("banchId = ?", reqBody.BanchId).
		Updates(targetData)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "刪除成功",
		},
	)
}

func GetSelector(Request *gin.Context) {
	// 權限驗證
	session := &method.SessionStruct{
		Request: Request,
		ReqBodyValidation: false,
		PermissionValidation: false,
	}
	if session.SessionHandler() != nil {return}

	// 獲取部門
	var targetData []Model.CompanyBanch
	Model.DB.
		Where("companyId = ?", session.CompanyId).
		Find(&targetData)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data": targetData,
		},
	)
}