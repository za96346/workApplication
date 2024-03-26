package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CompanyBanchRepo struct {
	db        *gorm.DB
	tableName string
}

func NewCompanyBanchRepository(db *gorm.DB) *CompanyBanchRepo {
	return &CompanyBanchRepo{db, "company_banch"}
}

var _ repository.CompanyBanchRepository = &CompanyBanchRepo{}

// 查詢 該公司 部門
func (r *CompanyBanchRepo) GetCompanyBanches(
	companyBanchEntity *entities.CompanyBanch,
	scopeBanch *[]int,
) (*[]entities.CompanyBanch, error) {

	var companyBanches []entities.CompanyBanch

	searchQuery := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", companyBanchEntity.CompanyId).
		Where("deleteFlag = ?", "N").
		Order("sort asc")

	// banch name
	if companyBanchEntity.BanchName != "" {
		searchQuery = searchQuery.Where("banchName like ?", "%"+companyBanchEntity.BanchName+"%")
	}

	// 部門範圍
	if scopeBanch != nil {
		searchQuery = searchQuery.Where("banchId in (?)", *scopeBanch)
	}

	err := searchQuery.Find(&companyBanches).Error
	return &companyBanches, err
}

// 查詢 選擇器 該公司 部門
func (r *CompanyBanchRepo) GetCompanyBanchesSelector(
	companyId int,
) (*[]entities.CompanyBanch, error) {

	// 獲取部門
	var companyBanches []entities.CompanyBanch
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", companyId).
		Find(&companyBanches).
		Error

	return &companyBanches, err
}

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

func (r *CompanyBanchRepo) GetBanchesId(banchEntity *entities.CompanyBanch) *[]int {
	var banchesIdArr []int

	r.db.
		Debug().
		Table(r.tableName).
		Select("banchId").
		Where("companyId = ?", banchEntity.CompanyId).
		Where("deleteFlag = ?", "N").
		Find(&entities.CompanyBanch{}).
		Pluck("banchId", &banchesIdArr)

	return &banchesIdArr
}

func (r *CompanyBanchRepo) GetBanchesIdByScopeBanch(banchEntity *entities.CompanyBanch, scopeBanch *[]int) *[]int {
	var banchesIdArr []int

	r.db.
		Debug().
		Table(r.tableName).
		Select("banchId").
		Where("companyId = ?", banchEntity.CompanyId).
		Where("deleteFlag = ?", "N").
		Where("banchId in (?)", *scopeBanch).
		Find(&entities.CompanyBanch{}).
		Pluck("banchId", &banchesIdArr)

	return &banchesIdArr
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

// 更新 該公司部門
func (r *CompanyBanchRepo) UpdateCompanyBanch(
	companyBanchEntity *entities.CompanyBanch,
) (*entities.CompanyBanch, error) {

	// 添加固定欄位
	now := time.Now()

	companyBanchEntity.LastModify = &now
	companyBanchEntity.DeleteTime = nil
	companyBanchEntity.DeleteFlag = "N"

	if r.IsBanchNameDuplicated(companyBanchEntity) {
		return nil, errors.New("部門名稱重複")
	}

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", (*companyBanchEntity).CompanyId).
		Where("banchId = ?", (*companyBanchEntity).BanchId).
		Updates(companyBanchEntity).
		Error

	return companyBanchEntity, err
}

// 新增 該公司部門
func (r *CompanyBanchRepo) SaveCompanyBanch(
	companyBanchEntity *entities.CompanyBanch,
) (*entities.CompanyBanch, error) {

	// 添加固定欄位
	now := time.Now()
	(*companyBanchEntity).BanchId = (*r).GetNewBanchID((*companyBanchEntity).CompanyId)

	(*companyBanchEntity).LastModify = &now
	(*companyBanchEntity).CreateTime = &now
	(*companyBanchEntity).DeleteTime = nil
	(*companyBanchEntity).DeleteFlag = "N"

	if r.IsBanchNameDuplicated(companyBanchEntity) {
		return nil, errors.New("部門名稱重複")
	}

	err := r.db.
		Debug().
		Table(r.tableName).
		Create(companyBanchEntity).
		Error

	return companyBanchEntity, err
}

// 刪除 該公司部門
func (r *CompanyBanchRepo) DeleteCompanyBanch(
	companyBanchEntity *entities.CompanyBanch,
) (*entities.CompanyBanch, error) {

	// 加入固定欄位
	now := time.Now()
	(*companyBanchEntity).DeleteFlag = "Y"
	(*companyBanchEntity).DeleteTime = &now
	(*companyBanchEntity).LastModify = &now

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", companyBanchEntity.CompanyId).
		Where("banchId = ?", companyBanchEntity.BanchId).
		Updates(companyBanchEntity).
		Error

	return companyBanchEntity, err
}
