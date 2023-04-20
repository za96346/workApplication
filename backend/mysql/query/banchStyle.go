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
	and bs.delFlag="N"
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
		update banchStyle bs
		left join companyBanch cb
			on cb.id=bs.banchId
		left join company c
			on cb.companyId=c.companyId
		set
			bs.delFlag="Y"
		where bs.styleId=? and c.companyCode=?;
	`
	sqlQueryInstance.BanchStyle.SelectSingleByStyleId = `
		select * from banchStyle
		where styleId = ?
		and delFlag="N";
	`;
	sqlQueryInstance.BanchStyle.SelectAll = `select * from banchStyle where delFlag="N";`;
	sqlQueryInstance.BanchStyle.Delete = `
		update banchStyle
		set
			bs.delFlag="Y"
		where styleId=?;
	`;
	sqlQueryInstance.BanchStyle.SelectAllByBanchId = `
		select * from banchStyle
		where banchId = ?
		and delFlag="N";
	`;
}