package query

func AddBanchStyleQuery() {
	sqlQueryInstance.BanchStyle.InsertAll = `
		insert into banchStyle(
			banchId,
			icon,
			restTime,
			timeRangeName,
			onShiftTime,
			offShiftTime,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?
		);
	`
	sqlQueryInstance.BanchStyle.UpdateSingle = `
		update banchStyle
		set
			icon=?,
			restTime=?,
			timeRangeName=?,
			onShiftTime=?,
			offShiftTime=?,
			lastModify=?
		where styleId=?;
	`;
	sqlQueryInstance.BanchStyle.SelectByCompanyCode = `
	select bs.* from banchStyle as bs
	left join companyBanch cb
		on cb.id=bs.banchId
	left join company c
		on cb.companyId=c.companyId
	where
		bs.banchId=?
	and
		c.companyCode=?;
	`
	sqlQueryInstance.BanchStyle.UpdateByCompanyCode = `
		update banchStyle bs
		left join companyBanch cb
			on cb.id=bs.banchId
		left join company c
			on cb.companyId=c.companyId
		set
			bs.icon=?,
			bs.restTime=?,
			bs.timeRangeName=?,
			bs.onShiftTime=?,
			bs.offShiftTime=?,
			bs.lastModify=?
		where bs.styleId=? and c.companyCode=?;
	`
	sqlQueryInstance.BanchStyle.DeleteByCompanyCode = `
		delete bs from banchStyle bs
		left join companyBanch cb
			on cb.id=bs.banchId
		left join company c
			on cb.companyId=c.companyId
		where bs.styleId=? and c.companyCode=?;
	`
	sqlQueryInstance.BanchStyle.SelectSingleByStyleId = `select * from banchStyle where styleId = ?;`;
	sqlQueryInstance.BanchStyle.SelectAll = `select * from banchStyle;`;
	sqlQueryInstance.BanchStyle.Delete = `delete from banchStyle where styleId=?;`;
	sqlQueryInstance.BanchStyle.SelectAllByBanchId = `select * from banchStyle where banchId = ?;`;
}