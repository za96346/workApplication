package persistence

import (
	"backend/domain/repository"

	"gorm.io/gorm"
)

type LogRepo struct {
	db        *gorm.DB
	tableName string
}

func NewLogRepository(db *gorm.DB) *LogRepo {
	return &LogRepo{db, "log"}
}

var _ repository.LogRepository = &LogRepo{}

// 拿取新的 log id ( max count + 1 )
func (r *LogRepo) GetNewLogId(companyId int) int {
	var MaxCount int64
	r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", companyId).
		Select("max(logId)").
		Row().
		Scan(&MaxCount)

	return int(MaxCount) + 1
}
