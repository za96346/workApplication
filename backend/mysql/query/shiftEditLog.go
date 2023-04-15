package query

func AddShiftEditLogQuery() {
	sqlQueryInstance.ShiftEditLog.InsertAll = `
	insert into shiftEditLog(
		year,
		month,
		banchId,
		msg
		) values(
		?, ?, ?, ?
	);`;
	sqlQueryInstance.ShiftEditLog.SelectByBanchId = `
		select * from shiftEditLog
		where
			banchId=?
		and
			year=?
		and
			month=?
		order by lastModify desc;
	`
}