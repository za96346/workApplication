package query

func AddQuitWorkUserQuery() {
	sqlQueryInstance.QuitWorkUser.InsertAll = `
		insert into quitWorkUser(
			userId,
			companyCode,
			userName,
			employeeNumber,
			account,
			onWorkDay,
			banch,
			permession,
			monthSalary,
			partTimeSalary,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		);
	`;
	sqlQueryInstance.QuitWorkUser.InsertBySelectUser = `
			insert into quitWorkUser(
				userId,
				companyCode,
				userName,
				employeeNumber,
				account,
				onWorkDay,
				banch,
				permession,
				monthSalary,
				partTimeSalary,
				createTime,
				lastModify
			)
			select
				userId,
				companyCode,
				userName,
				employeeNumber,
				account,
				onWorkDay,
				banch,
				permession,
				monthSalary,
				partTimeSalary,
				createTime,
				lastModify
			from user
			where userId=? and companyCode=?;
	`
	sqlQueryInstance.QuitWorkUser.UpdateSingle = `
		update quitWorkUser
		set
			userId=?,
			companyCode=?,
			userName=?,
			employeeNumber=?,
			account=?,
			onWorkDay=?,
			banch=?,
			permession=?,
			monthSalary=?,
			partTimeSalary=?,
			createTime=?,
			lastModify=?
		where quitId=?;
	`
	sqlQueryInstance.QuitWorkUser.DeleteByJoinUser = `
		delete qw from quitWorkUser qw
		left join user u
		on u.userId=qw.userId
		where
			qw.userId=?
		and
			qw.companyCode=?
		and
			(u.companyCode is null or u.companyCode='');
	`
	sqlQueryInstance.QuitWorkUser.SelectAll = `select * from quitWorkUser;`
	sqlQueryInstance.QuitWorkUser.SelectAllByCompanyCode = `select * from quitWorkUser where companyCode=?;`
	sqlQueryInstance.QuitWorkUser.SelectSingleByUserId = `select * from quitWorkUser where userId=?;`
	sqlQueryInstance.QuitWorkUser.SelectSingleByQuitId = `select * from quitWorkUser where quitId=?;`
	sqlQueryInstance.QuitWorkUser.SelectSingleByCompanyCodeAndUserId = `select * from quitWorkUser where companyCode = ? and userId = ?;`
	sqlQueryInstance.QuitWorkUser.Delete = `delete from quitWorkUser where quitId=?;`;
}