package query

func AddPerformance(){
	sqlQueryInstance.Performance.SelectSingleByAdmin = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			performanceId=?
		and (u.companyCode=?
			or qu.companyCode=?);
	`
	sqlQueryInstance.Performance.SelectSingleByManager = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			performanceId=?
		and
			(u.companyCode=?
			or qu.companyCode=?)
		and
			(p.banchId=? or p.banchName=?);
	`
	sqlQueryInstance.Performance.SelectSingleByPerson = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			performanceId=?
		and
			u.userId=? or qu.userId=?;
	`
	sqlQueryInstance.Performance.SelectAllByAdmin = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			(u.companyCode=?
			or qu.companyCode=?)
			and 
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) >= ?
			and
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) <= ?
			and u.userName=if(?='' or ?=null, u.userName, ?)
			order by p.year asc, p.month asc;
	`;
	sqlQueryInstance.Performance.SelectAllByManager = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			(u.companyCode=?
				or qu.companyCode=?)
			and (p.banchId=?
				or p.banchName=?)
				and 
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) >= ?
			and
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) <= ?
			and u.userName=if(?='' or ?=null, u.userName, ?)
			order by p.year asc, p.month asc;
	`;
	sqlQueryInstance.Performance.SelectAllByPerson = `
		select
			p.*,
			ifnull(u.userName, ''),
			ifnull(c.companyId, -1)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			p.userId=? 
			and 
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) >= ?
			and
				concat(
					p.year,
					if(p.month < 10, concat('0', p.month), p.month)
				) <= ?
		order by p.year asc, p.month asc;
	`;
	sqlQueryInstance.Performance.UpdateByAdmin = `
		update performance p
		left join user u
			on u.userId=p.userId
		left join quitWorkUser qu
			on qu.userId=p.userId
		set
			banchId=?,
			goal=?,
			attitude=?,
			efficiency=?,
			professional=?,
			directions=?,
			beLate=?,
			dayOffNotOnRule=?,
			banchName=?,
			p.lastModify=?
		where
			p.performanceId=?
			and (qu.companyCode=?
			or u.companyCode=?);
	`
	sqlQueryInstance.Performance.UpdateByManager = `
		update performance p
		left join user u
			on u.userId=p.userId
		left join quitWorkUser qu
			on qu.userId=p.userId
		set
			banchId=?,
			goal=?,
			attitude=?,
			efficiency=?,
			professional=?,
			directions=?,
			beLate=?,
			dayOffNotOnRule=?,
			banchName=?,
			p.lastModify=?
		where
			p.performanceId=?
			and (qu.companyCode=?
				or u.companyCode=?)
			and (p.banchId=? or p.banchName=?);
	`;
	sqlQueryInstance.Performance.InsertAll = `
		insert into performance (
			userId,
			year,
			month,
			banchId,
			goal,
			attitude,
			efficiency,
			professional,
			directions,
			beLate,
			dayOffNotOnRule,
			banchName,
			createTime,
			lastModify
		) values (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		);
	`
	sqlQueryInstance.Performance.SelectYearPerformanceByAdmin = `
		select
			p.userId,
			p.year,
			ifnull(u.userName, ''),
			round((sum(p.attitude) + sum(p.professional) + sum(p.efficiency)) / 3.6, 2)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			(u.companyCode=?
			or qu.companyCode=?)
			and 
				p.year>=?
			and
				p.year<=?
			and u.userName=if(?='' or ?=null, u.userName, ?)
			group by p.userId, p.year;
	`
	sqlQueryInstance.Performance.SelectYearPerformanceByManage = `
		select
			p.userId,
			p.year,
			ifnull(u.userName, ''),
			round((sum(p.attitude) + sum(p.professional) + sum(p.efficiency)) / 3.6, 2)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			(u.companyCode=?
			or qu.companyCode=?)
			and
			(p.banchId=?
				or p.banchName=?)
			and 
				p.year>=?
			and
				p.year<=?
			and u.userName=if(?='' or ?=null, u.userName, ?)
			group by p.userId, p.year;
	`
	sqlQueryInstance.Performance.SelectYearPerformanceByPerson = `
		select
			p.userId,
			p.year,
			ifnull(u.userName, ''),
			round((sum(p.attitude) + sum(p.professional) + sum(p.efficiency)) / 3.6, 2)
		from performance as p
		left join user u
			on u.userId=p.userId
		left join companyBanch cb
			on cb.id=p.banchId
		left join company c
			on u.companyCode=c.companyCode
		left join quitWorkUser qu
			on qu.userId=p.userId
		where
			(u.companyCode=?
			or qu.companyCode=?)
			and
				p.userId=?
			and 
				p.year>=?
			and
				p.year<=?
			and u.userName=if(?='' or ?=null, u.userName, ?)
			group by p.userId, p.year;
	`
	sqlQueryInstance.Performance.DeleteByAdmin = `
		delete p from performance p
		where performanceId=?;
	`
	sqlQueryInstance.Performance.DeleteByManage = `
		delete p from performance p
		where performanceId=? && p.banchId=? && p.userId!=?;
	`
	sqlQueryInstance.Performance.UpdateByPerson = `
		update performance p
		set
			goal=?,
			p.lastModify=?
		where p.performanceId=? and p.userId=?;
	`;
}