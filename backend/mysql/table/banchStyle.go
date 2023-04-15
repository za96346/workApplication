package table

import (
	"time"

)

//部門班表的樣式
type BanchStyle struct {
	StyleId int64 `json:"StyleId"` // 此樣式的id
	BanchId int64 `json:"BanchId"` // 部門id
	Icon string `json:"Icon"` // 時段圖標
	RestTime string `json:"RestTime"` // 休息時間
	TimeRangeName string `json:"TimeRangeName"` // 時段名稱
	OnShiftTime string  `json:"OnShiftTime"`// 開始上班時間
	OffShiftTime string `json:"OffShiftTime"` //結束上班的時間
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}