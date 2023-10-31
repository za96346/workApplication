package table

import (
	"time"

)

// 個人時數
type WorkTime struct {
	WorkTimeId int64 `json:"WorkTimeId"` // 時數id
	UserId int64 `json:"UserId"` // 使用者id
	Year int `json:"Year"` // 年份
	Month int `json:"Month"` // 月份
	WorkHours int `json:"WorkHours"` // 應當工作時數
	TimeOff int `json:"TimeOff"` // 應當休假天數
	UsePaidVocation int `json:"UsePaidVocation"` // 使用特休的次數
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

type WorkTimeExtend struct {
	WorkTime
	UserName string `json:"UserName"` // 名字
	Banch int64 `json:"Banch"` // 部門
	EmployeeNumber string `json:"EmployeeNumber"` // 員工編號
	BanchName string `json:"BanchName"` // 部門名
}
