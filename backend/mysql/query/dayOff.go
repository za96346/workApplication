package query

func AddDayOffQuery() {
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