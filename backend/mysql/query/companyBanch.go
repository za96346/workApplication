package query

func AddCompanyBanchQuery() {
	sqlQueryInstance.CompanyBanch.InsertAll = `
	insert into companyBanch(
		companyId,
		banchName,
		banchShiftStyle,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.CompanyBanch.UpdateSingle = `
	update companyBanch
	set
		banchName=?,
		banchShiftStyle=?,
		lastModify=?
	where id=?;
	`;
	sqlQueryInstance.CompanyBanch.UpdateByCompanyCode = `
	update companyBanch b
	left join company c
	on b.companyId=c.companyId
	set
		b.banchName=?,
		b.banchShiftStyle=?,
		b.lastModify=?
	where b.id=? and c.companyCode=?;
	`
	sqlQueryInstance.CompanyBanch.DeleteByCompanyCode = `
	delete b from companyBanch b
	left join company c
	on b.companyId=c.companyId
	where b.id=? and c.companyCode=?;
	`
	sqlQueryInstance.CompanyBanch.SelectByCompanyCodeAndBanchID = `
		select
			cb.*
			count(u.userId) as userTotal
		from companyBanch as cb
		left join user u
			on u.banch=cb.id
		where id=? and companyId=?;
	`
	sqlQueryInstance.CompanyBanch.SelectAll = `
		select
			cb.*
			count(u.userId) as userTotal
		from companyBanch as cb
		left join user u
			on u.banch=cb.id 
		from companyBanch`;
	sqlQueryInstance.CompanyBanch.Delete = `delete from companyBanch where id = ?;`;
	sqlQueryInstance.CompanyBanch.SelectSingleByCompanyId = `
		select
			cb.*,
			count(u.userId) as userTotal
		from companyBanch as cb
		left join user u
		on u.banch=cb.id
		where 
			u.companyCode=(
				select
					companyCode
				from company as c
				where c.companyId=?
			)
		group by cb.id;
	`
	sqlQueryInstance.CompanyBanch.SelectSingleById = `
		select
			cb.*
			count(u.userId) as userTotal
		from companyBanch as cb
		left join user u
			on u.banch=cb.id
		from companyBanch where id = ?;`
}