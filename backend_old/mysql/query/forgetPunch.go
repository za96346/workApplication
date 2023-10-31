package query

func AddForgetPunchQuery() {
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