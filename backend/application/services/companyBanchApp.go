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
	GetCompanyBanches(
		companyId int,
		scopeBanch *[]int,
		banchName *string,
	) (*[]entities.CompanyBanch, *map[string]string)
	GetCompanyBanchesSelector(int) (*[]entities.CompanyBanch, *map[string]string)
	UpdateCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)
	SaveCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)
	DeleteCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)
}

func (c *CompanyBanchApp) GetCompanyBanches(
	companyId int,
	scopeBanch *[]int,
	banchName *string,
) (*[]entities.CompanyBanch, *map[string]string) {
	return c.companyBanchRepo.GetCompanyBanches(
		companyId,
		scopeBanch,
		banchName,
	)
}

func (c *CompanyBanchApp) GetCompanyBanchesSelector(companyId int) (*[]entities.CompanyBanch, *map[string]string) {
	return c.companyBanchRepo.GetCompanyBanchesSelector(companyId)
}

func (c *CompanyBanchApp) UpdateCompanyBanch(companyBanchEntity *entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string) {
	return c.companyBanchRepo.UpdateCompanyBanch(companyBanchEntity)
}

func (c *CompanyBanchApp) SaveCompanyBanch(companyBanchEntity *entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string) {
	return c.companyBanchRepo.SaveCompanyBanch(companyBanchEntity)
}

func (c *CompanyBanchApp) DeleteCompanyBanch(companyBanchEntity *entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string) {
	return c.companyBanchRepo.DeleteCompanyBanch(companyBanchEntity)
}