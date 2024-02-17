package persistence

import (
	"backend/domain/repository"
	"github.com/jinzhu/gorm"
)


type RoleStrcutRepo struct {
	db *gorm.DB
	tableName string
}

func NewRoleStructRepository(db *gorm.DB) *RoleStrcutRepo {
	return &RoleStrcutRepo{db, "role_struct"}
}

var _ repository.RoleStructRepository = &RoleStrcutRepo{}