package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"time"

	"github.com/jinzhu/gorm"
)


type CompanyRepo struct {
	db *gorm.DB
	tableName string
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepo {
	return &CompanyRepo{db, "company"}
}

var _ repository.CompanyRepository = &CompanyRepo{}

func (r *CompanyRepo) GetCompany(companyEntity *entities.Company) (*entities.Company, *map[string]string) {
	var company entities.Company
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("id = ?", (*companyEntity).CompanyId).
		First(&company).
		Error

	return &company, persistenceErrorHandler(err)
}


func (r *CompanyRepo) UpdateCompany(companyEntity *entities.Company) (*entities.Company, *map[string]string) {
	now := time.Now()
	(*companyEntity).LastModify = &now

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", (*companyEntity).CompanyId).
		Updates(&companyEntity).
		Error

	return companyEntity, persistenceErrorHandler(err)
}
