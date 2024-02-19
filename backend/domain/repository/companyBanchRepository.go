package repository

import "backend/domain/entities"


type CompanyBanchRepository interface {
	GetCompanyBanches(int, *[]int, *string) (*[]entities.CompanyBanch, *map[string]string)

	GetCompanyBanchesSelector(int) (*[]entities.CompanyBanch, *map[string]string)

	UpdateCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)

	SaveCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)

	DeleteCompanyBanch(*entities.CompanyBanch) (*entities.CompanyBanch, *map[string]string)

	GetNewBanchID(int) int
	GetBanchesId(*entities.CompanyBanch) *[]int
	GetBanchesIdByScopeBanch(*entities.CompanyBanch, *[]int) *[]int
	IsBanchNameDuplicated(*entities.CompanyBanch) bool
}