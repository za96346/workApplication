package method

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/domain/entities"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
	請求body, param
*/
type ReqStruct struct {
	ReqBodyValidation bool // 是否開啟 請求 json binding 驗證
	ReqBodyStruct interface{} // 請求結構 ( please give it as a pointer. )
	ReqParamsValidation bool // 是否開啟 請求 json binding 驗證
	ReqParamsStruct interface{} // 請求結構 ( please give it as a pointer. )
}

type SessionStruct struct {
	Request *gin.Context
	IsLogin bool // 是否成功登入 "Y" | "N"
	User *entities.User
	Permission interface{}
}

/*
	session 基本處理
*/
func NewSession(
	Request *gin.Context,
	req *ReqStruct,
) (*SessionStruct, error){
	instance := new(SessionStruct)
	session := sessions.Default(Request)

	// 公司 id
	companyId, err := strconv.Atoi(session.Get("companyId").(string))
	if err != nil {
		Request.JSON(
			http.StatusOK,
			gin.H {
				"message": "公司id Error",
			},
		)
		return nil, err
	}

	// 是否登入
	isLogin := false
	if session.Get("isLogin") == "Y" {isLogin = true}

	// 權限
	permission := session.Get("permission")

	// 使用者id
	userId, err := strconv.Atoi(session.Get("userId").(string))
	if err != nil {
		Request.JSON(
			http.StatusOK,
			gin.H {
				"message": "使用者id Error",
			},
		)
		return nil, err
	}

	//角色id
	roleId, err := strconv.Atoi(session.Get("roleId").(string))
	if err != nil {
		Request.JSON(
			http.StatusOK,
			gin.H {
				"message": "使用者角色id Error",
			},
		)
		return nil, err
	}

	// 部門id
	banchId, err := strconv.Atoi(session.Get("banchId").(string))
	if err != nil {
		Request.JSON(
			http.StatusOK,
			gin.H {
				"message": "使用者部門id Error",
			},
		)
		return nil, err
	}

	// 使用者姓名
	userName := session.Get("userName").(string)

	// 使用者 員工編號
	employeeNumber := session.Get("employeeNumber").(string)

	if req != nil {
		// 請求資料驗證 body
		if (*req).ReqBodyValidation {
			bindError := Request.ShouldBindJSON((*req).ReqBodyStruct)
		
			if bindError != nil {
				Request.JSON(
					http.StatusOK,
					gin.H {
						"message": fmt.Sprintf("Request Data 格式不正確 %s", bindError),
					},
				)
				return nil, bindError
			}	
		}

		// 請求資料驗證 params
		if (*req).ReqParamsValidation {
			bindError := Request.ShouldBindQuery((*req).ReqParamsStruct)
		
			if bindError != nil {
				Request.JSON(
					http.StatusOK,
					gin.H {
						"message": fmt.Sprintf("Request Params 格式不正確 %s", bindError),
					},
				)
				return nil, bindError
			}	
		}
	}

	// 綁定物件
	(*instance).IsLogin = isLogin
	(*instance).User = &entities.User{
		CompanyId: companyId,
		UserId: userId,
		RoleId: roleId,
		BanchId: &banchId,
		UserName: userName,
		EmployeeNumber: employeeNumber,
	}
	(*instance).Permission = permission

	return instance, nil
}
