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
	) (*[]dtos.PerformanceDetailDto, *error)

	GetYearPerformances(
		performanceEntity *entities.Performance,
		userName string,
		startYear string,
		endYear string,
		scopeBanch *[]int,
		scopeRole *[]int,
	) (*[]entities.YearPerformance, *error)
	GetPerformance(performanceEntity *entities.Performance) (*entities.Performance, *error)

	SavePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error)
	UpdatePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error)
	DeletePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error)
	GetNewPerformanceID(int) int
	IsYearMonthDuplicated(*entities.Performance) bool
}