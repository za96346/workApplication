package repository

import "backend/domain/entities"


type OperationItemRepository interface {
	GetOperationItems() (*[]entities.OperationItem, *error)
}