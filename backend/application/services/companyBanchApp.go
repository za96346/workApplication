package application

import (
	"backend/domain/aggregates"
	"backend/domain/entities"
	"backend/domain/repository"
	"backend/enum"
	"backend/interfaces/method"
)

type CompanyBanchApp struct {
	CompanyBanchRepo repository.CompanyBanchRepository
	RoleRepo repository.RoleRepository
}

var _ CompanyBanchAppInterface = &CompanyBanchApp{}

type CompanyBanchAppInterface interface {
	GetCompanyBanches(
		companyBanchEntity *entities.CompanyBanch,
		sessionStruct *method.SessionStruct,
	) (*[]entities.CompanyBanch, error)
	GetCompanyBanchesSelector(sessionStruct *method.SessionStruct) (*[]entities.CompanyBanch, error)
	UpdateCompanyBanch(*entities.CompanyBanch, *method.SessionStruct) (*entities.CompanyBanch, error)
	SaveCompanyBanch(*entities.CompanyBanch, *method.SessionStruct) (*entities.CompanyBanch, error)
	DeleteCompanyBanch(*entities.CompanyBanch, *method.SessionStruct) (*entities.CompanyBanch, error)
}

func (c *CompanyBanchApp) GetCompanyBanches(
	companyBanchEntity *entities.CompanyBanch,
	sessionStruct *method.SessionStruct,
) (*[]entities.CompanyBanch, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		c.RoleRepo,
		c.CompanyBanchRepo,
		true,
		string(enum.BanchManage),
		string(enum.Inquire),
	)
	
	if err != nil {
		return nil, err
	}

	return c.CompanyBanchRepo.GetCompanyBanches(
		&entities.CompanyBanch{
			CompanyId: authAggregate.User.CompanyId,
			BanchName: companyBanchEntity.BanchName,
		},
		&authAggregate.CurrentPermissionScopeBanch,
	)
}

func (c *CompanyBanchApp) GetCompanyBanchesSelector(sessionStruct *method.SessionStruct) (*[]entities.CompanyBanch, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		c.RoleRepo,
		c.CompanyBanchRepo,
		false,
		"",
		"",
	)
	
	if err != nil {
		return nil, err
	}

	return c.CompanyBanchRepo.GetCompanyBanchesSelector(authAggregate.User.CompanyId)
}

func (c *CompanyBanchApp) UpdateCompanyBanch(companyBanchEntity *entities.CompanyBanch, sessionStruct *method.SessionStruct) (*entities.CompanyBanch, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		c.RoleRepo,
		c.CompanyBanchRepo,
		true,
		string(enum.BanchManage),
		string(enum.Edit),
	)
	
	if err != nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeBanchValidation((*companyBanchEntity).BanchId); err != nil {
		return nil, err
	}

	companyBanchEntity.CompanyId = authAggregate.User.CompanyId

	return c.CompanyBanchRepo.UpdateCompanyBanch(companyBanchEntity)
}

func (c *CompanyBanchApp) SaveCompanyBanch(companyBanchEntity *entities.CompanyBanch, sessionStruct *method.SessionStruct) (*entities.CompanyBanch, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		c.RoleRepo,
		c.CompanyBanchRepo,
		true,
		string(enum.BanchManage),
		string(enum.Add),
	)
	
	if err != nil {
		return nil, err
	}

	companyBanchEntity.CompanyId = authAggregate.User.CompanyId

	return c.CompanyBanchRepo.SaveCompanyBanch(companyBanchEntity)
}

func (c *CompanyBanchApp) DeleteCompanyBanch(companyBanchEntity *entities.CompanyBanch, sessionStruct *method.SessionStruct) (*entities.CompanyBanch, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		c.RoleRepo,
		c.CompanyBanchRepo,
		true,
		string(enum.BanchManage),
		string(enum.Delete),
	)
	
	if err != nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeBanchValidation((*companyBanchEntity).BanchId); err != nil {
		return nil, err
	}

	companyBanchEntity.CompanyId = authAggregate.User.CompanyId

	return c.CompanyBanchRepo.DeleteCompanyBanch(companyBanchEntity)
}