package query

func AddWorkTime () {
	sqlQueryInstance.WorkTime.InsertAll = `
		insert into workTime(
			userId,
			year,
			month,
			workHours,
			timeOff,
			usePaidVocation,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?
		);
	`;
	sqlQueryInstance.WorkTime.UpdateSingle = `
		update workTime
		left join user
		on user.userId=workTime.userId
		set
			workTime.year=?,
			workTime.month=?,
			workTime.workHours=?,
			workTime.timeOff=?,
			workTime.usePaidVocation=?,
			workTime.lastModify=?
		where workTime.workTimeId=? and user.companyCode=?;
	`;
	sqlQueryInstance.WorkTime.SelectAll = `
		select
			workTime.*
			user.userName,
			user.banch,
			user.employeeNumber,
			ifnull(companyBanch.banchName, '')
		from workTime
		left join user
		on workTime.userId=user.userId
		left join companyBanch
		on user.banch=companyBanch.id
		where user.companyCode=?;
	`;
	sqlQueryInstance.WorkTime.Delete = `delete from workTime where workTimeId=?;`;
	sqlQueryInstance.WorkTime.SelectAllByUserId = `
		select
			workTime.*,
			user.userName,
			user.banch,
			user.employeeNumber,
			ifnull(companyBanch.banchName, '')
		from workTime
		left join user
		on workTime.userId=user.userId
		left join companyBanch
		on user.banch=companyBanch.id
		where
			workTime.userId=? and
			user.companyCode=?;
		`;
	sqlQueryInstance.WorkTime.SelectAllByTime = `
		select
			workTime.*,
			user.userName,
			user.banch,
			user.employeeNumber,
			ifnull(companyBanch.banchName, '')
		from workTime
		left join user
		on workTime.userId=user.userId
		left join companyBanch
		on user.banch=companyBanch.id
		where
			workTime.year=? and
			workTime.month=? and
			user.companyCode=?;
	`;
	sqlQueryInstance.WorkTime.SelectAllByPrimaryKey = `
	select
		workTime.*,
		user.userName,
		user.banch,
		user.employeeNumber,
		ifnull(companyBanch.banchName, '')
	from workTime
	inner join user
	on workTime.userId=user.userId
	left join companyBanch
	on user.banch=companyBanch.id
	where
		workTime.year=? and
		workTime.month=? and
		workTime.userId=? and
		user.companyCode=?;
	`;
	sqlQueryInstance.WorkTime.DeleteByCompanyAndId = `
	delete wt from workTime wt
		left join user
		on
			user.userId=wt.userId
		where
			wt.workTimeId=?
			and
			user.companyCode=?
	;`;

}