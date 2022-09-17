package database

import (
	"fmt"
	"sync"
	"time"
)
var userTableMux sync.RWMutex
var UserTableInstance *UserTable
var queryMux *sync.Mutex
var SqlQuery *sqlQuery


//共同的欄位
type commonColumn struct {
	Reason string // 申請理由
	CaseProcess string // 此案件的進度狀態到哪
	timeColumn
	specifyTagColumn
}

type timeColumn struct {
	CreateTime time.Time //創建的時間
	LastModify time.Time // 上次修改的時間
}
type specifyTagColumn struct {
	SpecifyTag string // 特別的備註
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
	ShiftId int64 // 班表的編號
	UserId int64 // 使用者的編號
	OnShiftTime string // 開始上班時間
	OffShiftTime string //結束上班的時間
	PunchIn time.Time // 上班卡
	PunchOut time.Time // 下班卡
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
	CompanyId int64 // 公司編號
	BanchName string // 公司部們名稱
	BanchShiftStyle string // 部門班表樣式
	timeColumn
}







// sql query 
type sqlQuery struct {
	User userQuery
	UserPreference userPreferenceQuery
	Company companyQuery
	CompanyBanch companyBanchQuery
	Shift shiftQuery
	ShiftChange shiftChangeQuery
	ShiftOverTime shiftOverTimeQuery
	ForgetPunch forgetPunchQuery
	LateExcused lateExcusedQuery
	DayOff dayOffQuery
}
type queryCommonColumn struct {
	InsertAll string
	SelectAll string
}

type userQuery struct {
	queryCommonColumn
	InsertSingle string
	SelectSingle string
}
type userPreferenceQuery struct {
	queryCommonColumn
	InsertSingle string
}
type companyQuery struct {
	queryCommonColumn
	InsertSingle string
	SelectSingle string
}
type companyBanchQuery struct {
	queryCommonColumn
}
type shiftQuery struct {
	queryCommonColumn
	SelectSingleByUserId string
	SelectSingleByShiftId string
}
type shiftChangeQuery struct {
	queryCommonColumn
}
type shiftOverTimeQuery struct {
	queryCommonColumn
}
type forgetPunchQuery struct {
	queryCommonColumn
}
type lateExcusedQuery struct {
	queryCommonColumn
}
type  dayOffQuery struct {
	queryCommonColumn
}


// sql query end
func UserTableSingleton() *UserTable {
	if UserTableInstance == nil {
		userTableMux.Lock()
		defer userTableMux.Unlock()
		if UserTableInstance == nil {
			UserTableInstance = &UserTable{}
			return UserTableInstance
		}
	}
	return UserTableInstance
}

func SqlQuerySingleton() *sqlQuery {
	queryMux = new(sync.Mutex)
	if SqlQuery == nil {
		queryMux.Lock()
		if SqlQuery == nil {
			SqlQuery = &sqlQuery{}
			addUserQuery()
			addCompanyQuery()
			addUserPreferenceQuery()
			addCompanyBanchQuery()
			addShiftQuery()
			addShiftChangeQuery()
			addShiftOverTimeQuery()
			addForgetPunchQuery()
			addLateExcusedQuery()
			addDayOffQuery()
			defer queryMux.Unlock()
			return SqlQuery
		}
	}
	return SqlQuery
}

func addUserQuery() {
	fmt.Println(SqlQuery)
	SqlQuery.User.InsertAll = `insert into user(
		companyCode,
		account,
		password,
		onWorkDay,
		banch,
		permession,
		workState,
		createTime,
		lastModify,
		monthSalary,
		partTimeSalary
		) values(
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.User.SelectAll = `select * from user;`;
	SqlQuery.User.SelectSingle = `select * from user where account=? or userId=?;`;
}
func addUserPreferenceQuery() {
	SqlQuery.UserPreference.InsertAll = `insert into userPreference(
		userId,
		style,
		fontSize,
		selfPhoto,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.UserPreference.SelectAll = `select * from userPreference;`;
}
func addCompanyQuery() {
	SqlQuery.Company.InsertAll = `insert into company(
		companyCode,
		companyName,
		companyLocation,
		companyPhoneNumber,
		termStart,
		termEnd,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.Company.SelectSingle = `select * from company where companyId=? or companyCode=?;`;
	SqlQuery.Company.SelectAll = `select * from company;`;
}
func addCompanyBanchQuery() {
	SqlQuery.CompanyBanch.InsertAll = `insert into companyBanch(
		companyId,
		banchName,
		banchShiftStyle,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?
	);`;
	SqlQuery.CompanyBanch.SelectAll = `select * from companyBanch`;
}
func addShiftQuery() {
	SqlQuery.Shift.InsertAll = `insert into shift(
	userId,
	onShiftTime,
	offShiftTime,
	punchIn,
	punchOut,
	specifyTag,
	createTime,
	lastModify
	) values(
		?, ?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.Shift.SelectSingleByUserId = `select * from shift where userId=?;`;
	SqlQuery.Shift.SelectSingleByShiftId = `select * from shift where shiftId=?;`;
	SqlQuery.Shift.SelectAll = `select * from shift;`;
}
func addShiftChangeQuery() {
	SqlQuery.ShiftChange.InsertAll = `insert into shiftChange(
		initiatorShiftId,
		requestedShiftId,
		reson,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
	) values(
		?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.ShiftChange.SelectAll = `select * from shiftChange;`;
}
func  addShiftOverTimeQuery() {
	SqlQuery.ShiftOverTime.InsertAll = `insert into shiftOverTime(
		shiftId,
		initiatorOnOverTime,
		initiatorOffOverTime,
		reson,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
	) values(
		?, ?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.ShiftOverTime.SelectAll = `select * from shiftOverTime;`;
}
func addForgetPunchQuery() {
	SqlQuery.ShiftOverTime.InsertAll = `insert into forgetPunch(
		shiftId,
		targetPunch,
		reson,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
	) values(
		?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.ForgetPunch.SelectAll = `select * from forgetPunch;`;
}
func addLateExcusedQuery() {
	SqlQuery.LateExcused.InsertAll = `insert into lateExcused(
		shiftId,
		lateExcusedType,
		reson,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
	) values(
		?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.LateExcused.SelectAll = `select * from lateExcused;`;
}
func addDayOffQuery() {
	SqlQuery.DayOff.InsertAll = `insert into dayOff(
		shiftId,
		dayOffType,
		reson,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
	) values(
		?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.DayOff.SelectAll = `select * from dayOff;`;
}