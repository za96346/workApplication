package query

func AddShiftQuery() {
	sqlQueryInstance.Shift.InsertAll = `
	insert into shift(
		userId,
		banchStyleId,
		banchId,
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
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
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

	sqlQueryInstance.Shift.SelectSingleByUserId = `
		select
			sf.*,
			u.userName,
			u.permession,
			u.banch,
			u.employeeNumber
		from shift sf
		left join user u
			on u.userId=sf.userId
		where sf.userId=?;`;
	sqlQueryInstance.Shift.SelectSingleByShiftId = `
		select
			sf.*,
			u.userName,
			u.permession,
			u.banch,
			u.employeeNumber
		from shift sf
		left join user u
			on u.userId=sf.userId
		where shiftId=?;`;
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
		left join companyBanch cb
			on cb.id=sf.banchId
		where
			sf.banchId=?
		and 
			cb.companyId=?
		and sf.year=?
		and sf.month=?;
	`;
	sqlQueryInstance.Shift.SelectTotal = `
		select
			sf.userId,
			sf.year,
			sf.month,
			sf.banchId,
			u.userName,
			u.permession,
			u.employeeNumber,
			count(sc.caseProcess) as changeCocunt,
			count(so.caseProcess) as overTimeCount,
			count(fp.caseProcess) as forgetPunchCount,
			count(dof.caseProcess) as dayOffCount,
			count(led.caseProcess) as lateExcusedCount,
			sum(
				timestampdiff(
					SECOND, sf.onShiftTime, sf.offShiftTime
				)
					- TIME_TO_SEC(sf.restTime)
			) / 60 / 60 as hours
		from shift as sf
		left join user u
			on u.userId=sf.userId
		left join companyBanch cb
			on cb.id=sf.banchId
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
			sf.banchId=?
		and 
			cb.companyId=?
		and sf.year=?
		and sf.month=?
		group by sf.userId, sf.year, sf.month, sf.banchId;
	`
	// 列 指的是 對人 總計
	sqlQueryInstance.Shift.SelectRowTotal = `
		select
			sf.userId,
			round(
				sum(
					timestampdiff(
						SECOND, sf.onShiftTime, sf.offShiftTime
					)
					- TIME_TO_SEC(sf.restTime)
				) / 60 / 60,
				2
			) as hours
		from shift as sf
		left join user u
			on u.userId=sf.userId
		left join companyBanch cb
			on cb.id=sf.banchId
		where
			sf.banchId=?
		and 
			cb.companyId=?
		and sf.year=?
		and sf.month=?
		group by sf.userId;
	`
	// 欄 指的是 對日期 總計
	sqlQueryInstance.Shift.SelectColumnTotal = `
		select
			date_format(sf.onShiftTime, '%Y-%m-%d') as dates,
			round(
				sum(
					timestampdiff(
						SECOND, sf.onShiftTime, sf.offShiftTime
					)
					- TIME_TO_SEC(sf.restTime)
				) / 60 / 60,
				2
			) as hours
		from shift as sf
		left join user u
			on u.userId=sf.userId
		left join companyBanch cb
			on cb.id=sf.banchId
		where
			sf.banchId=?
		and 
			cb.companyId=?
		and sf.year=?
		and sf.month=?
		group by dates;
	`
	sqlQueryInstance.Shift.Delete = `delete from shift where shiftId = ?;`;
}