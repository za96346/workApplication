package persistence

import (
	"backend/domain/repository"
	"github.com/jinzhu/gorm"
)


type OperationItemRepo struct {
	db *gorm.DB
	tableName string
}

func NewOperationItemRepository(db *gorm.DB) *OperationItemRepo {
	return &OperationItemRepo{db, "operation_item"}
}

var _ repository.OperationItemRepository = &OperationItemRepo{}
