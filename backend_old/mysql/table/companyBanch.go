package table

import (
	"time"

)

//公司部們
type CompanyBanchTable struct {
	Id int64 `json:"Id"`
	CompanyId int64 `json:"CompanyId"` // 公司編號
	BanchName string `json:"BanchName"` // 公司部們名稱
	BanchShiftStyle string `json:"BanchShiftStyle"` // 部門班表樣式
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
	UserTotal int // 總人數
}