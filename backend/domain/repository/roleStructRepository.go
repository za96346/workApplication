package repository

import (
	"backend/domain/entities"

	"gorm.io/gorm"
)


type RoleStructRepository interface {
	GetRoleStructs(roleStructEntity *entities.RoleStruct) (*[]entities.RoleStruct, error)
	GetRoleStructsFuncCode(roleStructEntity *entities.RoleStruct) (*[]string, error)
	DeleteRoleStructs(roleStructEntity *entities.RoleStruct, Tx *gorm.DB) (*entities.RoleStruct, error)
	SaveRoleStruct(roleStructEntity *entities.RoleStruct, Tx *gorm.DB) (*entities.RoleStruct, error)
}