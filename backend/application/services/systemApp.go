package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"encoding/json"
	"strings"
)

type SystemApp struct {
	roleRepo repository.RoleRepository
	roleStructRepo repository.RoleStructRepository
	functionItemRepo repository.FunctionItemRepository
	functionRoleBanchRelationRepo repository.FunctionRoleBanchRelationRepository
	operationItemRepo repository.OperationItemRepository
	companyBanchRepo repository.CompanyBanchRepository
	userRepo repository.UserRepository
}

var _ SystemAppInterface = &SystemApp{}

type SystemAppInterface interface {
	GetAuth(companyId int, roleId int) (*[]entities.FunctionItem, *map[string](map[string](map[string]interface{})))
	GetFunc() (
		*[]entities.FunctionItem,
		*[]entities.OperationItem,
		*map[string]map[string]map[string][]string,
	)
	GetRoleBanchList(
		companyId int,
		currentPermissionScopeBanch *[]int,
		currentPermissionScopeRole *[]int,
	) (
		*map[string]map[string]interface{},
		*map[string]map[string]interface{},
		*map[string]map[string]interface{},
		*[]entities.CompanyBanch,
		*[]entities.Role,
		*[]entities.User,
	)
}

// 獲取權限
func (s *SystemApp) GetAuth(companyId int, roleId int) (*[]entities.FunctionItem, *map[string](map[string](map[string]interface{}))) {
	permission := &map[string](map[string](map[string]interface{})){}

	// 角色結構
	roleStructs, _ := s.roleStructRepo.GetRoleStructs(&entities.RoleStruct{
		CompanyId: companyId,
		RoleId: roleId,
	})

	// 該角色的角色結構表中所擁有的 funcCode
	roleSturctFuncCodes, _ := s.roleStructRepo.GetRoleStructsFuncCode(&entities.RoleStruct{
		CompanyId: companyId,
		RoleId: roleId,
	})

	// 功能項目表
	functionItem, _ := s.functionItemRepo.GetFunctionItemsByFuncCodes(roleSturctFuncCodes)

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

	return functionItem, permission
}

// 拿取功能項目表
func (s *SystemApp) GetFunc() (
	*[]entities.FunctionItem,
	*[]entities.OperationItem,
	*map[string]map[string]map[string][]string,
) {
	// 拿取功能項目表
	functionItem, _ := s.functionItemRepo.GetFunctionItems()

	// 拿取 功能項目 與 角色部門關聯表
	functionRoleBanchRelation, _ := s.functionRoleBanchRelationRepo.GetFunctionRoleBanchRelations()

	// 拿取操作項目表
	operationItem, _ := s.operationItemRepo.GetOperationItems()

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
	companyId int,
	currentPermissionScopeBanch *[]int,
	currentPermissionScopeRole *[]int,
) (
	*map[string]map[string]interface{},
	*map[string]map[string]interface{},
	*map[string]map[string]interface{},
	*[]entities.CompanyBanch,
	*[]entities.Role,
	*[]entities.User,
) {
	// 拿取功能項目表
	functionItem, _ := s.functionItemRepo.GetFunctionItems()

	// 拿取操作項目表
	operationItem, _ := s.operationItemRepo.GetOperationItems()
	
	// data struct init
	scopeRole := make(map[string]map[string]interface{})
	scopeBanch := make(map[string]map[string]interface{})
	scopeUser := make(map[string]map[string]interface{})

	availableBanch, _ := s.companyBanchRepo.GetCompanyBanches(
		&entities.CompanyBanch{
			CompanyId: companyId,
		},
		nil,
	)

	availableRole, _ := s.roleRepo.GetRoles(&entities.Role{
		CompanyId: companyId,
	})

	availableUser, _ := s.userRepo.GetUsers(&entities.User{
		CompanyId: companyId,
	}, nil, nil)

	for _, FunctionItem := range *functionItem {
		scopeBanch[FunctionItem.FuncCode] = make(map[string]interface{})
		scopeRole[FunctionItem.FuncCode] = make(map[string]interface{})
		scopeUser[FunctionItem.FuncCode] = make(map[string]interface{})

		for _, operation := range *operationItem {

			scopeBanch[FunctionItem.FuncCode][operation.OperationCode] = currentPermissionScopeBanch
			scopeRole[FunctionItem.FuncCode][operation.OperationCode] = currentPermissionScopeRole

			var scopeUserIdArray []int
			userDatas, _ := s.userRepo.GetUsers(
				&entities.User{
					CompanyId: companyId,
				},
				currentPermissionScopeRole,
				currentPermissionScopeBanch,
			)

			for _, user := range *userDatas{
				scopeUserIdArray = append(scopeUserIdArray, user.UserId)
			}

			scopeUser[FunctionItem.FuncCode][operation.OperationCode] = scopeUserIdArray
		}
	}

	return &scopeRole, &scopeBanch, &scopeUser, availableBanch, availableRole, availableUser
}