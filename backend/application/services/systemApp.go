package application

import (
	"backend/domain/aggregates"
	"backend/domain/entities"
	"backend/domain/repository"
	"backend/interfaces/method"
	"encoding/json"
	"strings"
)

type SystemApp struct {
	RoleRepo repository.RoleRepository
	RoleStructRepo repository.RoleStructRepository
	FunctionItemRepo repository.FunctionItemRepository
	FunctionRoleBanchRelationRepo repository.FunctionRoleBanchRelationRepository
	OperationItemRepo repository.OperationItemRepository
	CompanyBanchRepo repository.CompanyBanchRepository
	UserRepo repository.UserRepository
}

var _ SystemAppInterface = &SystemApp{}

type SystemAppInterface interface {
	GetAuth(sessionStruct *method.SessionStruct) (
		*[]entities.FunctionItem,
		*map[string](map[string](map[string]interface{})),
		*error,
	)
	GetFunc() (
		*[]entities.FunctionItem,
		*[]entities.OperationItem,
		*map[string]map[string]map[string][]string,
	)
	GetRoleBanchList(
		sessionStruct *method.SessionStruct,
	) (
		*map[string]map[string]interface{},
		*map[string]map[string]interface{},
		*map[string]map[string]interface{},
		*[]entities.CompanyBanch,
		*[]entities.Role,
		*[]entities.User,
		*error,
	)
}

// 獲取權限
func (s *SystemApp) GetAuth(
	sessionStruct *method.SessionStruct,
) (
	*[]entities.FunctionItem,
	*map[string](map[string](map[string]interface{})),
	*error,
) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		s.RoleRepo,
		s.CompanyBanchRepo,
		false,
		"",
		"",
	)
	
	if err != nil {
		return nil, nil, err
	}

	permission := &map[string](map[string](map[string]interface{})){}

	// 角色結構
	roleStructs, _ := s.RoleStructRepo.GetRoleStructs(&entities.RoleStruct{
		CompanyId: authAggregate.User.CompanyId,
		RoleId: authAggregate.User.RoleId,
	})

	// 該角色的角色結構表中所擁有的 funcCode
	roleSturctFuncCodes, _ := s.RoleStructRepo.GetRoleStructsFuncCode(&entities.RoleStruct{
		CompanyId: authAggregate.User.CompanyId,
		RoleId: authAggregate.User.RoleId,
	})

	// 功能項目表
	functionItem, _ := s.FunctionItemRepo.GetFunctionItemsByFuncCodes(roleSturctFuncCodes)

	// 整理權限的資料結構
	for _, value := range *roleStructs {
		// empty map init
		if (*permission)[value.FuncCode] == nil {
			(*permission)[value.FuncCode] = make(map[string]map[string]interface{})
		}

		if (*permission)[value.FuncCode][value.ItemCode] == nil {
			(*permission)[value.FuncCode][value.ItemCode] = make(map[string]interface{})
		}

		// 可編輯角色範圍 json decode
		if value.ScopeRole != "all" && value.ScopeRole != "self" {
			var scopeRole []int
			json.Unmarshal([]byte(value.ScopeRole), &scopeRole)
			(*permission)[value.FuncCode][value.ItemCode]["scopeRole"] = scopeRole
		} else {
			(*permission)[value.FuncCode][value.ItemCode]["scopeRole"] = value.ScopeRole
		}

		// 可編輯部門範圍 json decode
		if value.ScopeBanch != "all" && value.ScopeBanch != "self" {
			var scopeBanch []int
			json.Unmarshal([]byte(value.ScopeBanch), &scopeBanch)
			(*permission)[value.FuncCode][value.ItemCode]["scopeBanch"] = scopeBanch
		} else {
			(*permission)[value.FuncCode][value.ItemCode]["scopeBanch"] = value.ScopeBanch
		}
	}

	return functionItem, permission, nil
}

