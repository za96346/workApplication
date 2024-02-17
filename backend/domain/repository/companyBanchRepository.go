package repository

import "backend/domain/entities"


type CompanyBanchRepository interface {
	GetNewBanchID(companyId int) int
	IsBanchNameDuplicated(companyBanchEntity *entities.CompanyBanch) bool
}