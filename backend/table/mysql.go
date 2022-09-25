package table

import (
	"time"
)


//共同的欄位
type commonColumn struct {
	CaseId int64
	Reason string // 申請理由
	CaseProcess string // 此案件的進度狀態到哪
	timeColumn
	specifyTagColumn
}

type timeColumn struct {
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}
type specifyTagColumn struct {
	SpecifyTag string `json:"SpecifyTag"` // 特別的備註
}

//使用者
type UserTable struct {
	UserId int64 // 使用者的編號
	CompanyCode string //公司碼
	Account string // 帳號
	Password string // 密碼
	OnWorkDay time.Time // 到職日
	Banch string // 部門
	Permession string // 權限
	WorkState string // 工作狀態 (到職on or 離職off)
	MonthSalary int64 // 月薪
	PartTimeSalary int64 // 時薪
	timeColumn
}

//使用者偏好
type UserPreferenceTable struct {
	UserId int64
	Style string
	FontSize string
	SelfPhoto string
	timeColumn
}

//班表
type ShiftTable struct {
	ShiftId int64 `json:"ShiftId"` // 班表的編號
	UserId int64 `json:"UserId"` // 使用者的編號
	OnShiftTime time.Time  `json:"OnShiftTime"`// 開始上班時間
	OffShiftTime time.Time `json:"OffShiftTime"` //結束上班的時間
	PunchIn time.Time `json:"PunchIn"` // 上班卡
	PunchOut time.Time `json:"PunchOut"`// 下班卡
	timeColumn
	specifyTagColumn
}

//遲到早退
type LateExcusedTable struct {
	ShiftId int64 // 班表的編號
	LateExcusedType  string // 遲到 或是 早退
	commonColumn
}

//加班
type ShiftOverTimeTable struct {
	ShiftId int64 // 班表的編號
	InitiatorOnOverTime time.Time // 申請人 開始加班時間
	InitiatorOffOverTime time.Time // 申請人 結束加班時間
	commonColumn
}

//換班
type ShiftChangeTable struct {
	InitiatorShiftId int64 // 發起人班表的編號
	RequestedShiftId int64 // 被請求人的班表編號
	commonColumn
}

//請假
type DayOffTable struct {
	ShiftId int64 // 班表的編號
	DayOffType string // 請假類型
	commonColumn
}

//忘記打卡
type ForgetPunchTable struct {
	ShiftId int64 // 班表的編號
	TargetPunch string // 上班卡 或是 下班卡
	commonColumn
}

//公司
type CompanyTable struct {
	CompanyId int64 // 公司編號
	CompanyCode string // 公司碼
	CompanyName string // 公司名稱
	CompanyLocation string // 公司位置
	CompanyPhoneNumber string // 公司電話
	TermStart time.Time // 開始期間
	TermEnd time.Time // 結束期間
	timeColumn
}

//公司部們
type CompanyBanchTable struct {
	Id int64
	CompanyId int64 // 公司編號
	BanchName string // 公司部們名稱
	BanchShiftStyle string // 部門班表樣式
	timeColumn
}






