package repository

import "backend/domain/entities"


type FunctionRoleBanchRelationRepository interface {
	GetFunctionRoleBanchRelations() (*[]entities.FuncRoleBanchRelation, *error)
}