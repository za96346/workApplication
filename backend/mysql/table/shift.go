package table

import (
	"database/sql"
	"time"

)


//班表
type ShiftTable struct {
	ShiftId int64 `json:"ShiftId"` // 班表的編號
	UserId int64 `json:"UserId"` // 使用者的編號
	BanchStyleId int64 `json:"BanchStyleId"` // 班表樣式id
	BanchId int64 `json:"BanchId"` // 部門id
	Year int `json:"Year"` // 紀錄 年
	Month int `json:"Month"` // 紀錄 月
	Icon string `json:"Icon"` // 圖標
	OnShiftTime time.Time  `json:"OnShiftTime"`// 開始上班時間
	OffShiftTime time.Time `json:"OffShiftTime"` //結束上班的時間
	RestTime string `json:"RestTime"` // 休息時間
	PunchIn sql.NullTime `json:"PunchIn"` // 上班卡
	PunchOut sql.NullTime `json:"PunchOut"`// 下班卡
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
	SpecifyTag string `json:"SpecifyTag"` // 特別的備註
}

type ShiftExtend struct {
	ShiftTable
	UserName string // 使用者名稱
	Permission int // 權限
	Banch int64 // 部門編號
	EmployeeNumber string // 員工編號
}

// 班表統計
type ShiftTotal struct {
	UserId int64 `json:"UserId"`
	Year int `json:"Year"`
	Month int `json:"Month"`
	BanchId int64 `json:"BanchId"`
	UserName string `json:"UserName"`
	Permession int `json:"Permession"`
	EmployeeNumber string `json:"EmployeeNumber"`
	ChangeCocunt int `json:"ChangeCocunt"`
	OverTimeCount int `json:"OverTimeCount"`
	ForgetPunchCount int `json:"ForgetPunchCount"`
	DayOffCount int `json:"DayOffCount"`
	LateExcusedCount int `json:"LateExcusedCount"`
	TotalHours float32 `json:"TotalHours"`
}