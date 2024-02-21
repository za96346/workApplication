package repository

import (
	"backend/domain/entities"
	"backend/domain/dtos"
)


type PerformanceRepository interface {
	GetPerformances(
		performanceEntity *entities.Performance,
		userName string,
		startDate string,
		endDate string,
		scopeBanch *[]int,
		scopeRole *[]int,
	) (*[]dtos.PerformanceDetailDto, *map[string]string)

	GetYearPerformances(
		performanceEntity *entities.Performance,
		userName string,
		startYear string,
		endYear string,
		scopeBanch *[]int,
		scopeRole *[]int,
	) (*[]entities.YearPerformance, *map[string]string)
	GetPerformance(*entities.Performance) (*entities.Performance, *map[string]string)

	SavePerformance(*entities.Performance) (*entities.Performance, *map[string]string)
	UpdatePerformance(*entities.Performance) (*entities.Performance, *map[string]string)
	DeletePerformance(*entities.Performance) (*entities.Performance, *map[string]string)
	GetNewPerformanceID(int) int
	IsYearMonthDuplicated(*entities.Performance) bool
}