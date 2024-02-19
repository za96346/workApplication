package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"

	"gorm.io/gorm"
)


type RoleStrcutRepo struct {
	db *gorm.DB
	tableName string
}

func NewRoleStructRepository(db *gorm.DB) *RoleStrcutRepo {
	return &RoleStrcutRepo{db, "role_struct"}
}

var _ repository.RoleStructRepository = &RoleStrcutRepo{}

func (r *RoleStrcutRepo) GetRoleStructs(roleStructEntity *entities.RoleStruct) (*[]entities.RoleStruct, *map[string]string) {
	var roleStructs []entities.RoleStruct

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleStructEntity.CompanyId).
		Where("roleId = ?", roleStructEntity.RoleId).
		Find(&roleStructs).
		Error

	return &roleStructs, persistenceErrorHandler(err)
}

func (r *RoleStrcutRepo) DeleteRoleStructs(roleStructEntity *entities.RoleStruct, Tx *gorm.DB) (*entities.RoleStruct, *map[string]string) {
	var roleStruct entities.RoleStruct

	err := Tx.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleStructEntity.CompanyId).
		Where("roleId = ?", roleStructEntity.RoleId).
		Delete(&roleStruct).
		Error

	return &roleStruct, persistenceErrorHandler(err)
}

func (r *RoleStrcutRepo) SaveRoleStruct(roleStructEntity *entities.RoleStruct, Tx *gorm.DB) (*entities.RoleStruct, *map[string]string) {
	var roleStruct entities.RoleStruct

	err := Tx.
		Debug().
		Table(r.tableName).
		Create(&roleStruct).
		Error

	return &roleStruct, persistenceErrorHandler(err)
}