package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type RoleApp struct {
	roleRepo repository.RoleRepository
}

var _ RoleAppInterface = &RoleApp{}

type RoleAppInterface interface {
	GetRole(*entities.Role) (*entities.Role, error)
	GetRoles(*entities.Role) (*[]entities.Role, error)
	GetRolesSelector(*entities.Role) (*[]entities.Role, error)

	UpdateRole(*entities.Role) (*entities.Role, *map[string]string)
	SaveRole(*entities.Role) (*entities.Role, *map[string]string)
	DeleteRole(*entities.Role) (*entities.Role, *map[string]string)
}