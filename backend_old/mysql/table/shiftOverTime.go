package table

import (
	"time"
)

//加班
type ShiftOverTimeTable struct {
	ShiftId int64 `json:"ShiftId"` // 班表的編號
	InitiatorOnOverTime time.Time `json:"InitiatorOnOverTime"` // 申請人 開始加班時間
	InitiatorOffOverTime time.Time `json:"InitiatorOffOverTime"` // 申請人 結束加班時間
	CaseId int64 `json:"CaseId"`
	Reason string `json:"Reason"` // 申請理由
	CaseProcess string `json:"CaseProcess"` // 此案件的進度狀態到哪
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
	SpecifyTag string `json:"SpecifyTag"` // 特別的備註
}