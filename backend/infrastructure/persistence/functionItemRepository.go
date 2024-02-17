package persistence

import (
	"backend/domain/repository"
	"github.com/jinzhu/gorm"
)


type FunctionItemRepo struct {
	db *gorm.DB
	tableName string
}

func NewFunctionItemRepository(db *gorm.DB) *FunctionItemRepo {
	return &FunctionItemRepo{db, "function_item"}
}

var _ repository.FunctionItemRepository = &FunctionItemRepo{}