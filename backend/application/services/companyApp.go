package application

import (
	"backend/domain/aggregates"
	"backend/domain/entities"
	"backend/domain/repository"
	"backend/enum"
	"backend/interfaces/method"
)

type CompanyApp struct {
	companyRepo repository.CompanyRepository
	companyBanchRepo repository.CompanyBanchRepository
	roleRepo repository.RoleRepository
}

var _ CompanyAppInterface = &CompanyApp{}

type CompanyAppInterface interface {
	UpdateCompany(companyEntity *entities.Company, sessionStruct *method.SessionStruct) (*entities.Company, *error)
	GetCompany(companyEntity *entities.Company, sessionStruct *method.SessionStruct) (*entities.Company, *error)
}

func (c *CompanyApp) UpdateCompany(
	companyEntity *entities.Company,
	sessionStruct *method.SessionStruct,
) (*entities.Company, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		c.roleRepo,
		c.companyBanchRepo,
		true,
		string(enum.CompanyData),
		string(enum.Edit),
	)
	
	if err != nil {
		return nil, err
	}

	(*companyEntity).CompanyId = authAggregate.User.CompanyId

	return c.companyRepo.UpdateCompany(companyEntity)
}

func (c *CompanyApp) GetCompany(
	companyEntity *entities.Company,
	sessionStruct *method.SessionStruct,
) (*entities.Company, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		c.roleRepo,
		c.companyBanchRepo,
		true,
		string(enum.CompanyData),
		string(enum.Inquire),
	)
	
	if err != nil {
		return nil, err
	}

	(*companyEntity).CompanyId = authAggregate.User.CompanyId

	return c.companyRepo.GetCompany(companyEntity)
}
