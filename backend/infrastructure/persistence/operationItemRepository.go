package persistence

import (
	"backend/domain/entities"
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

func (o *OperationItemRepo) GetOperationItems() (*[]entities.OperationItem, *error) {
	var operationItems []entities.OperationItem

	err := o.db.
		Debug().
		Table(o.tableName).
		Order("sort asc").
		Find(&operationItems).
		Error

	return &operationItems, &err
}