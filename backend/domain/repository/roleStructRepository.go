package repository

import (
	"backend/domain/entities"

	"gorm.io/gorm"
)


type RoleStructRepository interface {
	GetRoleStructs(*entities.RoleStruct) (*[]entities.RoleStruct, *map[string]string)
	DeleteRoleStructs(*entities.RoleStruct, *gorm.DB) (*entities.RoleStruct, *map[string]string)
	SaveRoleStruct(*entities.RoleStruct,*gorm.DB) (*entities.RoleStruct, *map[string]string)
}