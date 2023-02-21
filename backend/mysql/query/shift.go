package query

func AddShiftQuery() {
	sqlQueryInstance.Shift.InsertAll = `
	insert into shift(
		userId,
		banchStyleId,
		year,
		month,
		Icon,
		onShiftTime,
		offShiftTime,
		restTime,
		punchIn,
		punchOut,
		specifyTag,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.Shift.UpdateSingle = `
	update shift
	set
		banchStyleId=?,
		icon=?,
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
	sqlQueryInstance.Shift.SelectAll = `
		select
			sf.*,
			u.userName,
			u.permession,
			u.banch,
			u.employeeNumber
		from shift sf
		left join user u
			on u.userId=sf.userId
		left join quitWorkUser qu
			on qu.userId=sf.userId
		where
			(u.banch=? or qu.banch=?)
		and 
			(u.companyCode=? or qu.companyCode=?)
		and sf.year=?
		and sf.month=?;
	`;
	sqlQueryInstance.Shift.SelectTotal = `
		select
			sf.userId,
			sf.year,
            sf.month,
			u.userName,
			u.permession,
			u.banch,
			u.employeeNumber,
			count(sc.caseProcess) as changeCocunt,
			count(so.caseProcess) as overTimeCount,
			count(fp.caseProcess) as forgetPunchCount,
			count(dof.caseProcess) as dayOffCount,
			count(led.caseProcess) as lateExcusedCount
		from shift as sf
		left join user u
			on u.userId=sf.userId
		left join quitWorkUser qu
			on qu.userId=sf.userId
		left join shiftChange sc
			on (sf.shiftId=sc.initiatorShiftId
			or sf.shiftId=sc.requestedShiftId)
			and sc.caseProcess='ok'
		left join shiftOverTime so
			on sf.shiftId=so.shiftId
			and so.caseProcess='ok'
		left join forgetPunch fp
			on sf.shiftId=fp.shiftId
			and fp.caseProcess='ok'
		left join dayOff dof
			on sf.shiftId=dof.shiftId
			and dof.caseProcess='ok'
		left join lateExcused led
			on sf.shiftId=led.shiftId
			and led.caseProcess='ok'
		where
			(u.banch=? or qu.banch=?)
		and 
			(u.companyCode=? or qu.companyCode=?)
		and sf.year=?
		and sf.month=?
		group by sf.userId, sf.year, sf.month;
	`
	sqlQueryInstance.Shift.Delete = `delete from shift where shiftId = ?;`;
}