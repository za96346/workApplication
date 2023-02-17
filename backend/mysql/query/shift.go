package query

func AddShiftQuery() {
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