package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"

	"github.com/jinzhu/gorm"
)


type PerformanceRepo struct {
	db *gorm.DB
	tableName string
}

func NewPerformanceRepository(db *gorm.DB) *PerformanceRepo {
	return &PerformanceRepo{db, "performance"}
}

var _ repository.PerformanceRepository = &PerformanceRepo{}

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