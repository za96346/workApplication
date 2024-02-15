package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type companyApp struct {
	companyRepo repository.CompanyRepository
}

var _ CompanyAppInterface = &companyApp{}

type CompanyAppInterface interface {
	UpdateCompany(*entities.Company) (*entities.Company, *map[string]string)
	GetCompany(*entities.Company) (*entities.Company, error)
}

func (c *companyApp) UpdateCompany(companyEntity *entities.Company) (*entities.Company, *map[string]string) {
	return c.companyRepo.UpdateCompany(companyEntity)
}

func (c *companyApp) GetCompany(companyEntity *entities.Company) (*entities.Company, error) {
	return c.companyRepo.GetCompany(companyEntity)
}
