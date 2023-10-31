package table

import (
	"time"
)


//班表
type ShiftEditLog struct {
	LogId int64 `json:"LogId"`
	Year int `json:"Year"`
	Month int `json:"Month"`
	BanchId int64 `json:"BanchId"`
	Msg string `json:"Msg"`
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}
