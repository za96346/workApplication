package table

import (
	"time"
)

type WorkTimeExtend struct {
	WorkTime
	UserName string `json:"UserName"` // 名字
	Banch int64 `json:"Banch"` // 部門
	EmployeeNumber string `json:"EmployeeNumber"` // 員工編號
	BanchName string `json:"BanchName"` // 部門名
}

type UserExtend struct {
	UserTable
	BanchName string `json:"BanchName"` // 公司部們名稱
	CompanyName string `json:"CompanyName"` // 公司名稱
	CompanyId int64 `json:"CompanyId"` // 公司編號
}

type PerformanceExtend struct {
	Performance
	UserName string `json:"UserName"` // 名字
	CompanyId int64 `json:"CompanyId"` // 公司編號
}

//使用者
type UserTable struct {
	UserId int64 `json:"UserId"`// 使用者的編號
	CompanyCode string `json:"CompanyCode"` //公司碼
	Account string `json:"Account"`// 帳號
	Password string `json:"Password"`// 密碼
	UserName string `json:"UserName"` // 名字
	EmployeeNumber string `json:"EmployeeNumber"` // 員工編號
	OnWorkDay time.Time `json:"OnWorkDay"` // 到職日
	Banch int64 `json:"Banch"` // 部門
	Permession int `json:"Permession"` // 權限  (100 admin , 1 manager, 2 personal)
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
	BanchStyleId int64 `json:"BanchStyleId"` // 班表樣式id
	OnShiftTime time.Time  `json:"OnShiftTime"`// 開始上班時間
	OffShiftTime time.Time `json:"OffShiftTime"` //結束上班的時間
	RestTime string `json:"RestTime"` // 休息時間
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
	BossId int64 `json:"BossId"` //負責人 id
	SettlementDate int `json:"SettlementDate"` // 結算日
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

//部門班表的樣式
type BanchStyle struct {
	StyleId int64 `json:"StyleId"` // 此樣式的id
	BanchId int64 `json:"BanchId"` // 部門id
	Icon string `json:"Icon"` // 時段圖標
	RestTime string `json:"RestTime"` // 休息時間
	TimeRangeName string `json:"TimeRangeName"` // 時段名稱
	OnShiftTime string  `json:"OnShiftTime"`// 開始上班時間
	OffShiftTime string `json:"OffShiftTime"` //結束上班的時間
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

// 個部門的 每天上班規則
type BanchRule struct {
	RuleId int64 `json:"RuleId"` // 此規則的id
	BanchId int64 `json:"BanchId"` // 部門id
	MinPeople int `json:"MinPeople"` // 限制最少員工
	MaxPeople int `json:"MaxPeople"` // 限制做多的員工
	WeekDay int `json:"WeekDay"` // 星期幾 (1, 2, 3, 4, 5, 6, 7)
	WeekType int `json:"WeekType"` // 寒暑假 或 平常(1, 2, 3)
	OnShiftTime string  `json:"OnShiftTime"`// 開始上班時間
	OffShiftTime string `json:"OffShiftTime"` //結束上班的時間
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

// 離職員工表
type QuitWorkUser struct {
	QuitId int64 `json:"QuitId"` //離職的唯一id
	UserId int64 `json:"UserId"`// 使用者的編號
	CompanyCode string `json:"CompanyCode"` //公司碼
	Account string `json:"Account"`// 帳號
	UserName string `json:"UserName"` // 名字
	EmployeeNumber string `json:"EmployeeNumber"` // 員工編號
	OnWorkDay time.Time `json:"OnWorkDay"` // 到職日
	Banch int64 `json:"Banch"` // 部門
	Permession int `json:"Permession"` // 權限  (100 admin , 1 manager, 2 personal)
	MonthSalary int `json:"MonthSalary"` // 月薪
	PartTimeSalary int `json:"PartTimeSalary"` // 時薪
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

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

// 公司假日的設定
type WeekendSetting struct {
	WeekendId int64 `json:"WeekendId"` // 假日的 唯一id
	CompanyId int64 `json:"CompanyId"` // 公司的唯一編號
	Date string `json:"Date"` // 每天 ex: 2022-02-22
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

// 個人時數
type WorkTime struct {
	WorkTimeId int64 `json:"WorkTimeId"` // 時數id
	UserId int64 `json:"UserId"` // 使用者id
	Year int `json:"Year"` // 年份
	Month int `json:"Month"` // 月份
	WorkHours int `json:"WorkHours"` // 應當工作時數
	TimeOff int `json:"TimeOff"` // 應當休假天數
	UsePaidVocation int `json:"UsePaidVocation"` // 使用特休的次數
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

// 個人特休
type PaidVocation struct {
	PaidVocationId int64 `json:"PaidVocationId"` // 特休id
	UserId int64 `json:"UserId"` // 使用者id
	Year int `json:"Year"` // 年份
	Count int `json:"Count"` // 次數
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

// log
type Log struct {
	LogId int64
	Msg string
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

type Performance struct {
	PerformanceId int64 `json:"PerformanceId"`
	UserId int64 `json:"UserId"`
	Year int `json:"Year"`
	Month int `json:"Month"`
	BanchId int64 `json:"BanchId"`
	Goal string `json:"Goal"`
	Attitude int `json:"Attitude"`
	Efficiency int `json:"Efficiency"`
	Professional int `json:"Professional"`
	Directions string `json:"Directions"`
	BeLate int `json:"BeLate"`
	DayOffNotOnRule int `json:"DayOffNotOnRule"`
	BanchName string `json:"BanchName"`
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

