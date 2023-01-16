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
	UpdateByAdmin string
	UpdateByManager string
	UpdateByPerson string
	DeleteByManage string
	DeleteByAdmin string
	InsertByAdmin string
	InsertByManager string
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
			addWorkTime()
			addPaidVocation()
			addLog()
			addPerformance()
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
	sqlQueryInstance.User.UpdateBoss = `
		update user
		set
			companyCode=?,
			banch=?,
			permession=?,
			lastModify=?
		where userId=?;
	`
	sqlQueryInstance.User.SelectAllByAdmin = `
		select
			u.userId,
			u.companyCode,
			u.userName,
			u.employeeNumber,
			u.onWorkDay,
			u.banch,
			u.permession,
			IF(q.userId is null = 1, 'on', 'off') AS workState,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user u
		left join quitWorkUser q
			on u.userId=q.userId
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where
			(u.companyCode=? or q.companyCode=?)
			and
			u.userName=if(?='' or ?=null, u.userName, ?)
		;
	`
	sqlQueryInstance.User.SelectAllByManager = `
	select
		u.userId,
		u.companyCode,
		u.userName,
		u.employeeNumber,
		u.onWorkDay,
		u.banch,
		u.permession,
		IF(q.userId is null = 1, 'on', 'off') AS workState,
		ifnull(cb.banchName, '') as banchName,
		ifnull(c.companyId, -1),
		ifnull(c.companyName, '') as companyName
	from user u
	left join quitWorkUser q
		on u.userId=q.userId
	left join companyBanch cb
		on cb.id=u.banch
	left join company c
		on u.companyCode=c.companyCode
	where
		(u.companyCode=? or q.companyCode=?)
		and
		(q.banch=? or u.banch=?)
		and
		u.userName=if(?='' or ?=null, u.userName, ?)
	;
	`;
	sqlQueryInstance.User.UpdateCompanyUser = `
	update user
	set
		employeeNumber=?,
		companyCode=?,
		onWorkDay=?,
		banch=?,
		permession=?,
		lastModify=?
	where userId=?
	and(
		companyCode=?
		or companyCode is null
		or companyCode=''
	);
	`
	sqlQueryInstance.User.SelectAllByUserIdAndCompanyCode = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.companyCode=? and u.userId=?;
	`
	sqlQueryInstance.User.SelectAllByCompanyCode = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.companyCode=?;
	`
	sqlQueryInstance.User.SelectAll = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode;
	`;
	sqlQueryInstance.User.SelectSingleByUserId = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.userId=?;
	`;
	sqlQueryInstance.User.SelectSingleByAccount = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.account=?;
	`;
	sqlQueryInstance.User.Delete = `delete from user where userId=?;`;
	sqlQueryInstance.User.SelectAllByBanchId = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.banch=?;
	`;
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
	sqlQueryInstance.CompanyBanch.UpdateByCompanyCode = `
	update companyBanch b
	left join company c
	on b.companyId=c.companyId
	set
		b.banchName=?,
		b.banchShiftStyle=?,
		b.lastModify=?
	where b.id=? and c.companyCode=?;
	`
	sqlQueryInstance.CompanyBanch.DeleteByCompanyCode = `
	delete b from companyBanch b
	left join company c
	on b.companyId=c.companyId
	where b.id=? and c.companyCode=?;
	`
	sqlQueryInstance.CompanyBanch.SelectByCompanyCodeAndBanchID = `
		select * from companyBanch where id=? and companyId=?;
	`
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
	sqlQueryInstance.BanchStyle.SelectByCompanyCode = `
	select bs.* from banchStyle as bs
	left join companyBanch cb
		on cb.id=bs.banchId
	left join company c
		on cb.companyId=c.companyId
	where
		bs.banchId=?
	and
		c.companyCode=?;
	`
	sqlQueryInstance.BanchStyle.UpdateByCompanyCode = `
		update banchStyle bs
		left join companyBanch cb
			on cb.id=bs.banchId
		left join company c
			on cb.companyId=c.companyId
		set
			bs.icon=?,
			bs.restTime=?,
			bs.timeRangeName=?,
			bs.onShiftTime=?,
			bs.offShiftTime=?,
			bs.lastModify=?
		where bs.styleId=? and c.companyCode=?;
	`
	sqlQueryInstance.BanchStyle.DeleteByCompanyCode = `
		delete bs from banchStyle bs
		left join companyBanch cb
			on cb.id=bs.banchId
		left join company c
			on cb.companyId=c.companyId
		where bs.styleId=? and c.companyCode=?;
	`
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
	sqlQueryInstance.QuitWorkUser.InsertBySelectUser = `
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
			)
			select
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
			from user
			where userId=? and companyCode=?;
	`
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
	sqlQueryInstance.QuitWorkUser.DeleteByJoinUser = `
		delete qw from quitWorkUser qw
		left join user u
		on u.userId=qw.userId
		where
			qw.userId=?
		and
			qw.companyCode=?
		and
			(u.companyCode is null or u.companyCode='');
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
func addWorkTime () {
	sqlQueryInstance.WorkTime.InsertAll = `
		insert into workTime(
			userId,
			year,
			month,
			workHours,
			timeOff,
			usePaidVocation,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?
		);
	`;
	sqlQueryInstance.WorkTime.UpdateSingle = `
		update workTime
		left join user
		on user.userId=workTime.userId
		set
			workTime.year=?,
			workTime.month=?,
			workTime.workHours=?,
			workTime.timeOff=?,
			workTime.usePaidVocation=?,
			workTime.lastModify=?
		where workTime.workTimeId=? and user.companyCode=?;
	`;
	sqlQueryInstance.WorkTime.SelectAll = `
		select
			workTime.*
			user.userName,
			user.banch,
			user.employeeNumber,
			ifnull(companyBanch.banchName, '')
		from workTime
		left join user
		on workTime.userId=user.userId
		left join companyBanch
		on user.banch=companyBanch.id
		where user.companyCode=?;
	`;
	sqlQueryInstance.WorkTime.Delete = `delete from workTime where workTimeId=?;`;
	sqlQueryInstance.WorkTime.SelectAllByUserId = `
		select
			workTime.*,
			user.userName,
			user.banch,
			user.employeeNumber,
			ifnull(companyBanch.banchName, '')
		from workTime
		left join user
		on workTime.userId=user.userId
		left join companyBanch
		on user.banch=companyBanch.id
		where
			workTime.userId=? and
			user.companyCode=?;
		`;
	sqlQueryInstance.WorkTime.SelectAllByTime = `
		select
			workTime.*,
			user.userName,
			user.banch,
			user.employeeNumber,
			ifnull(companyBanch.banchName, '')
		from workTime
		left join user
		on workTime.userId=user.userId
		left join companyBanch
		on user.banch=companyBanch.id
		where
			workTime.year=? and
			workTime.month=? and
			user.companyCode=?;
	`;
	sqlQueryInstance.WorkTime.SelectAllByPrimaryKey = `
	select
		workTime.*,
		user.userName,
		user.banch,
		user.employeeNumber,
		ifnull(companyBanch.banchName, '')
	from workTime
	inner join user
	on workTime.userId=user.userId
	left join companyBanch
	on user.banch=companyBanch.id
	where
		workTime.year=? and
		workTime.month=? and
		workTime.userId=? and
		user.companyCode=?;
	`;
	sqlQueryInstance.WorkTime.DeleteByCompanyAndId = `
	delete wt from workTime wt
		left join user
		on
			user.userId=wt.userId
		where
			wt.workTimeId=?
			and
			user.companyCode=?
	;`;

}
func addPaidVocation () {
	sqlQueryInstance.PaidVocation.InsertAll = `
		insert into paidVocation(
			userId,
			year,
			count,
			createTime,
			lastModify
		) values (
			?, ?, ?, ?, ?
		);
	`;
	sqlQueryInstance.PaidVocation.UpdateSingle = `
		update paidVocation
		set
			year=?,
			count=?,
			lastModify=?
		where paidVocationId=?;
	`;
	sqlQueryInstance.PaidVocation.SelectAll = `select * from paidVocation;`
	sqlQueryInstance.PaidVocation.Delete = `delete from paidVocation where paidVocationId=?;`
	sqlQueryInstance.PaidVocation.SelectAllByUserId = `select * from paidVocation where userId=?;`
	sqlQueryInstance.PaidVocation.SelectAllByTime = `select * from paidVocation where year=?;`
}
func addLog () {
	sqlQueryInstance.Log.InsertAll = `
		insert into log(
			msg,
			createTime,
			lastModify
		) values (
			?,?,?
		)
	;`;
}
func addPerformance(){
	sqlQueryInstance.Performance.SelectAllByAdmin = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			(u.companyCode=?
			or qu.companyCode=?)
			and 
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) >= ?
			and
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) <= ?
			and u.userName=if(?='' or ?=null, u.userName, ?)
			order by p.year asc, p.month asc;
	`;
	sqlQueryInstance.Performance.SelectAllByManager = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			(u.companyCode=?
				or qu.companyCode=?)
			and (p.banchId=?
				or p.banchName=?)
				and 
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) >= ?
			and
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) <= ?
			and u.userName=if(?='' or ?=null, u.userName, ?)
			order by p.year asc, p.month asc;
	`;
	sqlQueryInstance.Performance.SelectAllByPerson = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			p.userId=? 
			and 
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) >= ?
			and
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) <= ?
		order by p.year asc, p.month asc;
	`;
	sqlQueryInstance.Performance.UpdateByAdmin = `
		update performance p
		left join user u
			on u.userId=p.userId
		left join quitWorkUser qu
			on qu.userId=p.userId
		set
			banchId=?,
			goal=?,
			attitude=?,
			efficiency=?,
			professional=?,
			directions=?,
			beLate=?,
			dayOffNotOnRule=?,
			banchName=?,
			p.lastModify=?
		where
			p.performanceId=?
			and (qu.companyCode=?
			or u.companyCode=?);
	`
	sqlQueryInstance.Performance.UpdateByManager = `
		update performance p
		left join user u
			on u.userId=p.userId
		left join quitWorkUser qu
			on qu.userId=p.userId
		set
			banchId=?,
			goal=?,
			attitude=?,
			efficiency=?,
			professional=?,
			directions=?,
			beLate=?,
			dayOffNotOnRule=?,
			banchName=?,
			p.lastModify=?
		where
			p.performanceId=?
			and (qu.companyCode=?
				or u.companyCode=?)
			and (p.banchId=? or p.banchName=?);
	`;
	sqlQueryInstance.Performance.InsertAll = `
		insert into performance (
			userId,
			year,
			month,
			banchId,
			goal,
			attitude,
			efficiency,
			professional,
			directions,
			beLate,
			dayOffNotOnRule,
			banchName,
			createTime,
			lastModify
		) values (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		);
	`
	sqlQueryInstance.Performance.DeleteByAdmin = `
		delete p from performance p
		where performanceId=?;
	`
	sqlQueryInstance.Performance.DeleteByManage = `
		delete p from performance p
		where performanceId=? && p.banchId=? && p.userId!=?;
	`
	sqlQueryInstance.Performance.UpdateByPerson = `
		update performance p
		set
			goal=?,
			p.lastModify=?
		where p.performanceId=? and p.userId=?;
	`;
}