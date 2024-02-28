package application

import (
	"backend/domain/dtos"
	"backend/domain/entities"
	"backend/domain/repository"
)

type PerformanceApp struct {
	performanceRepo repository.PerformanceRepository
	userRepo repository.UserRepository
}

var _ PerformanceAppInterface = &PerformanceApp{}

type PerformanceAppInterface interface {
	GetPerformances(*entities.Performance) (*[]dtos.PerformanceDetailDto, *error)
	GetYearPerformances(*entities.Performance) (*[]entities.YearPerformance, *error)

	UpdatePerformance(*entities.Performance) (*entities.Performance, *error)
	SavePerformance(*entities.Performance) (*entities.Performance, *error)
	DeletePerformance(*entities.Performance) (*entities.Performance, *error)

	ChangeBanch(*entities.Performance) (*entities.Performance, *error)
}

func (p *PerformanceApp) GetPerformances(performanceEntity *entities.Performance) (*[]dtos.PerformanceDetailDto, *error) {
	v := []int{}
	return p.performanceRepo.GetPerformances(
		performanceEntity,
		"",
		"",
		"",
		&v,
		&v,
	)
}

func (p *PerformanceApp) GetYearPerformances(performanceEntity *entities.Performance) (*[]entities.YearPerformance, *error) {
	v := []int{}
	return p.performanceRepo.GetYearPerformances(
		performanceEntity,
		"",
		"",
		"",
		&v,
		&v,
	)
}

func (p *PerformanceApp) SavePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error) {
	user, err := p.userRepo.GetUser(&entities.User{
		CompanyId: performanceEntity.CompanyId,
		UserId: performanceEntity.UserId,
	})
	if user == nil {
		return nil, err
	}

	return p.performanceRepo.SavePerformance(performanceEntity)
}

func (p *PerformanceApp) UpdatePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error) {
	user, err := p.userRepo.GetUser(&entities.User{
		CompanyId: performanceEntity.CompanyId,
		UserId: performanceEntity.UserId,
	})
	if user == nil {
		return nil, err
	}

	return p.performanceRepo.UpdatePerformance(performanceEntity)
}

func (p *PerformanceApp) DeletePerformance(performanceEntity *entities.Performance) (*entities.Performance, *error) {
	return p.performanceRepo.DeletePerformance(performanceEntity)
}

func (p *PerformanceApp) ChangeBanch(performanceEntity *entities.Performance) (*entities.Performance, *error) {
	thisPerformance, _ := p.performanceRepo.GetPerformance(performanceEntity)
	thisPerformance.BanchId = performanceEntity.BanchId
	return p.performanceRepo.UpdatePerformance(thisPerformance)
}