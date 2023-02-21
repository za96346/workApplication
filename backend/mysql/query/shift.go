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
	sqlQueryInstance.Shift.Delete = `delete from shift where shiftId = ?;`;
}