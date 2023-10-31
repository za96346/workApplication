package query

func  AddShiftOverTimeQuery() {
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