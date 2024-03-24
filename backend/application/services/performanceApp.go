package application

import (
	appDtos "backend/application/dtos"
	"backend/domain/aggregates"
	domainDtos "backend/domain/dtos"
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
	CopyPerformance(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error)
	ChangeBanch(*entities.Performance, *method.SessionStruct) (*entities.Performance, error)
}

func (p *PerformanceApp) GetPerformances(
	performanceEntity *entities.Performance,
	queryParams *appDtos.PerformanceQueryParams,
	sessionStruct *method.SessionStruct,
) (*[]domainDtos.PerformanceDetailDto, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
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
		authAggregate.GetScopeBanchWithCustomize(&queryParams.BanchId),
		authAggregate.GetScopeRolehWithCustomize(&queryParams.RoleId),
		&authAggregate.CurrentPermissionScopeUser,
	)
}

func (p *PerformanceApp) GetYearPerformances(
	performanceEntity *entities.Performance,
	queryParams *appDtos.PerformanceQueryParams,
	sessionStruct *method.SessionStruct,
) (*[]entities.YearPerformance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
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
		authAggregate.GetScopeBanchWithCustomize(&queryParams.BanchId),
		authAggregate.GetScopeRolehWithCustomize(&queryParams.RoleId),
		&authAggregate.CurrentPermissionScopeUser,
	)
}

func (p *PerformanceApp) SavePerformance(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
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

	if err := authAggregate.CheckScopeUserValidation(user.UserId); err != nil {
		return nil, err
	}

	// 加入使用者的欄位
	performanceEntity.BanchId = *user.BanchId

	return p.PerformanceRepo.SavePerformance(performanceEntity)
}

func (p *PerformanceApp) CopyPerformance(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		true,
		string(enum.Performance),
		string(enum.Copy),
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

	if err := authAggregate.CheckScopeUserValidation(user.UserId); err != nil {
		return nil, err
	}

	// 加入使用者的欄位
	performanceEntity.BanchId = *user.BanchId

	return p.PerformanceRepo.SavePerformance(performanceEntity)
}

func (p *PerformanceApp) UpdatePerformance(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
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

	if err := authAggregate.CheckScopeUserValidation(user.UserId); err != nil {
		return nil, err
	}

	// 加入使用者的欄位
	performanceEntity.BanchId = *user.BanchId

	return p.PerformanceRepo.UpdatePerformance(performanceEntity)
}

func (p *PerformanceApp) DeletePerformance(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		true,
		string(enum.Performance),
		string(enum.Delete),
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

	if err := authAggregate.CheckScopeUserValidation(user.UserId); err != nil {
		return nil, err
	}

	performance, err := p.PerformanceRepo.DeletePerformance(performanceEntity)

	if err != nil {
		return nil, err
	}

	return performance, nil
}

func (p *PerformanceApp) ChangeBanch(performanceEntity *entities.Performance, sessionStruct *method.SessionStruct) (*entities.Performance, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		true,
		string(enum.Performance),
		string(enum.Edit),
	)
	if err != nil { return nil, err }

	performanceEntity.CompanyId = authAggregate.User.CompanyId

	thisPerformance, err := p.PerformanceRepo.GetPerformance(performanceEntity)
	if err != nil { return nil, err }

	// 檢查 banch
	if err := authAggregate.CheckScopeBanchValidation(thisPerformance.BanchId); err != nil {
		return nil, err
	}

	// 把 user id 寫回
	performanceEntity.UserId = thisPerformance.UserId

	user, err := p.UserRepo.GetUser(&entities.User{
		CompanyId: authAggregate.User.CompanyId,
		UserId: thisPerformance.UserId,
	})
	if err != nil { return nil, err }

	// 檢查 role
	if err := authAggregate.CheckScopeRoleValidation(user.RoleId); err != nil {
		return nil, err
	}

	thisPerformance.BanchId = performanceEntity.BanchId
	return p.PerformanceRepo.UpdatePerformance(thisPerformance)
}