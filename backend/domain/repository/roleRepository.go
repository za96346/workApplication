package repository

import "backend/domain/entities"

type RoleRepository interface {
	GetRole(*entities.Role) (*entities.Role, error)
	GetRoles() (*[]entities.Role, error)
	UpdateRole(*entities.Role) (*entities.Role, *map[string]string)
	SaveRole(*entities.Role) (*entities.Role, *map[string]string)
	
	GetNewRoleID(companyId int) int 
	IsRoleNameDuplicated(roleEntity *entities.Role) bool
}