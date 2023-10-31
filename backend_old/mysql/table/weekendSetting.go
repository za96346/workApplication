package table

import (
	"time"

)

// 公司假日的設定
type WeekendSetting struct {
	WeekendId int64 `json:"WeekendId"` // 假日的 唯一id
	CompanyId int64 `json:"CompanyId"` // 公司的唯一編號
	Date string `json:"Date"` // 每天 ex: 2022-02-22
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}