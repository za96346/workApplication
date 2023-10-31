package table

import (
	"time"

)

//公司
type CompanyTable struct {
	CompanyId int64 `json:"CompanyId"` // 公司編號
	CompanyCode string `json:"CompanyCode"` // 公司碼
	CompanyName string `json:"CompanyName"` // 公司名稱
	CompanyLocation string `json:"CompanyLocation"` // 公司位置
	CompanyPhoneNumber string `json:"CompanyPhoneNumber"` // 公司電話
	BossId int64 `json:"BossId"` //負責人 id
	SettlementDate int `json:"SettlementDate"` // 結算日
	TermStart time.Time `json:"TermStart"` // 開始期間
	TermEnd time.Time `json:"TermEnd"` // 結束期間
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}