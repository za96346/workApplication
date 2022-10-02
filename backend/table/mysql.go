package table

import (
	"time"
)

//使用者
type UserTable struct {
	UserId int64 `json:"UserId"`// 使用者的編號
	CompanyCode string `json:"CompanyCode"` //公司碼
	Account string `json:"Account"`// 帳號
	Password string `json:"Password"`// 密碼
	OnWorkDay time.Time `json:"OnWorkDay"` // 到職日
	Banch string `json:"Banch"` // 部門
	Permession string `json:"Permession"` // 權限
	WorkState string `json:"WorkState"` // 工作狀態 (到職on or 離職off)
	MonthSalary int `json:"MonthSalary"` // 月薪
	PartTimeSalary int `json:"PartTimeSalary"` // 時薪
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

//使用者偏好
type UserPreferenceTable struct {
	UserId int64 `json:"UserId"`
	Style string `json:"Style"`
	FontSize string `json:"FontSize"`
	SelfPhoto string `json:"SelfPhoto"`
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

//班表
type ShiftTable struct {
	ShiftId int64 `json:"ShiftId"` // 班表的編號
	UserId int64 `json:"UserId"` // 使用者的編號
	OnShiftTime time.Time  `json:"OnShiftTime"`// 開始上班時間
	OffShiftTime time.Time `json:"OffShiftTime"` //結束上班的時間
	PunchIn time.Time `json:"PunchIn"` // 上班卡
	PunchOut time.Time `json:"PunchOut"`// 下班卡
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
	SpecifyTag string `json:"SpecifyTag"` // 特別的備註
}

//遲到早退
type LateExcusedTable struct {
	ShiftId int64 `json:"ShiftId"` // 班表的編號
	LateExcusedType  string `json:"LateExcusedType"` // 遲到 或是 早退
	CaseId int64 `json:"CaseId"`
	Reason string `json:"Reason"` // 申請理由
	CaseProcess string `json:"CaseProcess"` // 此案件的進度狀態到哪
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
	SpecifyTag string `json:"SpecifyTag"` // 特別的備註
}

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

//請假
type DayOffTable struct {
	ShiftId int64 `json:"ShiftId"` // 班表的編號
	DayOffType string `json:"DayOffType"` // 請假類型
	CaseId int64 `json:"CaseId"`
	Reason string `json:"Reason"` // 申請理由
	CaseProcess string `json:"CaseProcess"` // 此案件的進度狀態到哪
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
	SpecifyTag string `json:"SpecifyTag"` // 特別的備註
}

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

//公司
type CompanyTable struct {
	CompanyId int64 `json:"CompanyId"` // 公司編號
	CompanyCode string `json:"CompanyCode"` // 公司碼
	CompanyName string `json:"CompanyName"` // 公司名稱
	CompanyLocation string `json:"CompanyLocation"` // 公司位置
	CompanyPhoneNumber string `json:"CompanyPhoneNumber"` // 公司電話
	TermStart time.Time `json:"TermStart"` // 開始期間
	TermEnd time.Time `json:"TermEnd"` // 結束期間
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

//公司部們
type CompanyBanchTable struct {
	Id int64 `json:"Id"`
	CompanyId int64 `json:"CompanyId"` // 公司編號
	BanchName string `json:"BanchName"` // 公司部們名稱
	BanchShiftStyle string `json:"BanchShiftStyle"` // 部門班表樣式
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}
