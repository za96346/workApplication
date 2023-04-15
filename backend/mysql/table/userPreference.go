package table

import (
	"time"
)


//使用者偏好
type UserPreferenceTable struct {
	UserId int64 `json:"UserId"`
	Style string `json:"Style"`
	FontSize string `json:"FontSize"`
	SelfPhoto string `json:"SelfPhoto"`
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}