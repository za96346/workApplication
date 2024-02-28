package persistence

import (
	"backend/domain/entities"
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

func (f *FunctionItemRepo) GetFunctionItemsByFuncCodes(funcCodes *[]string) (*[]entities.FunctionItem, *error) {
	var functionItem []entities.FunctionItem

	err := f.db.
		Debug().
		Table(f.tableName).
		Where(
			"funcCode in (?)",
			*funcCodes,
		).
		Find(functionItem).
		Error

	return &functionItem, &err
}

func (f *FunctionItemRepo) GetFunctionItems() (*[]entities.FunctionItem, *error) {
	var functionItem []entities.FunctionItem

	err := f.db.
		Debug().
		Table(f.tableName).
		Order("sort asc").
		Find(functionItem).
		Error

	return &functionItem, &err
}