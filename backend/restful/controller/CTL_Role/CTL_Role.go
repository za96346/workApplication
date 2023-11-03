package CTL_Role

import (
	"backend/Model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 錯誤處理
func ErrorHandle(Request *gin.Context, TX *gorm.DB, MSG string) {
	Request.JSON(
		http.StatusForbidden,
		gin.H {
			"message": "[角色]--" + MSG,
		},
	)
	TX.Rollback()
}

func checkRequest() {
	
}

// 獲取公司角色
func Get(Request *gin.Context) {
	data := new([]Model.Role)
	Model.DB.Find(data)

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":  data,
		},
	)
}

// 獲取公司單一角色
func GetSingle(Request *gin.Context) {
	roleData := &Model.Role{}
	rolePermission := &[]Model.RoleStruct{}
	rolePermissionMap := map[string](map[string][]int){}

	session := sessions.Default(Request)

	companySID, _ := strconv.Atoi(session.Get("companySID").(string))

	// 請求處理
	reqBody := new(struct {
		RoleId int
	})

	if Request.ShouldBindJSON(&reqBody) != nil {
		ErrorHandle(Request, Model.DB.Begin(), "Request Data 格式不正確")
		return
	}

	// 查詢DB
	Model.DB.Where("companyId = ?", companySID).Where("roleId = ?", reqBody.RoleId).First(roleData)
	Model.DB.Where("companyId = ?", companySID).Where("roleId = ?", reqBody.RoleId).Find(rolePermission)

	for _, v := range *rolePermission {
		if rolePermissionMap[v.FuncCode] == nil {
			rolePermissionMap[v.FuncCode] = make(map[string][]int)
		}
		rolePermissionMap[v.FuncCode][v.ItemCode] = []int{}
	}

	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "成功",
			"data":  map[string]interface{}{
				"Role": *roleData,
				"Permission": rolePermissionMap,
			},
		},
	)
}

// 更新角色結構
func Update(Request *gin.Context) {
	TX := Model.DB.Begin()
	session := sessions.Default(Request)

	companySID, _ := strconv.Atoi(session.Get("companySID").(string))

	// 請求處理
	reqBody := new(struct {
		RoleId int
		RoleName string
		StopFlag string
		Type string
		/**
			Data = {
				[funcCode]: {
					[itemCode]: []RoleId
				}
			}
		*/
		Data map[string](map[string][]int)

	})

	if Request.ShouldBindJSON(&reqBody) != nil {
		ErrorHandle(Request, TX, "Request Data 格式不正確")
		return
	}

	// 要更新的欄位
	updateRoleQuery := map[string]interface{}{
		"roleName": reqBody.RoleName,
		"stopFlag": reqBody.StopFlag,
		"lastModify": time.Now(),
	}

	// 更新 或 新增 role table
	if reqBody.Type == "add" {
		var MaxCount int64
		TX.Model(&Model.Role{}).Where("companyId = ?", companySID).Count(&MaxCount)
		updateRoleQuery["companyId"] = companySID
		updateRoleQuery["roleId"] = MaxCount + 1

		TX.Model(&Model.Role{}).Create(&updateRoleQuery)
	} else {
		
		TX.Model(&Model.Role{}).Where(
			"companyId = ?",
			companySID,
		).Where(
			"roleId = ?",
			reqBody.RoleId,
		).Updates(&updateRoleQuery)
	}

	// 先把 此role structure 的資料 刪除
	TX.Where(
		"companyId = ?",
		companySID,
	).Where(
		"roleId = ?",
		reqBody.RoleId,
	).Delete(&Model.RoleStruct{})

	// 在 寫入 新的 進入 db
	for funcCode, itemObject := range reqBody.Data {
		for itemCode, _ := range itemObject {
			updateData := &Model.RoleStruct{
				CompanyId: companySID,
				RoleId: reqBody.RoleId,
				FuncCode: funcCode,
				ItemCode: itemCode,
				ScopeRole: "[]",
				CreateTime: time.Now(),
				LastModify: time.Now(),
			}
			if TX.Create(updateData).Error != nil {
				ErrorHandle(Request, TX, "新增失敗")
				return
			}
			
		}
	}

	TX.Commit()
	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}

func Delete(Request *gin.Context) {
	TX := Model.DB.Begin()
	session := sessions.Default(Request)

	TX.Delete(&)

	TX.Commit()
	Request.JSON(
		http.StatusOK,
		gin.H {
			"message": "更新成功",
		},
	)
}