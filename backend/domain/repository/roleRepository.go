package repository

import (
	"backend/domain/entities"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRole(*entities.Role) (*entities.Role, *map[string]string)
	GetRoles(*entities.Role) (*[]entities.Role, *map[string]string)
	GetRolesSelector(*entities.Role) (*[]entities.Role, *map[string]string)
	UpdateRole(*entities.Role,*gorm.DB) (*entities.Role, *map[string]string)
	SaveRole(*entities.Role,*gorm.DB) (*entities.Role, *map[string]string)
	DeleteRole(*entities.Role) (*entities.Role, *map[string]string)
	
	GetNewRoleID(int) int
	GetRolesId(*entities.Role) *[]int
	GetRolesIdByScopeRole(*entities.Role,*[]int) *[]int
	IsRoleNameDuplicated(*entities.Role) bool
	

	Begin() *gorm.DB
}