package query

func AddWeekendSetting () {
	sqlQueryInstance.WeekendSetting.InsertAll = `
		insert into weekendSetting(
			companyId,
			date,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?
		);
	`
	sqlQueryInstance.WeekendSetting.UpdateSingle = `
		update weekendSetting
		set
			date=?,
			lastModify=?
		where weekendId=?;
	`
	sqlQueryInstance.WeekendSetting.SelectAll = `select * from weekendSetting;`;
	sqlQueryInstance.WeekendSetting.SelectSingleByWeekendId = `select * from weekendSetting where weekendId = ?;`
	sqlQueryInstance.WeekendSetting.SelectAllByCompanyId = `select * from weekendSetting where companyId = ?;`
	sqlQueryInstance.WeekendSetting.Delete = `delete from weekendSetting where weekendId = ?;`

}