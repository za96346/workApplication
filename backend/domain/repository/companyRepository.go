package repository

import "backend/domain/entities"

type CompanyRepository interface {
	GetCompany(companyEntity *entities.Company) (*entities.Company, *error)
	UpdateCompany(companyEntity *entities.Company) (*entities.Company, *error)
}