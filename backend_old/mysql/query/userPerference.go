package query

func AddUserPreferenceQuery() {
	sqlQueryInstance.UserPreference.InsertAll = `
	insert into userPreference(
		userId,
		style,
		fontSize,
		selfPhoto,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?
	);`;
	sqlQueryInstance.UserPreference.UpdateSingle = `
	update userPreference
	set
		style=?,
		fontSize=?,
		selfPhoto=?,
		lastModify=?
	where userId=?;
	`;
	sqlQueryInstance.UserPreference.SelectAll = `select * from userPreference;`;
	sqlQueryInstance.UserPreference.Delete = `delete from userPreference where userId = ?;`;
	sqlQueryInstance.UserPreference.SelectSingleByUserId = `select * from userPreference where userId = ?;`;
}