package application

import (
	"backend/domain/aggregates"
	"backend/domain/entities"
	"backend/domain/repository"
	"backend/enum"
	"backend/interfaces/method"
	"encoding/json"
)

type RoleApp struct {
	RoleRepo repository.RoleRepository
	RoleStructRepo repository.RoleStructRepository
	CompanyBanchRepo repository.CompanyBanchRepository
}

var _ RoleAppInterface = &RoleApp{}

type RoleAppInterface interface {
	GetRole(
		*entities.Role,
		*method.SessionStruct,
	) (
		*entities.Role,
		*map[string](map[string]map[string]interface{}),
		error,
	)
	GetRoles(*entities.Role, *method.SessionStruct) (*[]entities.Role, error)
	GetRolesSelector(*entities.Role, *method.SessionStruct) (*[]entities.Role, error)

	UpdateRole(
		*entities.Role,
		*map[string]map[string]map[string]interface{},
		*method.SessionStruct,
	) (*entities.Role, error)

	SaveRole(
		*entities.Role,
		*map[string]map[string]map[string]interface{},
		*method.SessionStruct,
	) (*entities.Role, error)
	DeleteRole(*entities.Role, *method.SessionStruct) (*entities.Role, error)
}

func (r *RoleApp) GetRoles(roleEntity *entities.Role, sessionStruct *method.SessionStruct) (*[]entities.Role, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		true,
		string(enum.RoleManage),
		string(enum.Inquire),
	)
	
	if err != nil {
		return nil, err
	}

	roleEntity.CompanyId = authAggregate.User.CompanyId

	return r.RoleRepo.GetRoles(roleEntity)
}

func (r *RoleApp) GetRole(
	roleEntity *entities.Role,
	sessionStruct *method.SessionStruct,
) (
	*entities.Role,
	*map[string](map[string]map[string]interface{}),
	error,
) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		true,
		string(enum.RoleManage),
		string(enum.Inquire),
	)
	
	if err != nil {
		return nil, nil, err
	}

	roleEntity.CompanyId = authAggregate.User.CompanyId

	rolePermissionMap := map[string](map[string]map[string]interface{}){}

	role, _ := r.RoleRepo.GetRole(roleEntity)
	roleStructs, _ := r.RoleStructRepo.GetRoleStructs(&entities.RoleStruct{
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
	return role, &rolePermissionMap, nil
}

func (r *RoleApp) GetRolesSelector(roleEntity *entities.Role, sessionStruct *method.SessionStruct) (*[]entities.Role, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		false,
		string(enum.RoleManage),
		string(enum.Inquire),
	)
	
	if err != nil {
		return nil, err
	}

	roleEntity.CompanyId = authAggregate.User.CompanyId

	return r.RoleRepo.GetRolesSelector(roleEntity)
}

func (r *RoleApp) UpdateRole(
	roleEntity *entities.Role,
	data *map[string]map[string]map[string]interface{},
	sessionStruct *method.SessionStruct,
) (*entities.Role, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		true,
		string(enum.RoleManage),
		string(enum.Edit),
	)
	
	if err != nil {
		return nil, err
	}

	roleEntity.CompanyId = authAggregate.User.CompanyId

	TX := r.RoleRepo.Begin()
	if _, err := r.RoleRepo.UpdateRole(roleEntity, TX); err != nil {
		TX.Rollback()
		return nil, err
	}

	roleAggregate := aggregates.NewRoleAggregate(
		roleEntity,
		&[]entities.RoleStruct{},
		r.RoleStructRepo,
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
	sessionStruct *method.SessionStruct,
) (*entities.Role, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		true,
		string(enum.RoleManage),
		string(enum.Add),
	)
	
	if err != nil {
		return nil, err
	}

	roleEntity.CompanyId = authAggregate.User.CompanyId

	TX := r.RoleRepo.Begin()
	if _, err := r.RoleRepo.SaveRole(roleEntity, TX); err != nil {
		TX.Rollback()
		return nil, err
	}

	roleAggregate := aggregates.NewRoleAggregate(
		roleEntity,
		&[]entities.RoleStruct{},
		r.RoleStructRepo,
	)

	if err := roleAggregate.AddRoleStruct(data, TX,); err != nil {
		TX.Rollback()
		return nil, err
	}

	TX.Commit()
	return nil, nil
}

func (r *RoleApp) DeleteRole(roleEntity *entities.Role, sessionStruct *method.SessionStruct) (*entities.Role, error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		true,
		string(enum.RoleManage),
		string(enum.Delete),
	)
	
	if err != nil {
		return nil, err
	}

	roleEntity.CompanyId = authAggregate.User.CompanyId

	return r.RoleRepo.DeleteRole(roleEntity)
}