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
		select * from companyBanch where id=? and companyId=?;
	`
	sqlQueryInstance.CompanyBanch.SelectAll = `select * from companyBanch`;
	sqlQueryInstance.CompanyBanch.Delete = `delete from companyBanch where id = ?;`;
	sqlQueryInstance.CompanyBanch.SelectSingleByCompanyId = `select * from companyBanch where companyId = ?;`
	sqlQueryInstance.CompanyBanch.SelectSingleById = `select * from companyBanch where id = ?;`
}