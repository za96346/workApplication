package persistence

import (
	appDtos "backend/application/dtos"
	domainDtos "backend/domain/dtos"
	"backend/domain/entities"
	"backend/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type PerformanceRepo struct {
	db        *gorm.DB
	tableName string
}

func NewPerformanceRepository(db *gorm.DB) *PerformanceRepo {
	return &PerformanceRepo{db, "performance"}
}

var _ repository.PerformanceRepository = &PerformanceRepo{}

func (r *PerformanceRepo) GetPerformances(
	performanceEntity *entities.Performance,
	queryParams *appDtos.PerformanceQueryParams,
	scopeBanch *[]int,
	scopeRole *[]int,
) (*[]domainDtos.PerformanceDetailDto, *error) {
	// 獲取資料
	var data []domainDtos.PerformanceDetailDto
	searchQuery := r.db.
		Debug().
		Table(r.tableName).
		Where("performance.companyId = ?", performanceEntity.CompanyId).
		Where("performance.banchId in (?)", *scopeBanch).
		Where("user.roleId in (?)", *scopeRole).
		Where("performance.deleteFlag = ?", "N").
		Joins(`
			left join user
			on user.userId = performance.userId
			and user.companyId = performance.companyId
		`).
		Joins(`
			left join company_banch
			on company_banch.companyId = performance.companyId
			and company_banch.banchId = performance.banchId
		`).
		Select(
			"performance.*",
			"user.userName as userName",
			"company_banch.banchName as banchName",
			"user.roleId",
		).
		Order("year desc, month desc, sort")

	// 使用者名稱
	if queryParams.UserName != nil {
		searchQuery.Where("userName like ?", "%"+*queryParams.UserName+"%")
	}

	// 日期塞選
	if queryParams.StartDate != nil {
		searchQuery.Where(
			`
				concat(
					performance.year,
					'-',
					if(
						performance.month < 10,
						concat('0', performance.month),
						performance.month
					)
				) >= ?
			`,
			*queryParams.StartDate,
		)
	}

	if queryParams.EndDate != nil {
		searchQuery.Where(
			`
				concat(
					performance.year,
					'-',
					if(
						performance.month < 10,
						concat('0', performance.month),
						performance.month
					)
				) <= ?
			`,
			*queryParams.EndDate,
		)
	}

	err := searchQuery.Find(&data).Error

	return &data, &err
}

func (r *PerformanceRepo) GetYearPerformances(
	performanceEntity *entities.Performance,
	queryParams *appDtos.PerformanceQueryParams,
	scopeBanch *[]int,
	scopeRole *[]int,
) (*[]entities.YearPerformance, *error) {
	var data []entities.YearPerformance

	searchQuery := r.db.
		Debug().
		Table(r.tableName).
		Where("performance.companyId = ?", performanceEntity.CompanyId).
		Where("performance.banchId in (?)", *scopeBanch).
		Where("user.roleId in (?)", *scopeRole).
		Where("performance.deleteFlag = ?", "N").
		Joins(`
			left join user
			on user.userId = performance.userId
			and user.companyId = performance.companyId
		`).
		Joins(`
			left join company_banch
			on company_banch.companyId = performance.companyId
			and company_banch.banchId = performance.banchId
		`).
		Group("performance.userId").
		Group("performance.year").
		Group("user.userName").
		Order("score desc").
		Select(
			"performance.year as year",
			"user.userName as userName",
			`
				round(
					(
						sum(performance.attitude)
						+ sum(performance.efficiency)
						+ sum(performance.professional)
					) / 36, 2
				) as score
			`,
		)

	// 使用者名稱
	if queryParams.UserName != nil {
		searchQuery.Where("user.userName like ?", "%"+*queryParams.UserName+"%")
	}

	// 年度塞選
	if queryParams.StartYear != nil {
		searchQuery.Where("performance.year >= ?", *queryParams.StartYear)
	}

	if queryParams.EndYear != nil {
		searchQuery.Where("performance.year <= ?", *queryParams.EndYear)
	}

	err := searchQuery.Find(&data).Error

	return &data, &err
}

func (r *PerformanceRepo) GetPerformance(performanceEntity *entities.Performance) (*entities.Performance, *error) {
	var performance entities.Performance
	searchQuery := r.db.
		Debug().
		Table(r.tableName)

	if &performanceEntity.PerformanceId != nil {
		searchQuery.Where("performanceId = ?", performanceEntity.PerformanceId)
	}

	if &performanceEntity.CompanyId != nil {
		searchQuery.Where("companyId = ?", performanceEntity.CompanyId)
	}

	err := searchQuery.First(&performance).Error

	return performanceEntity, &err
}

func (r *PerformanceRepo) SavePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error) {
	if r.IsYearMonthDuplicated(performanceEntity) {
		err := errors.New("年月份重複")
		return nil, &err
	}

	// 新增固定欄位
	now := time.Now()
	(*performanceEntity).PerformanceId = r.GetNewPerformanceID(performanceEntity.CompanyId)
	(*performanceEntity).DeleteFlag = "N"
	(*performanceEntity).DeleteTime = nil
	(*performanceEntity).CreateTime = &now
	(*performanceEntity).LastModify = &now

	err := r.db.
		Debug().
		Table(r.tableName).
		Create(performanceEntity).
		Error

	return performanceEntity, &err
}

func (r *PerformanceRepo) UpdatePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error) {
	if r.IsYearMonthDuplicated(performanceEntity) {
		err := errors.New("年月份重複")
		return nil, &err
	}

	// 新增固定欄位
	now := time.Now()
	(*performanceEntity).DeleteFlag = "N"
	(*performanceEntity).DeleteTime = nil
	(*performanceEntity).LastModify = &now

	err := r.db.
		Debug().
		Table(r.tableName).
		Updates(performanceEntity).
		Error

	return performanceEntity, &err
}

func (r *PerformanceRepo) DeletePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error) {
	// 新增固定欄位
	now := time.Now()
	(*performanceEntity).DeleteFlag = "Y"
	(*performanceEntity).DeleteTime = &now
	(*performanceEntity).LastModify = &now

	err := r.db.
		Debug().
		Table(r.tableName).
		Updates(performanceEntity).
		Error

	return performanceEntity, &err
}

// 拿取新的 performance id ( max count + 1 )
func (r *PerformanceRepo) GetNewPerformanceID(companyId int) int {
	var MaxCount int64

	r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", companyId).
		Select("max(performanceId)").
		Row().
		Scan(&MaxCount)

	return int(MaxCount) + 1
}

// 檢查績效年月是否重複
func (r *PerformanceRepo) IsYearMonthDuplicated(performanceEntity *entities.Performance) bool {
	var MaxCount int64

	r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", (*performanceEntity).CompanyId).
		Where("userId = ?", (*performanceEntity).UserId).
		Where("performanceId != ?", (*performanceEntity).PerformanceId).
		Where("Year = ?", (*performanceEntity).Year).
		Where("Month = ?", (*performanceEntity).Month).
		Where("deleteFlag = ?", "N").
		Count(&MaxCount)

	return MaxCount > 0
}
