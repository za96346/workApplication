package table

import (
	"time"

)

// 離職員工表
type QuitWorkUser struct {
	QuitId int64 `json:"QuitId"` //離職的唯一id
	UserId int64 `json:"UserId"`// 使用者的編號
	CompanyCode string `json:"CompanyCode"` //公司碼
	Account string `json:"Account"`// 帳號
	UserName string `json:"UserName"` // 名字
	EmployeeNumber string `json:"EmployeeNumber"` // 員工編號
	OnWorkDay time.Time `json:"OnWorkDay"` // 到職日
	Banch int64 `json:"Banch"` // 部門
	Permession int `json:"Permession"` // 權限  (100 admin , 1 manager, 2 personal)
	MonthSalary int `json:"MonthSalary"` // 月薪
	PartTimeSalary int `json:"PartTimeSalary"` // 時薪
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}