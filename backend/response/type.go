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
	OnWorkDay time.Time `json:"OnWorkDay"` // 到職日
	Banch string `json:"Banch"` // 部門
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