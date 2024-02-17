package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type PerformanceApp struct {
	performanceRepo repository.PerformanceRepository
}

var _ PerformanceAppInterface = &PerformanceApp{}

type PerformanceAppInterface interface {
	GetPerformances(*entities.Performance) (*[]entities.Performance, error)
	GetYearPerformances(*entities.Performance) (*[]entities.Performance, error)

	UpdatePerformance(*entities.Performance) (*entities.Performance, *map[string]string)
	SavePerformance(*entities.Performance) (*entities.Performance, *map[string]string)
	DeletePerformance(*entities.Performance) (*entities.Performance, *map[string]string)

	ChangeBanch(*entities.Performance) (*entities.Performance, *map[string]string)
}