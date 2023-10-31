package table

import (
	"time"

)

//換班
type ShiftChangeTable struct {
	InitiatorShiftId int64 `json:"InitiatorShiftId"` // 發起人班表的編號
	RequestedShiftId int64 `json:"RequestedShiftId"` // 被請求人的班表編號
	CaseId int64 `json:"CaseId"`
	Reason string `json:"Reason"` // 申請理由
	CaseProcess string `json:"CaseProcess"` // 此案件的進度狀態到哪
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
	SpecifyTag string `json:"SpecifyTag"` // 特別的備註
}