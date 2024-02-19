package repository

import "backend/domain/entities"


type PerformanceRepository interface {
	GetNewPerformanceID(int) int
	IsYearMonthDuplicated(*entities.Performance) bool
}