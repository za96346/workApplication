package persistence

import (
	"backend/domain/repository"
	"github.com/jinzhu/gorm"
)


type FunctionRoleBanchRelationRepo struct {
	db *gorm.DB
	tableName string
}

func NewFunctionRoleBanchRelationRepository(db *gorm.DB) *FunctionRoleBanchRelationRepo {
	return &FunctionRoleBanchRelationRepo{db, "func_role_banch_relation"}
}

var _ repository.FunctionRoleBanchRelationRepository = &FunctionRoleBanchRelationRepo{}