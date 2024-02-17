package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"

	"github.com/jinzhu/gorm"
)


type CompanyBanchRepo struct {
	db *gorm.DB
	tableName string
}

func NewCompanyBanchRepository(db *gorm.DB) *CompanyBanchRepo {
	return &CompanyBanchRepo{db, "company_banch"}
}

var _ repository.CompanyBanchRepository = &CompanyBanchRepo{}

// 拿取新的 banch id ( max count + 1 )
func (r *CompanyBanchRepo) GetNewBanchID(companyId int) int {
    var MaxCount int64
	r.db.
		Debug().
		Table(r.tableName).
        Where("companyId = ?", companyId).
        Select("max(banchId)").
        Row().
        Scan(&MaxCount)

    return int(MaxCount) + 1
}

// 查詢是否有重複banch name
func (r *CompanyBanchRepo) IsBanchNameDuplicated(companyBanchEntity *entities.CompanyBanch) bool {
    var MaxCount int64

	r.db.
		Debug().
		Table(r.tableName).
        Where("companyId = ?", (*companyBanchEntity).CompanyId).
        Where("banchName = ?", (*companyBanchEntity).BanchName).
        Where("banchId != ?", (*companyBanchEntity).BanchId).
        Where("deleteFlag = ?", "N").
        Count(&MaxCount)

    return int(MaxCount) > 0
}