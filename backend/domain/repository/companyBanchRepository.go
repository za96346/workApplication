package repository

import "backend/domain/entities"


type CompanyBanchRepository interface {
	GetCompanyBanches(
		companyBanchEntity *entities.CompanyBanch,
		scopeBanch *[]int,
	) (*[]entities.CompanyBanch, error)

	GetCompanyBanchesSelector(
		companyId int,
	) (*[]entities.CompanyBanch, error)

	UpdateCompanyBanch(
		companyBanchEntity *entities.CompanyBanch,
	) (*entities.CompanyBanch, error)

	SaveCompanyBanch(
		companyBanchEntity *entities.CompanyBanch,
	) (*entities.CompanyBanch, error)

	DeleteCompanyBanch(
		companyBanchEntity *entities.CompanyBanch,
	) (*entities.CompanyBanch, error)

	GetNewBanchID(int) int
	GetBanchesId(*entities.CompanyBanch) *[]int
	GetBanchesIdByScopeBanch(*entities.CompanyBanch, *[]int) *[]int
	IsBanchNameDuplicated(*entities.CompanyBanch) bool
}