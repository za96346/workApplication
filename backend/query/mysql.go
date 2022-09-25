package query

import (
	"fmt"
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
}
type queryCommonColumn struct {
	InsertAll string
	SelectAll string
	Delete string
	update string
}

type userQuery struct {
	queryCommonColumn
	InsertSingle string
	SelectSingleByUserId string
	SelectSingleByAccount string
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
}
type shiftQuery struct {
	queryCommonColumn
	SelectSingleByUserId string
	SelectSingleByShiftId string
}
type shiftChangeQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
}
type shiftOverTimeQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
}
type forgetPunchQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
}
type lateExcusedQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
}
type  dayOffQuery struct {
	queryCommonColumn
	SelectSingleByCaseId string
}

func MysqlSingleton() *sqlQuery {
	queryMux = new(sync.Mutex)
	if sqlQueryInstance == nil {
		queryMux.Lock()
		if sqlQueryInstance == nil {
			sqlQueryInstance = &sqlQuery{}
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
			return sqlQueryInstance
		}
	}
	return sqlQueryInstance
}

func addUserQuery() {
	fmt.Println(sqlQueryInstance)
	sqlQueryInstance.User.InsertAll = `insert into user(
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
	sqlQueryInstance.User.SelectAll = `select * from user;`;
	sqlQueryInstance.User.SelectSingleByUserId = `select * from user where userId=?;`;
	sqlQueryInstance.User.SelectSingleByAccount = `select * from user where account=?;`;
	sqlQueryInstance.User.Delete = `delete from user where userId=?;`;
}
func addUserPreferenceQuery() {
	sqlQueryInstance.UserPreference.InsertAll = `insert into userPreference(
		userId,
		style,
		fontSize,
		selfPhoto,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.UserPreference.SelectAll = `select * from userPreference;`;
	sqlQueryInstance.UserPreference.Delete = `delete from userPreference where userId = ?;`;
	sqlQueryInstance.UserPreference.SelectSingleByUserId = `select * from userPreference where userId = ?;`;
}
func addCompanyQuery() {
	sqlQueryInstance.Company.InsertAll = `insert into company(
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
	sqlQueryInstance.Company.SelectSingleByCompanyId = `select * from company where companyId = ?;`;
	sqlQueryInstance.Company.SelectSingleByCompanyCode = `select * from company where companyCode = ?;`;
	sqlQueryInstance.Company.SelectAll = `select * from company;`;
	sqlQueryInstance.Company.Delete = `delete from company where companyId = ?;`;
}
func addCompanyBanchQuery() {
	sqlQueryInstance.CompanyBanch.InsertAll = `insert into companyBanch(
		companyId,
		banchName,
		banchShiftStyle,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.CompanyBanch.SelectAll = `select * from companyBanch`;
	sqlQueryInstance.CompanyBanch.Delete = `delete from companyBanch where id = ?;`;
	sqlQueryInstance.CompanyBanch.SelectSingleByCompanyId = `select * from companyBanch where companyId = ?;`
	sqlQueryInstance.CompanyBanch.SelectSingleById = `select * from companyBanch where id = ?;`
}
func addShiftQuery() {
	sqlQueryInstance.Shift.InsertAll = `insert into shift(
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
	sqlQueryInstance.Shift.SelectSingleByUserId = `select * from shift where userId=?;`;
	sqlQueryInstance.Shift.SelectSingleByShiftId = `select * from shift where shiftId=?;`;
	sqlQueryInstance.Shift.SelectAll = `select * from shift;`;
	sqlQueryInstance.Shift.Delete = `delete from shift where shiftId = ?;`;
}
func addShiftChangeQuery() {
	sqlQueryInstance.ShiftChange.InsertAll = `insert into shiftChange(
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
	sqlQueryInstance.ShiftChange.SelectAll = `select * from shiftChange;`;
	sqlQueryInstance.ShiftChange.Delete = `delete from shiftChange where caseId = ?;`;
	sqlQueryInstance.ShiftChange.SelectSingleByCaseId = `select * from shiftChange where caseId = ?;`;
}
func  addShiftOverTimeQuery() {
	sqlQueryInstance.ShiftOverTime.InsertAll = `insert into shiftOverTime(
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
	sqlQueryInstance.ShiftOverTime.SelectAll = `select * from shiftOverTime;`;
	sqlQueryInstance.ShiftOverTime.Delete = `delete from shiftOverTime where caseId = ?;`;
	sqlQueryInstance.ShiftOverTime.SelectSingleByCaseId = `select * from shiftOverTime where caseId = ?;`;
}
func addForgetPunchQuery() {
	sqlQueryInstance.ForgetPunch.InsertAll = `insert into forgetPunch(
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
	sqlQueryInstance.ForgetPunch.SelectAll = `select * from forgetPunch;`;
	sqlQueryInstance.ForgetPunch.Delete = `delete from forgetPunch where caseId = ?;`;
	sqlQueryInstance.ForgetPunch.SelectSingleByCaseId = `select * from forgetPunch where caseId = ?;`;
}
func addLateExcusedQuery() {
	sqlQueryInstance.LateExcused.InsertAll = `insert into lateExcused(
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
	sqlQueryInstance.LateExcused.SelectAll = `select * from lateExcused;`;
	sqlQueryInstance.LateExcused.Delete = `delete from lateExcused where caseId = ?;`;
	sqlQueryInstance.LateExcused.SelectSingleByCaseId = `select * from lateExcused where caseId = ?;`;
}
func addDayOffQuery() {
	sqlQueryInstance.DayOff.InsertAll = `insert into dayOff(
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
	sqlQueryInstance.DayOff.SelectAll = `select * from dayOff;`;
	sqlQueryInstance.DayOff.Delete = `delete from dayOff where caseId = ?;`;
	sqlQueryInstance.DayOff.SelectSingleByCaseId = `select * from dayOff where caseId = ?;`;
}