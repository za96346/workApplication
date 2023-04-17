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
	BanchName string `json:"BanchName"` // 公司部們名稱
	CompanyName string `json:"CompanyName"` // 公司名稱
	CompanyId int64 `json:"CompanyId"` // 公司編號
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

type ChangePassword struct {
	Captcha int `json:"Captcha"`
	OldPassword string `json:"OldPassword"`
	NewPassword string `json:"NewPassword"`
	NewPasswordConfirm string `json:"NewPasswordConfirm"`
}

type ForgetPassword struct {
	ChangePassword
	Email string `json:"Email"`
}

type Member struct {
	UserName string // 用戶名
	UserId int64 // 使用者編號
	BanchId int64 // 自己 部門id
	Permission int // 權限
	Pic string // 用戶的照片
	Color string // 用戶的編輯顏色
	Online int // 上線狀態
	Position int // 當前編輯的位置
}

type Shift struct {
	UserId int64 `json:"UserId"` // 使用者的編號
	// Position int `json:"Position"` // 位置
	Icon string `json:"Icon"`
	Date string `json:"Date"` // 每天 ex: 2022-02-22
	BanchStyleId int64
	RestTime string `json:"RestTime"` // 休息時間 ex: 01:00:00
	OnShiftTime string  `json:"OnShiftTime"`// 開始上班時間 ex: 09:00:00
	OffShiftTime string`json:"OffShiftTime"` //結束上班的時間 ex: 18:00:00
}
