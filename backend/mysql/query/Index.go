package query

import (
	"sync"
)

var queryMux *sync.Mutex
var sqlQueryInstance *sqlQuery

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
	BanchStyle banchStyle
	BanchRule banchRule
	QuitWorkUser quitWorkUser
	WaitCompanyReply waitCompanyReply
	WeekendSetting weekendSetting
	WorkTime workTime
	PaidVocation paidVocation
	Log log
	Performance performance
}
type queryCommonColumn struct {
	InsertAll string
	SelectAll string
	Delete string
	UpdateSingle string
}

type userQuery struct {
	queryCommonColumn
	InsertSingle string
	SelectSingleByUserId string
	SelectSingleByAccount string
	SelectAllByCompanyCode string
	SelectAllByBanchId string
	SelectAllByUserIdAndCompanyCode string
	SelectAllByAdmin string
	SelectAllByManager string
	UpdateCompanyUser string
	UpdateBoss string
}
type userPreferenceQuery struct {
	queryCommonColumn
	InsertSingle string
	SelectSingleByUserId string
}
type companyQuery struct {
	queryCommonColumn
	InsertSingle string
	SelectSingleByCompanyId string
	SelectSingleByCompanyCode string
}
type companyBanchQuery struct {
	queryCommonColumn
	SelectSingleByCompanyId string
	SelectSingleById string
	UpdateByCompanyCode string
	DeleteByCompanyCode string
	SelectByCompanyCodeAndBanchID string
}
type shiftQuery struct {
	queryCommonColumn
	SelectSingleByUserId string
	SelectSingleByShiftId string
	SelectTotal string
}
type shiftChangeQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
	SelectAllByInitiatorShiftId string
	SelectAllByRequestedShiftId string
}
type shiftOverTimeQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
	SelectAllByShiftId string
}
type forgetPunchQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
	SelectAllByShiftId string
}
type lateExcusedQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
	SelectAllByShiftId string
}
type  dayOffQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
	SelectAllByShiftId string
}
type banchStyle struct {
	queryCommonColumn
	SelectSingleByStyleId string
	SelectAllByBanchId string
	SelectByCompanyCode string
	UpdateByCompanyCode string
	DeleteByCompanyCode string
}
type banchRule struct {
	queryCommonColumn
	SelectSingleByRuleId string
	SelectAllByBanchId string
}
type quitWorkUser struct {
	queryCommonColumn
	SelectSingleByUserId string
	SelectAllByCompanyCode string
	SelectSingleByQuitId string
	SelectSingleByCompanyCodeAndUserId string
	InsertBySelectUser string
	DeleteByJoinUser string
}
type waitCompanyReply struct {
	queryCommonColumn
	SelectSingleByWaitId string
	SelectAllByUserId string
	SelectAllByCompanyId string
	SelectAllByCompanyIdAndUserId string
	SelectAllJoinUserTable string
}
type weekendSetting struct {
	queryCommonColumn
	SelectSingleByWeekendId string
	SelectAllByCompanyId string
}
type workTime struct {
	queryCommonColumn
	SelectAllByUserId string
	SelectAllByTime string
	SelectAllByPrimaryKey string
	DeleteByCompanyAndId string
}
type paidVocation struct {
	queryCommonColumn
	SelectAllByUserId string
	SelectAllByTime string
}
type log struct {
	queryCommonColumn
}
type performance struct {
	queryCommonColumn
	SelectAllByAdmin string
	SelectAllByManager string
	SelectAllByPerson string
	SelectSingleByAdmin string
	SelectSingleByManager string
	SelectSingleByPerson string
	UpdateByAdmin string
	UpdateByManager string
	UpdateByPerson string
	DeleteByManage string
	DeleteByAdmin string
	InsertByAdmin string
	InsertByManager string
	SelectYearPerformanceByAdmin string
	SelectYearPerformanceByManage string
	SelectYearPerformanceByPerson string
}
func MysqlSingleton() *sqlQuery {
	queryMux = new(sync.Mutex)
	if sqlQueryInstance == nil {
		queryMux.Lock()
		defer queryMux.Unlock()
		if sqlQueryInstance == nil {
			sqlQueryInstance = &sqlQuery{}
			AddUserQuery()
			AddCompanyQuery()
			AddUserPreferenceQuery()
			AddCompanyBanchQuery()
			AddShiftQuery()
			AddShiftChangeQuery()
			AddShiftOverTimeQuery()
			AddForgetPunchQuery()
			AddLateExcusedQuery()
			AddDayOffQuery()
			AddBanchStyleQuery()
			AddBanchRuleQuery()
			AddQuitWorkUserQuery()
			AddWaitCompanyReply()
			AddWeekendSetting()
			AddWorkTime()
			AddPaidVocation()
			AddLog()
			AddPerformance()
			return sqlQueryInstance
		}
	}
	return sqlQueryInstance
}