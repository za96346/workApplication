package application

import (
	"backend/domain/aggregates"
	"backend/domain/entities"
	"backend/domain/repository"
	"encoding/json"
)

type RoleApp struct {
	roleRepo repository.RoleRepository
	roleStructRepo repository.RoleStructRepository
}

var _ RoleAppInterface = &RoleApp{}

type RoleAppInterface interface {
	GetRole(*entities.Role) (*entities.Role, *map[string](map[string]map[string]interface{}))
	GetRoles(*entities.Role) (*[]entities.Role, *map[string]string)
	GetRolesSelector(*entities.Role) (*[]entities.Role, *map[string]string)

	UpdateRole(*entities.Role, *map[string]map[string]map[string]interface{}) (*entities.Role, *map[string]string)

	SaveRole(*entities.Role, *map[string]map[string]map[string]interface{}) (*entities.Role, *map[string]string)
	DeleteRole(*entities.Role) (*entities.Role, *map[string]string)
}

func (r *RoleApp) GetRoles(roleEntity *entities.Role) (*[]entities.Role, *map[string]string) {
	return r.roleRepo.GetRoles(roleEntity)
}

func (r *RoleApp) GetRole(roleEntity *entities.Role) (*entities.Role, *map[string](map[string]map[string]interface{})) {
	rolePermissionMap := map[string](map[string]map[string]interface{}){}

	role, _ := r.roleRepo.GetRole(roleEntity)
	roleStructs, _ := r.roleStructRepo.GetRoleStructs(&entities.RoleStruct{
		CompanyId: roleEntity.CompanyId,
		RoleId: roleEntity.RoleId,
	})

	for _, v := range *roleStructs {
		if rolePermissionMap[v.FuncCode] == nil {
			rolePermissionMap[v.FuncCode] = make(map[string]map[string]interface{})
		}
		if rolePermissionMap[v.FuncCode][v.ItemCode] == nil {
			rolePermissionMap[v.FuncCode][v.ItemCode] = make(map[string]interface{})
		}

		// handle scope banch
		if v.ScopeBanch != "all" && v.ScopeBanch != "self" {
			var scopeBanch []int
			json.Unmarshal([]byte(v.ScopeBanch), &scopeBanch)
			rolePermissionMap[v.FuncCode][v.ItemCode]["scopeBanch"] = scopeBanch
		} else {
			rolePermissionMap[v.FuncCode][v.ItemCode]["scopeBanch"] = v.ScopeBanch
		}

		// handle scope role
		if v.ScopeRole != "all" && v.ScopeRole != "self" {
			var scopeRole []int
			json.Unmarshal([]byte(v.ScopeRole), &scopeRole)
			rolePermissionMap[v.FuncCode][v.ItemCode]["scopeRole"] = scopeRole
		} else {
			rolePermissionMap[v.FuncCode][v.ItemCode]["scopeRole"] = v.ScopeRole
		}
	}
	return role, &rolePermissionMap
}

func (r *RoleApp) GetRolesSelector(roleEntity *entities.Role) (*[]entities.Role, *map[string]string) {
	return r.roleRepo.GetRolesSelector(roleEntity)
}

func (r *RoleApp) UpdateRole(
	roleEntity *entities.Role,
	data *map[string]map[string]map[string]interface{},
) (*entities.Role, *map[string]string) {
	TX := r.roleRepo.Begin()
	if _, err := r.roleRepo.UpdateRole(roleEntity, TX); err != nil {
		TX.Rollback()
		return nil, err
	}

	roleAggregate := aggregates.NewRoleAggregate(
		roleEntity,
		&[]entities.RoleStruct{},
		r.roleStructRepo,
	)

	if err := roleAggregate.AddRoleStruct(data, TX,); err != nil {
		TX.Rollback()
		return nil, err
	}

	TX.Commit()
	return nil, nil
}

func (r *RoleApp) SaveRole(
	roleEntity *entities.Role,
	data *map[string]map[string]map[string]interface{},
) (*entities.Role, *map[string]string) {
	TX := r.roleRepo.Begin()
	if _, err := r.roleRepo.SaveRole(roleEntity, TX); err != nil {
		TX.Rollback()
		return nil, err
	}

	roleAggregate := aggregates.NewRoleAggregate(
		roleEntity,
		&[]entities.RoleStruct{},
		r.roleStructRepo,
	)

	if err := roleAggregate.AddRoleStruct(data, TX,); err != nil {
		TX.Rollback()
		return nil, err
	}

	TX.Commit()
	return nil, nil
}

func (r *RoleApp) DeleteRole(roleEntity *entities.Role) (*entities.Role, *map[string]string) {
	return r.roleRepo.DeleteRole(roleEntity)
}