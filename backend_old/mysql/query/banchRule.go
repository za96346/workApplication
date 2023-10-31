package query

func AddBanchRuleQuery() {
	sqlQueryInstance.BanchRule.InsertAll = `
		insert into banchRule(
			banchId,
			maxPeople,
			minPeople,
			weekDay,
			weekType,
			onShiftTime,
			offShiftTime,
			createTime,
			lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?, ?
		);
	`
	sqlQueryInstance.BanchRule.UpdateSingle = `
		update banchRule
		set
			maxPeople=?,
			minPeople=?,
			weekDay=?,
			weekType=?,
			onShiftTime=?,
			offShiftTime=?,
			lastModify=?
		where ruleId=?;
	`
	sqlQueryInstance.BanchRule.SelectSingleByRuleId = `select * from banchRule where ruleId = ?;`;
	sqlQueryInstance.BanchRule.SelectAll = `select * from banchRule;`;
	sqlQueryInstance.BanchRule.Delete = `delete from banchRule where ruleId=?;`;
	sqlQueryInstance.BanchRule.SelectAllByBanchId = `select * from banchRule where banchId = ?;`;
}