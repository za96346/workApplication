package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"time"

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

func (r *RoleStrcutRepo) GetRoleStructs(roleStructEntity *entities.RoleStruct) (*[]entities.RoleStruct, *error) {
	var roleStructs []entities.RoleStruct

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleStructEntity.CompanyId).
		Where("roleId = ?", roleStructEntity.RoleId).
		Find(&roleStructs).
		Error

	return &roleStructs, &err
}

func (r *RoleStrcutRepo) GetRoleStructsFuncCode(roleStructEntity *entities.RoleStruct) (*[]string, *error) {
	var roleStructs []string

	err := r.db.
		Debug().
		Table(r.tableName).
		Distinct().
		Select("funcCode").
		Where("companyId = ?", roleStructEntity.CompanyId).
		Where("roleId = ?", roleStructEntity.RoleId).
		Find(&roleStructs).
		Error

	return &roleStructs, &err
}

func (r *RoleStrcutRepo) DeleteRoleStructs(roleStructEntity *entities.RoleStruct, Tx *gorm.DB) (*entities.RoleStruct, *error) {
	var roleStruct entities.RoleStruct

	err := Tx.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleStructEntity.CompanyId).
		Where("roleId = ?", roleStructEntity.RoleId).
		Delete(&roleStruct).
		Error

	return &roleStruct, &err
}

func (r *RoleStrcutRepo) SaveRoleStruct(roleStructEntity *entities.RoleStruct, Tx *gorm.DB) (*entities.RoleStruct, *error) {
	now := time.Now()
	(*roleStructEntity).CreateTime = &now
	(*roleStructEntity).LastModify = &now

	err := Tx.
		Debug().
		Table(r.tableName).
		Create(&roleStructEntity).
		Error

	return roleStructEntity, &err
}