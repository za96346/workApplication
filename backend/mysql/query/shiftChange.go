package query

func AddShiftChangeQuery() {
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