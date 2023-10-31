package table

import (
	"time"

)

// 等待配對公司
type WaitCompanyReply struct {
	WaitId int64 `json:"WaitId"` // 等待的唯一id
	UserId int64 `json:"UserId"` // 使用者的編號
	UserName   string    `json:"UserName"`   // 名字 // 這個不會出現在 mysql table
	CompanyId int64 `json:"CompanyId"` // 公司碼
	SpecifyTag string `json:"SpecifyTag"` // 備註
	IsAccept int `json:"IsAccept"` // 是否接受 (1 等待確認, 2 接受, 3 拒絕)
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}