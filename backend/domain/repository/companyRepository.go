package repository

import "backend/domain/entities"

type CompanyRepository interface {
	GetCompany(*entities.Company) (*entities.Company, *map[string]string)
	UpdateCompany(*entities.Company) (*entities.Company, *map[string]string)
}