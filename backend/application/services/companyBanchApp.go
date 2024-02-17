package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type CompanyBanchApp struct {
	companyBanchRepo repository.CompanyBanchRepository
}

var _ CompanyBanchAppInterface = &CompanyBanchApp{}

type CompanyBanchAppInterface interface {
	GetCompanyBanches(*entities.CompanyBanch) (*[]entities.CompanyBanch, error)
	GetCompanyBanchesSelector(*entities.CompanyBanch) (*[]entities.CompanyBanch, error)
	UpdateCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)
	SaveCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)
	DeleteCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)
}

func (c *CompanyBanchApp) GetCompanyBanches(*entities.CompanyBanch) (*[]entities.CompanyBanch, error) {
	return &[]entities.CompanyBanch{}, nil
}

func (c *CompanyBanchApp) GetCompanyBanchesSelector(*entities.CompanyBanch) (*[]entities.CompanyBanch, error) {
	return &[]entities.CompanyBanch{}, nil
}

func (c *CompanyBanchApp) UpdateCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string) {
	return &entities.CompanyBanch{}, &map[string]string{}
}

func (c *CompanyBanchApp) SaveCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string) {
	return &entities.CompanyBanch{}, &map[string]string{}
}

func (c *CompanyBanchApp) DeleteCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string) {
	return &entities.CompanyBanch{}, &map[string]string{}
}