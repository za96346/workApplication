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
	BanchStyle banchStyle
	BanchRule banchRule
	QuitWorkUser quitWorkUser
	WaitCompanyReply waitCompanyReply
	WeekendSetting weekendSetting
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

func MysqlSingleton() *sqlQuery {
	queryMux = new(sync.Mutex)
	if sqlQueryInstance == nil {
		queryMux.Lock()
		defer queryMux.Unlock()
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
			addBanchStyleQuery()
			addBanchRuleQuery()
			addQuitWorkUserQuery()
			addWaitCompanyReply()
			addWeekendSetting()
			return sqlQueryInstance
		}
	}
	return sqlQueryInstance
}

func addUserQuery() {
	fmt.Println(sqlQueryInstance)
	sqlQueryInstance.User.InsertAll = `
	insert into user(
		companyCode,
		account,
		password,
		userName,
		employeeNumber,
		onWorkDay,
		banch,
		permession,
		createTime,
		lastModify,
		monthSalary,
		partTimeSalary
		) values(
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.User.UpdateSingle = `
		update user
		set
			employeeNumber=?,
			companyCode=?,
			password=?,
			userName=?,
			onWorkDay=?,
			banch=?,
			permession=?,
			lastModify=?,
			monthSalary=?,
			partTimeSalary=?
		where userId=?;
	`;
	sqlQueryInstance.User.SelectAllByCompanyCode = `select * from user where companyCode=?;`
	sqlQueryInstance.User.SelectAll = `select * from user;`;
	sqlQueryInstance.User.SelectSingleByUserId = `select * from user where userId=?;`;
	sqlQueryInstance.User.SelectSingleByAccount = `select * from user where account=?;`;
	sqlQueryInstance.User.Delete = `delete from user where userId=?;`;
	sqlQueryInstance.User.SelectAllByBanchId = `select * from user where banch = ?;`;
}
func addUserPreferenceQuery() {
	sqlQueryInstance.UserPreference.InsertAll = `
	insert into userPreference(
		userId,
		style,
		fontSize,
		selfPhoto,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.UserPreference.UpdateSingle = `
	update userPreference
	set
		style=?,
		fontSize=?,
		selfPhoto=?,
		lastModify=?
	where userId=?;
	`;
	sqlQueryInstance.UserPreference.SelectAll = `select * from userPreference;`;
	sqlQueryInstance.UserPreference.Delete = `delete from userPreference where userId = ?;`;
	sqlQueryInstance.UserPreference.SelectSingleByUserId = `select * from userPreference where userId = ?;`;
}
func addCompanyQuery() {
	sqlQueryInstance.Company.InsertAll = `
	insert into company(
		companyCode,
		companyName,
		companyLocation,
		companyPhoneNumber,
		bossId,
		settlementDate,
		termStart,
		termEnd,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.Company.UpdateSingle = `
	update company
	set
		companyName=?,
		companyLocation=?,
		companyPhoneNumber=?,
		bossId=?,
		settlementDate=?,
		termStart=?,
		termEnd=?,
		lastModify=?
	where companyId=?;
	`;
	sqlQueryInstance.Company.SelectSingleByCompanyId = `select * from company where companyId = ?;`;
	sqlQueryInstance.Company.SelectSingleByCompanyCode = `select * from company where companyCode = ?;`;
	sqlQueryInstance.Company.SelectAll = `select * from company;`;
	sqlQueryInstance.Company.Delete = `delete from company where companyId = ?;`;
}
func addCompanyBanchQuery() {
	sqlQueryInstance.CompanyBanch.InsertAll = `
	insert into companyBanch(
		companyId,
		banchName,
		banchShiftStyle,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.CompanyBanch.UpdateSingle = `
	update companyBanch
	set
		banchName=?,
		banchShiftStyle=?,
		lastModify=?
	where id=?;
	`;
	sqlQueryInstance.CompanyBanch.SelectAll = `select * from companyBanch`;
	sqlQueryInstance.CompanyBanch.Delete = `delete from companyBanch where id = ?;`;
	sqlQueryInstance.CompanyBanch.SelectSingleByCompanyId = `select * from companyBanch where companyId = ?;`
	sqlQueryInstance.CompanyBanch.SelectSingleById = `select * from companyBanch where id = ?;`
}
func addShiftQuery() {
	sqlQueryInstance.Shift.InsertAll = `
	insert into shift(
		userId,
		banchStyleId,
		onShiftTime,
		offShiftTime,
		restTime,
		punchIn,
		punchOut,
		specifyTag,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.Shift.UpdateSingle = `
	update shift
	set
		banchStyleId=?,
		onShiftTime=?,
		offShiftTime=?,
		restTime=?,
		punchIn=?,
		punchOut=?,
		specifyTag=?,
		lastModify=?
	where shiftId=?;
	`;

	sqlQueryInstance.Shift.SelectSingleByUserId = `select * from shift where userId=?;`;
	sqlQueryInstance.Shift.SelectSingleByShiftId = `select * from shift where shiftId=?;`;
	sqlQueryInstance.Shift.SelectAll = `select * from shift;`;
	sqlQueryInstance.Shift.Delete = `delete from shift where shiftId = ?;`;
}
func addShiftChangeQuery() {
	sqlQueryInstance.ShiftChange.InsertAll = `
	insert into shiftChange(
		initiatorShiftId,
		requestedShiftId,
		reason,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.ShiftChange.UpdateSingle = `
	update shiftChange
	set
		initiatorShiftId=?,
		requestedShiftId=?,
		reason=?,
		caseProcess=?,
		specifyTag=?,
		lastModify=?
	where caseId=?;
	`;
	sqlQueryInstance.ShiftChange.SelectAll = `select * from shiftChange;`;
	sqlQueryInstance.ShiftChange.Delete = `delete from shiftChange where caseId = ?;`;
	sqlQueryInstance.ShiftChange.SelectSingleByCaseId = `select * from shiftChange where caseId = ?;`;
	sqlQueryInstance.ShiftChange.SelectAllByInitiatorShiftId = `select * from shiftChange where initiatorShiftId = ?;`
	sqlQueryInstance.ShiftChange.SelectAllByRequestedShiftId = `select * from shiftChange where requestedShiftId = ?;`
}
func  addShiftOverTimeQuery() {
	sqlQueryInstance.ShiftOverTime.InsertAll = `
	insert into shiftOverTime(
		shiftId,
		initiatorOnOverTime,
		initiatorOffOverTime,
		reason,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.ShiftOverTime.UpdateSingle = `
	update shiftOverTime
	set
		initiatorOnOverTime=?,
		initiatorOffOverTime=?,
		reason=?,
		caseProcess=?,
		specifyTag=?,
		lastModify=?
	where caseId=?;
	`;
	sqlQueryInstance.ShiftOverTime.SelectAll = `select * from shiftOverTime;`;
	sqlQueryInstance.ShiftOverTime.Delete = `delete from shiftOverTime where caseId = ?;`;
	sqlQueryInstance.ShiftOverTime.SelectSingleByCaseId = `select * from shiftOverTime where caseId = ?;`;
	sqlQueryInstance.ShiftOverTime.SelectAllByShiftId = `select * from shiftOverTime where shiftId = ?;`;
}
func addForgetPunchQuery() {
	sqlQueryInstance.ForgetPunch.InsertAll = `
	insert into forgetPunch(
		shiftId,
		targetPunch,
		reason,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.ForgetPunch.UpdateSingle = `
	update forgetPunch
	set
		targetPunch=?,
		reason=?,
		caseProcess=?,
		specifyTag=?,
		lastModify=?
	where caseId=?;
	`;
	sqlQueryInstance.ForgetPunch.SelectAll = `select * from forgetPunch;`;
	sqlQueryInstance.ForgetPunch.Delete = `delete from forgetPunch where caseId = ?;`;
	sqlQueryInstance.ForgetPunch.SelectSingleByCaseId = `select * from forgetPunch where caseId = ?;`;
	sqlQueryInstance.ForgetPunch.SelectAllByShiftId = `select * from forgetPunch where shiftId = ?;`
}
func addLateExcusedQuery() {
	sqlQueryInstance.LateExcused.InsertAll = `
	insert into lateExcused(
		shiftId,
		lateExcusedType,
		reason,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.LateExcused.UpdateSingle = `
	update lateExcused
	set
		lateExcusedType=?,
		reason=?,
		caseProcess=?,
		specifyTag=?,
		lastModify=?
	where caseId=?;
	`;
	sqlQueryInstance.LateExcused.SelectAll = `select * from lateExcused;`;
	sqlQueryInstance.LateExcused.Delete = `delete from lateExcused where caseId = ?;`;
	sqlQueryInstance.LateExcused.SelectSingleByCaseId = `select * from lateExcused where caseId = ?;`;
	sqlQueryInstance.LateExcused.SelectAllByShiftId = `select * from lateExcused where shiftId = ?;`;
}
func addDayOffQuery() {
	sqlQueryInstance.DayOff.InsertAll = `
	insert into dayOff(
		shiftId,
		dayOffType,
		reason,
		caseProcess,
		specifyTag,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.DayOff.UpdateSingle = `
	update dayOff
	set
		dayOffType=?,
		reason=?,
		caseProcess=?,
		specifyTag=?,
		lastModify=?
	where caseId=?;
	`;
	sqlQueryInstance.DayOff.SelectAll = `select * from dayOff;`;
	sqlQueryInstance.DayOff.Delete = `delete from dayOff where caseId = ?;`;
	sqlQueryInstance.DayOff.SelectSingleByCaseId = `select * from dayOff where caseId = ?;`;
	sqlQueryInstance.DayOff.SelectAllByShiftId = `select * from dayOff where shiftId = ?;`;
}
func addBanchStyleQuery() {
	sqlQueryInstance.BanchStyle.InsertAll = `
		insert into banchStyle(
			banchId,
			icon,
			restTime,
			timeRangeName,
			onShiftTime,
			offShiftTime,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?
		);
	`
	sqlQueryInstance.BanchStyle.UpdateSingle = `
		update banchStyle
		set
			icon=?,
			restTime=?,
			timeRangeName=?,
			onShiftTime=?,
			offShiftTime=?,
			lastModify=?
		where styleId=?;
	`;
	sqlQueryInstance.BanchStyle.SelectSingleByStyleId = `select * from banchStyle where styleId = ?;`;
	sqlQueryInstance.BanchStyle.SelectAll = `select * from banchStyle;`;
	sqlQueryInstance.BanchStyle.Delete = `delete from banchStyle where styleId=?;`;
	sqlQueryInstance.BanchStyle.SelectAllByBanchId = `select * from banchStyle where banchId = ?;`;
}
func addBanchRuleQuery() {
	sqlQueryInstance.BanchRule.InsertAll = `
		insert into banchRule(
			banchId,
			maxPeople,
			minPeople,
			weekDay,
			weekType,
			onShiftTime,
			offShiftTime,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?, ?
		);
	`
	sqlQueryInstance.BanchRule.UpdateSingle = `
		update banchRule
		set
			maxPeople=?,
			minPeople=?,
			weekDay=?,
			weekType=?,
			onShiftTime=?,
			offShiftTime=?,
			lastModify=?
		where ruleId=?;
	`
	sqlQueryInstance.BanchRule.SelectSingleByRuleId = `select * from banchRule where ruleId = ?;`;
	sqlQueryInstance.BanchRule.SelectAll = `select * from banchRule;`;
	sqlQueryInstance.BanchRule.Delete = `delete from banchRule where ruleId=?;`;
	sqlQueryInstance.BanchRule.SelectAllByBanchId = `select * from banchRule where banchId = ?;`;
}

func addQuitWorkUserQuery() {
	sqlQueryInstance.QuitWorkUser.InsertAll = `
		insert into quitWorkUser(
			userId,
			companyCode,
			userName,
			employeeNumber,
			account,
			onWorkDay,
			banch,
			permession,
			monthSalary,
			partTimeSalary,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		);
	`;
	sqlQueryInstance.QuitWorkUser.UpdateSingle = `
		update quitWorkUser
		set
			userId=?,
			companyCode=?,
			userName=?,
			employeeNumber=?,
			account=?,
			onWorkDay=?,
			banch=?,
			permession=?,
			monthSalary=?,
			partTimeSalary=?,
			createTime=?,
			lastModify=?
		where quitId=?;
	`
	sqlQueryInstance.QuitWorkUser.SelectAll = `select * from quitWorkUser;`
	sqlQueryInstance.QuitWorkUser.SelectAllByCompanyCode = `select * from quitWorkUser where companyCode=?;`
	sqlQueryInstance.QuitWorkUser.SelectSingleByUserId = `select * from quitWorkUser where userId=?;`
	sqlQueryInstance.QuitWorkUser.SelectSingleByQuitId = `select * from quitWorkUser where quitId=?;`
	sqlQueryInstance.QuitWorkUser.SelectSingleByCompanyCodeAndUserId = `select * from quitWorkUser where companyCode = ? and userId = ?;`
	sqlQueryInstance.QuitWorkUser.Delete = `delete from quitWorkUser where quitId=?;`;
}

func addWaitCompanyReply () {
	sqlQueryInstance.WaitCompanyReply.InsertAll = `
		insert into waitCompanyReply(
			userId,
			companyId,
			specifyTag,
			isAccept,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?
		);
	`
	sqlQueryInstance.WaitCompanyReply.UpdateSingle = `
		update waitCompanyReply
		set
			specifyTag=?,
			isAccept=?,
			lastModify=?
		where waitId=?;
	`
	sqlQueryInstance.WaitCompanyReply.SelectAll = `select * from waitCompanyReply;`;
	sqlQueryInstance.WaitCompanyReply.SelectSingleByWaitId = `select * from waitCompanyReply where waitId = ?;`
	sqlQueryInstance.WaitCompanyReply.SelectAllByUserId = `select * from waitCompanyReply where userId = ?;`
	sqlQueryInstance.WaitCompanyReply.SelectAllByCompanyId = `select * from waitCompanyReply where companyId = ?;`
	sqlQueryInstance.WaitCompanyReply.SelectAllByCompanyIdAndUserId = `select * from waitCompanyReply where companyId = ? and userId = ?;`
	sqlQueryInstance.WaitCompanyReply.Delete = `delete from waitCompanyReply where waitId = ?;`
	sqlQueryInstance.WaitCompanyReply.SelectAllJoinUserTable = `
		select 
			w.waitId,
			w.userId,
			u.userName,
			w.companyId,
			w.specifyTag,
			w.isAccept,
			w.createTime,
			w.lastModify
		from waitCompanyReply as w left join user as u on w.userId=u.userId where w.companyId=?;
	`
}
func addWeekendSetting () {
	sqlQueryInstance.WeekendSetting.InsertAll = `
		insert into weekendSetting(
			companyId,
			date,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?
		);
	`
	sqlQueryInstance.WeekendSetting.UpdateSingle = `
		update weekendSetting
		set
			date=?,
			lastModify=?
		where weekendId=?;
	`
	sqlQueryInstance.WeekendSetting.SelectAll = `select * from weekendSetting;`;
	sqlQueryInstance.WeekendSetting.SelectSingleByWeekendId = `select * from weekendSetting where weekendId = ?;`
	sqlQueryInstance.WeekendSetting.SelectAllByCompanyId = `select * from weekendSetting where companyId = ?;`
	sqlQueryInstance.WeekendSetting.Delete = `delete from weekendSetting where weekendId = ?;`

}