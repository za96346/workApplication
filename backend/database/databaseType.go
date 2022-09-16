package database

import (
	"fmt"
	"sync"
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
	CreateTime string //創建的時間
	LastModify string // 上次修改的時間
}
type specifyTagColumn struct {
	SpecifyTag string // 特別的備註
}

//使用者
type UserTable struct {
	UserId float64 // 使用者的編號
	CompanyCode string //公司碼
	Account string // 帳號
	Password string // 密碼
	OnWorkDay string // 到職日
	Banch string // 部門
	Permession string // 權限
	WorkState string // 工作狀態 (到職on or 離職off)
	MonthSalary float64 // 月薪
	PartTimeSalary float64 // 時薪
	timeColumn
}

//班表
type ShiftTable struct {
	ShiftId float64 // 班表的編號
	UserId float64 // 使用者的編號
	OnShiftTime string // 開始上班時間
	OffShiftTime string //結束上班的時間
	PunchIn string // 上班卡
	PunchOut string // 下班卡
	timeColumn
	specifyTagColumn
}

//遲到早退
type LateExcusedTable struct {
	ShiftId float64 // 班表的編號
	LateExcusedType  string // 遲到 或是 早退
	commonColumn
}

//加班
type ShiftOverTimeTable struct {
	ShiftId float64 // 班表的編號
	InitiatorOnOverTime string // 申請人 開始加班時間
	InitiatorOffOverTime string // 申請人 結束加班時間
	commonColumn
}

//換班
type ShiftChangeTable struct {
	InitiatorShiftId float64 // 發起人班表的編號
	RequestedShiftId float64 // 被請求人的班表編號
	commonColumn
}

//請假
type DayOffTable struct {
	ShiftId float64 // 班表的編號
	DayOffType string // 請假類型
	commonColumn
}

//忘記打卡
type ForgetPunchTable struct {
	ShiftId float64 // 班表的編號
	TargetPunch string // 上班卡 或是 下班卡
	commonColumn
}

//公司
type CompanyTable struct {
	CompanyId float64 // 公司編號
	CompanyCode string // 公司碼
	CompanyName string // 公司名稱
	CompanyLocation string // 公司位置
	CompanyPhoneNumber string // 公司電話
	TermStart string // 開始期間
	TermEnd string // 結束期間
	timeColumn
}

//公司部們
type CompanyBanchTable struct {
	CompanyId float64 // 公司編號
	BanchName string // 公司部們名稱
	BanchShiftStyle string // 部門班表樣式
	timeColumn
}







// sql query 
type sqlQuery struct {
	User userQuery
	UserPreference userPerferenceQuery
	Company companyQuery
	CompanyBanch companyBanchQuery
}

type userQuery struct {
	InsertAll string
	InsertSingle string
	selectAll string
	selectSingle string
}
type userPerferenceQuery struct {
	InsertAll string
	InsertSingle string
}
type companyQuery struct {
	InsertAll string
	InsertSingle string
	SelectSingle string
}
type companyBanchQuery struct {
	InsertAll string
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
	SqlQuery.User.selectAll = `select * from user;`;
	SqlQuery.User.selectSingle = `select * from user where account=? or userId=?;`;
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
}