package aggregates

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"backend/interfaces/method"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

/**
  @TODO 可以把 array param 換成是 pointer
*/
func inArray(array interface{}, val interface{}) (exists bool, index int) {
    exists = false
    index = -1

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }

    return
}

func convertSliceToInt(in []any) (out []int) {
    out = make([]int, 0, len(in))
    for _, v := range in {
        out = append(out, int(v.(float64)))
    }
    return
}

type AutAggregate struct {
	ID string
	User entities.User
	/**
		Permission = {
			[funcCode]: {
				[itemCode]: {
					scopeBanch: []BanchId | all | self, 
					scopeRole: []RoleId | all | self,
				}
			}
		}
	*/
	Permission map[string](map[string](map[string]interface{}))
	CurrentPermission map[string]interface{} // 當前的權限
	CurrentPermissionScopeBanch []int // 當前的 scope banch
	CurrentPermissionScopeRole []int // 當前的 scope role
	

	IsCurrentScopeBanchAll bool // scope banch 是否是 all (給add 看的)
	IsCurrentScopeRoleAll bool // scope role 是否是 all (給add 看的)
}

/*
	auth 基本處理
*/
func NewAuthAggregate(
	sessionStruct *method.SessionStruct,
	roleRepo repository.RoleRepository,
	banchRepo repository.CompanyBanchRepository,
	permissionValidation bool, // 是否開啟 權限驗證
	permissionFuncCode string,
	permissionItemCode string,
) (*AutAggregate, error){
	instance := new(AutAggregate)

	(*instance).IsCurrentScopeBanchAll = false
	(*instance).IsCurrentScopeRoleAll = false
	(*instance).User = *sessionStruct.User

	// 權限json decode
	if sessionStruct.Permission != nil {
		json.Unmarshal([]byte(sessionStruct.Permission.(string)), &(*instance).Permission)

		permission, OK := (*instance).Permission[permissionFuncCode][permissionItemCode]
		if !OK && permissionValidation {
			errMSG := fmt.Sprintf(
				"權限驗證失敗--[funcCode: '%s'][itemCode: '%s']",
				permissionFuncCode,
				permissionItemCode,
			)
			return nil, errors.New(errMSG)
		}

		// 可編輯角色範圍 的資料 搜尋 ( 分為自己，所有，自訂 )
		var scopeRole []int
		if permission["scopeRole"] == "all" {
			scopeRole = *roleRepo.GetRolesId(&entities.Role{
				CompanyId: sessionStruct.User.CompanyId,
			})

			// 設定 is current scope role all
			(*instance).IsCurrentScopeRoleAll = true
		} else if permission["scopeRole"] == "self" {
			scopeRole = append(scopeRole, sessionStruct.User.RoleId)
		} else if permission["scopeRole"] != nil {
			scopeRoleSlice := convertSliceToInt(
				permission["scopeRole"].([]any),
			)

			// 要把　自訂義裡面　可能被刪除的　roleId 過濾掉
			scopeRole = *roleRepo.GetRolesIdByScopeRole(
				&entities.Role{
					CompanyId: sessionStruct.User.CompanyId,
				},
				&scopeRoleSlice,
			)
		}

		// 可編輯部門範圍 的資料 搜尋 ( 分為自己，所有，自訂 )
		var scopeBanch  []int
		if permission["scopeBanch"] == "all" {
			scopeBanch = *banchRepo.GetBanchesId(&entities.CompanyBanch{
				CompanyId: sessionStruct.User.CompanyId,
			})
			
			// 設定 is current scope banch all
			(*instance).IsCurrentScopeBanchAll = true
		} else if permission["scopeBanch"] == "self" {
			scopeBanch = append(scopeBanch, *sessionStruct.User.BanchId)
		} else if permission["scopeBanch"] != nil {
			scopeBanchSlice := convertSliceToInt(
				permission["scopeBanch"].([]any),
			)

			// 要把 自訂義裡面 可能被刪除的 banchId過濾掉
			scopeBanch = *banchRepo.GetBanchesIdByScopeBanch(
				&entities.CompanyBanch{
					CompanyId: sessionStruct.User.CompanyId,
				},
				&scopeBanchSlice,
			)
		}

		// 綁定物件
		(*instance).CurrentPermission = permission
		(*instance).CurrentPermissionScopeBanch = scopeBanch
		(*instance).CurrentPermissionScopeRole = scopeRole
	}

	return instance, nil
}

// 檢查可編輯的 部門範圍
func(instance *AutAggregate) CheckScopeBanchValidation(banchId int) error {
	// 檢查是否可以加入此部門
	if exists, _ := inArray(
		(*instance).CurrentPermissionScopeBanch,
		banchId);
		!exists {
		return errors.New("無法插入此部門，尚無權限")
	}
	return nil
}

// 檢查可編輯的 角色範圍
func(instance *AutAggregate) CheckScopeRoleValidation(roleId int) error {
	// 檢查是否可以加入此角色
	if exists, _ := inArray(
		(*instance).CurrentPermissionScopeRole,
		roleId);
		!exists {
		return errors.New("無法插入此角色，尚無權限")
	}

	return nil
}

// 指定的banch id 與 目前可查詢的 banch
func(instance *AutAggregate) GetScopeBanchWithCustomize(banchId *int) *[]int {
	if banchId != nil {
		if exists, _ := inArray(
			(*instance).CurrentPermissionScopeBanch,
			banchId,
		); !exists {
			return &[]int{*banchId}
		}
		return &[]int{}
	}
	return &(*instance).CurrentPermissionScopeBanch
}

// 指定的role id 與 目前可查詢的 role
func(instance *AutAggregate) GetScopeRolehWithCustomize(roleId *int) *[]int {
	if roleId != nil {
		if exists, _ := inArray(
			(*instance).CurrentPermissionScopeRole,
			roleId,
		); !exists {
			return &[]int{*roleId}
		}
		return &[]int{}
	}
	return &(*instance).CurrentPermissionScopeRole
}