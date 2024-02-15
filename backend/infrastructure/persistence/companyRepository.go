package persistence

import (
	"errors"
	"backend/domain/entities"
	"backend/domain/repository"
	"github.com/jinzhu/gorm"
	"strings"
)


type CompanyRepo struct {
	db *gorm.DB
	tableName string
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepo {
	return &CompanyRepo{db, "company"}
}

var _ repository.CompanyRepository = &CompanyRepo{}

func (r *CompanyRepo) GetCompany(companyEntity *entities.Company) (*entities.Company, error) {
	var company entities.Company
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("id = ?", (*companyEntity).CompanyId).
		Take(&companyEntity).
		Error

	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("company not found")
	}
	return &company, nil
}


func (r *CompanyRepo) UpdateCompany(companyEntity *entities.Company) (*entities.Company, *map[string]string) {
	dbErr := map[string]string{}
	err := r.db.
		Debug().
		Table(r.tableName).
		Save(&companyEntity).
		Error

	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, &dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, &dbErr
	}
	return companyEntity, nil
}
