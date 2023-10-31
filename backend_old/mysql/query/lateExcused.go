package query

func AddLateExcusedQuery() {
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