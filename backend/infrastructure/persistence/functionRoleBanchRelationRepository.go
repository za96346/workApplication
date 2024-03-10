package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type FunctionRoleBanchRelationRepo struct {
	db        *gorm.DB
	tableName string
}

func NewFunctionRoleBanchRelationRepository(db *gorm.DB) *FunctionRoleBanchRelationRepo {
	return &FunctionRoleBanchRelationRepo{db, "func_role_banch_relation"}
}

var _ repository.FunctionRoleBanchRelationRepository = &FunctionRoleBanchRelationRepo{}

func (f *FunctionRoleBanchRelationRepo) GetFunctionRoleBanchRelations() (*[]entities.FuncRoleBanchRelation, error) {
	var funcRoleBanchRelation []entities.FuncRoleBanchRelation

	err := f.db.
		Debug().
		Table(f.tableName).
		Find(&funcRoleBanchRelation).
		Error

	return &funcRoleBanchRelation, err
}
