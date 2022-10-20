package response

import "time"

// "backend/database"

type Response struct {
	Message string `binding:"required"`
	Data string	`binding:"required"`
	Status bool `binding:"required"`
}

type User struct {
	UserId int64 `json:"UserId"`// 使用者的編號
	CompanyCode string `json:"CompanyCode"` //公司碼
	EmployeeNumber string `json:"EmployeeNumber"` // 員工編號
	OnWorkDay time.Time `json:"OnWorkDay"` // 到職日
	UserName string `json:"UserName"`  // 姓名
	Banch int64 `json:"Banch"` // 部門
	Permession int `json:"Permession"` // 權限
	WorkState string `json:"WorkState"` // 工作狀態 (到職on or 離職off)
}

func NewResponse(message string, data string, status bool) *Response {
	res := new(Response)
	res.Message = message
	res.Data = data
	res.Status = status
	return res
}
type OnlyEmail struct {
	Email string `json: "Email"`
}
type Register struct {
	Account string `json:"Account"`
	Captcha	int `json:"Captcha"`
	Password string `json:"Password"`
	PasswordConfirm string `json:"PasswordConfirm"`
	CompanyCode string `json:"CompanyCode"`
	UserName string `json:"UserName"`
}

type ForgetPassword struct {
	Captcha int `json:"Captcha"`
	OldPassword string `json:"OldPassword"`
	NewPassword string `json:"NewPassword"`
	NewPasswordConfirm string `json:"NewPasswordConfirm"`
	Email string `json:"Email"`
}