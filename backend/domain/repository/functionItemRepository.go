package repository

import "backend/domain/entities"


type FunctionItemRepository interface {
	GetFunctionItemsByFuncCodes(funcCodes *[]string) (*[]entities.FunctionItem, error)
	GetFunctionItems() (*[]entities.FunctionItem, error)
}