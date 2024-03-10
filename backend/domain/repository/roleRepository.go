package repository

import (
	"backend/domain/entities"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRole(roleEntity *entities.Role) (*entities.Role, error)
	GetRoles(roleEntity *entities.Role) (*[]entities.Role, error)
	GetRolesSelector(roleEntity *entities.Role) (*[]entities.Role, error)
	UpdateRole(roleEntity *entities.Role, TX *gorm.DB) (*entities.Role, error)
	SaveRole(roleEntity *entities.Role, TX *gorm.DB) (*entities.Role, error)
	DeleteRole(roleEntity *entities.Role) (*entities.Role, error)
	
	GetNewRoleID(int) int
	GetRolesId(*entities.Role) *[]int
	GetRolesIdByScopeRole(*entities.Role,*[]int) *[]int
	IsRoleNameDuplicated(*entities.Role) bool
	

	Begin() *gorm.DB
}