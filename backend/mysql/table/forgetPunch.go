package table

import (
	"time"

)

//忘記打卡
type ForgetPunchTable struct {
	ShiftId int64 `json:"ShiftId"` // 班表的編號
	TargetPunch string `json:"TargetPunch"` // 上班卡 或是 下班卡
	CaseId int64 `json:"CaseId"`
	Reason string `json:"Reason"` // 申請理由
	CaseProcess string `json:"CaseProcess"` // 此案件的進度狀態到哪
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
	SpecifyTag string `json:"SpecifyTag"` // 特別的備註
}