package application

import (
	"backend/domain/aggregates"
	domainDtos "backend/domain/dtos"
	appDtos "backend/application/dtos"
	"backend/domain/entities"
	"backend/domain/repository"
	"backend/enum"
	"backend/interfaces/method"
)

type PerformanceApp struct {
	PerformanceRepo repository.PerformanceRepository
	UserRepo repository.UserRepository
	CompanyBanchRepo repository.CompanyBanchRepository
	RoleRepo repository.RoleRepository
}

var _ PerformanceAppInterface = &PerformanceApp{}

type PerformanceAppInterface interface {
	GetPerformances(
		performanceEntity *entities.Performance,
		queryParams *appDtos.PerformanceQueryParams,
		sessionStruct *method.SessionStruct,
	) (*[]domainDtos.PerformanceDetailDto, error)
	GetYearPerformances(
		performanceEntity *entities.Performance,
		queryParams *appDtos.PerformanceQueryParams,
		sessionStruct *method.SessionStruct,
	) (*[]entities.YearPerformance, error)

	UpdatePerformance(*entities.Performance, *method.SessionStruct) (*entities.Performance, error)
	SavePerformance(*entities.Performance, *method.SessionStruct) (*entities.Performance, error)
	DeletePerformance(*entities.Performance, *method.SessionStruct) (*entities.Performance, error)

	ChangeBanch(*entities.Performance, *method.SessionStruct) (*entities.Performance, error)
}

func (p *PerformanceApp) GetPerformances(
	performanceEntity *entities.Performance,
	queryParams *appDtos.PerformanceQueryParams,
	sessionStruct *method.SessionStruct,
) (*[]domainDtos.PerformanceDetailDto, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		p.RoleRepo,
		p.CompanyBanchRepo,
		true,
		string(enum.Performance),
		string(enum.Inquire),
	)
	
	if err != nil {
		return nil, err
	}

	performanceEntity.CompanyId = authAggregate.User.CompanyId

	return p.PerformanceRepo.GetPerformances(
		performanceEntity,
		queryParams,
		&authAggregate.CurrentPermissionScopeBanch,
		&authAggregate.CurrentPermissionScopeRole,
	)
}

func (p *PerformanceApp) GetYearPerformances(
	performanceEntity *entities.Performance,
	queryParams *appDtos.PerformanceQueryParams,
	sessionStruct *method.SessionStruct,
) (*[]entities.YearPerformance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		p.RoleRepo,
		p.CompanyBanchRepo,
		true,
		string(enum.Performance),
		string(enum.Inquire),
	)
	
	if err != nil {
		return nil, err
	}

	performanceEntity.CompanyId = authAggregate.User.CompanyId

	return p.PerformanceRepo.GetYearPerformances(
		performanceEntity,
		queryParams,
		&authAggregate.CurrentPermissionScopeBanch,
		&authAggregate.CurrentPermissionScopeRole,
	)
}

func (p *PerformanceApp) SavePerformance(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		p.RoleRepo,
		p.CompanyBanchRepo,
		true,
		string(enum.Performance),
		string(enum.Add),
	)
	
	if err != nil {
		return nil, err
	}

	performanceEntity.CompanyId = authAggregate.User.CompanyId

	user, err := p.UserRepo.GetUser(&entities.User{
		CompanyId: performanceEntity.CompanyId,
		UserId: performanceEntity.UserId,
	})

	if user == nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeBanchValidation(*user.BanchId); err != nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeRoleValidation(user.RoleId); err != nil {
		return nil, err
	}

	return p.PerformanceRepo.SavePerformance(performanceEntity)
}

func (p *PerformanceApp) UpdatePerformance(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		p.RoleRepo,
		p.CompanyBanchRepo,
		true,
		string(enum.Performance),
		string(enum.Edit),
	)
	
	if err != nil {
		return nil, err
	}

	performanceEntity.CompanyId = authAggregate.User.CompanyId

	user, err := p.UserRepo.GetUser(&entities.User{
		CompanyId: authAggregate.User.CompanyId,
		UserId: performanceEntity.UserId,
	})

	if user == nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeBanchValidation(*user.BanchId); err != nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeRoleValidation(user.RoleId); err != nil {
		return nil, err
	}

	return p.PerformanceRepo.UpdatePerformance(performanceEntity)
}

func (p *PerformanceApp) DeletePerformance(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		p.RoleRepo,
		p.CompanyBanchRepo,
		true,
		string(enum.Performance),
		string(enum.Delete),
	)
	
	if err != nil {
		return nil, err
	}

	performanceEntity.CompanyId = authAggregate.User.CompanyId

	performance, err := p.PerformanceRepo.DeletePerformance(performanceEntity)

	if err != nil {
		return nil, err
	}

	return performance, nil
}

func (p *PerformanceApp) ChangeBanch(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		p.RoleRepo,
		p.CompanyBanchRepo,
		true,
		string(enum.Performance),
		string(enum.Edit),
	)
	
	if err != nil {
		return nil, err
	}

	performanceEntity.CompanyId = authAggregate.User.CompanyId

	user, err := p.UserRepo.GetUser(&entities.User{
		CompanyId: authAggregate.User.CompanyId,
		UserId: performanceEntity.UserId,
	})

	// 檢查 role
	if err := authAggregate.CheckScopeRoleValidation(user.RoleId); err != nil {
		return nil, err
	}

	thisPerformance, _ := p.PerformanceRepo.GetPerformance(performanceEntity)

	// 檢查 banch
	if err := authAggregate.CheckScopeBanchValidation(thisPerformance.BanchId); err != nil {
		return nil, err
	}

	thisPerformance.BanchId = performanceEntity.BanchId
	return p.PerformanceRepo.UpdatePerformance(thisPerformance)
}