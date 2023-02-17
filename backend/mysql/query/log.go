package query

func AddLog () {
	sqlQueryInstance.Log.SelectAll = `
		select * from log where createTime>? && createTime<?;
	`
	sqlQueryInstance.Log.InsertAll = `
		insert into log(
			userId,
			userName,
			companyId,
			companyCode,
			permession,
			routes,
			ip,
			params,
			msg,
			createTime,
			lastModify
		) values (
			?,?,?,?,?,?,?,?,?,?,?
		)
	;`;
}