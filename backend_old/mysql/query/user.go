package query

func AddUserQuery() {
	sqlQueryInstance.User.InsertAll = `
	insert into user(
		companyCode,
		account,
		password,
		userName,
		employeeNumber,
		onWorkDay,
		banch,
		permession,
		createTime,
		lastModify,
		monthSalary,
		partTimeSalary
		) values(
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.User.UpdateSingle = `
		update user
		set
			employeeNumber=?,
			companyCode=?,
			password=?,
			userName=?,
			onWorkDay=?,
			banch=?,
			permession=?,
			lastModify=?,
			monthSalary=?,
			partTimeSalary=?
		where userId=?;
	`;
	sqlQueryInstance.User.UpdateBoss = `
		update user
		set
			companyCode=?,
			banch=?,
			permession=?,
			lastModify=?
		where userId=?;
	`
	sqlQueryInstance.User.SelectAllByAdmin = `
		select
			u.userId,
			u.companyCode,
			u.userName,
			u.employeeNumber,
			u.onWorkDay,
			u.banch,
			u.permession,
			IF(q.userId is null = 1, 'on', 'off') AS workState,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user u
		left join quitWorkUser q
			on u.userId=q.userId
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where
			(u.companyCode=? or q.companyCode=?)
			and
			u.userName=if(?='' or ?=null, u.userName, ?)
			and
			u.permession=if(?='Y', 1, u.permession)
		;
	`
	sqlQueryInstance.User.SelectAllByManager = `
	select
		u.userId,
		u.companyCode,
		u.userName,
		u.employeeNumber,
		u.onWorkDay,
		u.banch,
		u.permession,
		IF(q.userId is null = 1, 'on', 'off') AS workState,
		ifnull(cb.banchName, '') as banchName,
		ifnull(c.companyId, -1),
		ifnull(c.companyName, '') as companyName
	from user u
	left join quitWorkUser q
		on u.userId=q.userId
	left join companyBanch cb
		on cb.id=u.banch
	left join company c
		on u.companyCode=c.companyCode
	where
		(u.companyCode=? or q.companyCode=?)
		and
		(q.banch=? or u.banch=?)
		and
		u.userName=if(?='' or ?=null, u.userName, ?)
	;
	`;
	sqlQueryInstance.User.UpdateCompanyUser = `
	update user
	set
		employeeNumber=?,
		companyCode=?,
		onWorkDay=?,
		banch=?,
		permession=?,
		lastModify=?
	where userId=?
	and(
		companyCode=?
		or companyCode is null
		or companyCode=''
	);
	`
	sqlQueryInstance.User.SelectAllByUserIdAndCompanyCode = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.companyCode=? and u.userId=?;
	`
	sqlQueryInstance.User.SelectAllByCompanyCode = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.companyCode=?;
	`
	sqlQueryInstance.User.SelectAll = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode;
	`;
	sqlQueryInstance.User.SelectSingleByUserId = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.userId=?;
	`;
	sqlQueryInstance.User.SelectSingleByAccount = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.account=?;
	`;
	sqlQueryInstance.User.Delete = `delete from user where userId=?;`;
	sqlQueryInstance.User.SelectAllByBanchId = `
		select
			u.*,
			ifnull(cb.banchName, '') as banchName,
			ifnull(c.companyId, -1),
			ifnull(c.companyName, '') as companyName
		from user as u
		left join companyBanch cb
			on cb.id=u.banch
		left join company c
			on u.companyCode=c.companyCode
		where u.banch=? and u.companyCode=?;
	`;
}