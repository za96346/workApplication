package query

func AddCompanyQuery() {
	sqlQueryInstance.Company.InsertAll = `
	insert into company(
		companyCode,
		companyName,
		companyLocation,
		companyPhoneNumber,
		bossId,
		settlementDate,
		termStart,
		termEnd,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.Company.UpdateSingle = `
	update company
	set
		companyName=?,
		companyLocation=?,
		companyPhoneNumber=?,
		bossId=?,
		settlementDate=?,
		termStart=?,
		termEnd=?,
		lastModify=?
	where companyId=?;
	`;
	sqlQueryInstance.Company.SelectSingleByCompanyId = `select * from company where companyId = ?;`;
	sqlQueryInstance.Company.SelectSingleByCompanyCode = `select * from company where companyCode = ?;`;
	sqlQueryInstance.Company.SelectAll = `select * from company;`;
	sqlQueryInstance.Company.Delete = `delete from company where companyId = ?;`;
}