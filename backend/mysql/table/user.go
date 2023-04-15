package table

import (
	"time"

)

type UserExtend struct {
	UserTable
	BanchName string `json:"BanchName"` // 公司部們名稱
	CompanyName string `json:"CompanyName"` // 公司名稱
	CompanyId int64 `json:"CompanyId"` // 公司編號
}

//使用者
type UserTable struct {
	UserId int64 `json:"UserId"`// 使用者的編號
	CompanyCode string `json:"CompanyCode"` //公司碼
	Account string `json:"Account"`// 帳號
	Password string `json:"Password"`// 密碼
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