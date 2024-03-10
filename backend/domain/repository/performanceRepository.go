package repository

import (
	"backend/domain/entities"
	domainDto "backend/domain/dtos"
	appDtos "backend/application/dtos"
)


type PerformanceRepository interface {
	GetPerformances(
		performanceEntity *entities.Performance,
		queryParams *appDtos.PerformanceQueryParams,
		scopeBanch *[]int,
		scopeRole *[]int,
	) (*[]domainDto.PerformanceDetailDto, error)

	GetYearPerformances(
		performanceEntity *entities.Performance,
		queryParams *appDtos.PerformanceQueryParams,
		scopeBanch *[]int,
		scopeRole *[]int,
	) (*[]entities.YearPerformance, error)
	GetPerformance(performanceEntity *entities.Performance) (*entities.Performance, error)

	SavePerformance(performanceEntity *entities.Performance) (*entities.Performance, error)
	UpdatePerformance(performanceEntity *entities.Performance) (*entities.Performance, error)
	DeletePerformance(performanceEntity *entities.Performance) (*entities.Performance, error)
	GetNewPerformanceID(int) int
	IsYearMonthDuplicated(*entities.Performance) bool
}