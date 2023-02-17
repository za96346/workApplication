package query

func AddPaidVocation () {
	sqlQueryInstance.PaidVocation.InsertAll = `
		insert into paidVocation(
			userId,
			year,
			count,
			createTime,
			lastModify
		) values (
			?, ?, ?, ?, ?
		);
	`;
	sqlQueryInstance.PaidVocation.UpdateSingle = `
		update paidVocation
		set
			year=?,
			count=?,
			lastModify=?
		where paidVocationId=?;
	`;
	sqlQueryInstance.PaidVocation.SelectAll = `select * from paidVocation;`
	sqlQueryInstance.PaidVocation.Delete = `delete from paidVocation where paidVocationId=?;`
	sqlQueryInstance.PaidVocation.SelectAllByUserId = `select * from paidVocation where userId=?;`
	sqlQueryInstance.PaidVocation.SelectAllByTime = `select * from paidVocation where year=?;`
}