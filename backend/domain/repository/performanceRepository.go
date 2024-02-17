package repository

import "backend/domain/entities"


type PerformanceRepository interface {
	GetNewPerformanceID(companyId int) int
	IsYearMonthDuplicated(performanceEntity *entities.Performance) bool
}