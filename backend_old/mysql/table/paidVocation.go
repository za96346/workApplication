package table

import (
	"time"

)

// 個人特休
type PaidVocation struct {
	PaidVocationId int64 `json:"PaidVocationId"` // 特休id
	UserId int64 `json:"UserId"` // 使用者id
	Year int `json:"Year"` // 年份
	Count int `json:"Count"` // 次數
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}