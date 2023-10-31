package query

func AddWaitCompanyReply () {
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