// 拿取功能項目表
func (s *SystemApp) GetFunc() (
	*[]entities.FunctionItem,
	*[]entities.OperationItem,
	*map[string]map[string]map[string][]string,
) {
	// 拿取功能項目表
	functionItem, _ := s.FunctionItemRepo.GetFunctionItems()

	// 拿取 功能項目 與 角色部門關聯表
	functionRoleBanchRelation, _ := s.FunctionRoleBanchRelationRepo.GetFunctionRoleBanchRelations()

	// 拿取操作項目表
	operationItem, _ := s.OperationItemRepo.GetOperationItems()

	// 整理 功能項目 與 角色部門關聯表
	functionRoleBanchRelationMap := map[string]map[string]map[string][]string{}
	for _, value := range *functionRoleBanchRelation {
		// 初始化
		if functionRoleBanchRelationMap[value.FuncCode] == nil {
			functionRoleBanchRelationMap[value.FuncCode] = map[string]map[string][]string{}
		}
		if functionRoleBanchRelationMap[value.FuncCode][value.ItemCode] == nil {
			functionRoleBanchRelationMap[value.FuncCode][value.ItemCode] = map[string][]string{}
		}
		functionRoleBanchRelationMap[value.FuncCode][value.ItemCode]["scopeRole"] = strings.Split(value.HasScopeRole, ",")
		functionRoleBanchRelationMap[value.FuncCode][value.ItemCode]["scopeBanch"] = strings.Split(value.HasScopeBanch, ",")
	}

	return functionItem, operationItem, &functionRoleBanchRelationMap
}

func (s *SystemApp) GetRoleBanchList(
	sessionStruct *method.SessionStruct,
) (
	*map[string]map[string]interface{},
	*map[string]map[string]interface{},
	*map[string]map[string]interface{},
	*[]entities.CompanyBanch,
	*[]entities.Role,
	*[]entities.User,
	*error,
) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		s.RoleRepo,
		s.CompanyBanchRepo,
		false,
		"",
		"",
	)
	
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	// 拿取功能項目表
	functionItem, _ := s.FunctionItemRepo.GetFunctionItems()

	// 拿取操作項目表
	operationItem, _ := s.OperationItemRepo.GetOperationItems()
	
	// data struct init
	scopeRole := make(map[string]map[string]interface{})
	scopeBanch := make(map[string]map[string]interface{})
	scopeUser := make(map[string]map[string]interface{})

	availableBanch, _ := s.CompanyBanchRepo.GetCompanyBanches(
		&entities.CompanyBanch{
			CompanyId: authAggregate.User.CompanyId,
		},
		nil,
	)

	availableRole, _ := s.RoleRepo.GetRoles(&entities.Role{
		CompanyId: authAggregate.User.CompanyId,
	})

	availableUser, _ := s.UserRepo.GetUsers(&entities.User{
		CompanyId: authAggregate.User.CompanyId,
	}, nil, nil)

	for _, FunctionItem := range *functionItem {
		scopeBanch[FunctionItem.FuncCode] = make(map[string]interface{})
		scopeRole[FunctionItem.FuncCode] = make(map[string]interface{})
		scopeUser[FunctionItem.FuncCode] = make(map[string]interface{})

		for _, operation := range *operationItem {

			scopeBanch[FunctionItem.FuncCode][operation.OperationCode] = authAggregate.CurrentPermissionScopeBanch
			scopeRole[FunctionItem.FuncCode][operation.OperationCode] = authAggregate.CurrentPermissionScopeRole

			var scopeUserIdArray []int
			userDatas, _ := s.UserRepo.GetUsers(
				&entities.User{
					CompanyId: authAggregate.User.CompanyId,
				},
				&authAggregate.CurrentPermissionScopeRole,
				&authAggregate.CurrentPermissionScopeBanch,
			)

			for _, user := range *userDatas{
				scopeUserIdArray = append(scopeUserIdArray, user.UserId)
			}

			scopeUser[FunctionItem.FuncCode][operation.OperationCode] = scopeUserIdArray
		}
	}

	return &scopeRole, &scopeBanch, &scopeUser, availableBanch, availableRole, availableUser, nil
}