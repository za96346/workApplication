package table

import (
	"time"

)

// 個部門的 每天上班規則
type BanchRule struct {
	RuleId int64 `json:"RuleId"` // 此規則的id
	BanchId int64 `json:"BanchId"` // 部門id
	MinPeople int `json:"MinPeople"` // 限制最少員工
	MaxPeople int `json:"MaxPeople"` // 限制做多的員工
	WeekDay int `json:"WeekDay"` // 星期幾 (1, 2, 3, 4, 5, 6, 7)
	WeekType int `json:"WeekType"` // 寒暑假 或 平常(1, 2, 3)
	OnShiftTime string  `json:"OnShiftTime"`// 開始上班時間
	OffShiftTime string `json:"OffShiftTime"` //結束上班的時間
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}