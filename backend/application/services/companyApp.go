package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type CompanyApp struct {
	companyRepo repository.CompanyRepository
}

var _ CompanyAppInterface = &CompanyApp{}

type CompanyAppInterface interface {
	UpdateCompany(*entities.Company) (*entities.Company, *map[string]string)
	GetCompany(*entities.Company) (*entities.Company, error)
}

func (c *CompanyApp) UpdateCompany(companyEntity *entities.Company) (*entities.Company, *map[string]string) {
	return c.companyRepo.UpdateCompany(companyEntity)
}

func (c *CompanyApp) GetCompany(companyEntity *entities.Company) (*entities.Company, error) {
	return c.companyRepo.GetCompany(companyEntity)
}
