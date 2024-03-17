package aggregates

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type RoleAggregate struct {
	ID string
	Role *entities.Role
	RoleStructs *[]entities.RoleStruct

	RoleStructRepo repository.RoleStructRepository
}

func NewRoleAggregate(
	role *entities.Role,
	roleStructs *[]entities.RoleStruct,
	roleStructRepo repository.RoleStructRepository,
) *RoleAggregate {
	return &RoleAggregate{
		ID: "",
		Role: role,
		RoleStructs: roleStructs,
		RoleStructRepo: roleStructRepo,
	}
}

func (r *RoleAggregate) AddRoleStruct(
	data *map[string](map[string](map[string]interface{})),
	TX *gorm.DB,
) error {
	now := time.Now()
	// 先把 此role structure 的資料 刪除
	r.RoleStructRepo.DeleteRoleStructs(
		&entities.RoleStruct{
			CompanyId: r.Role.CompanyId,
			RoleId: r.Role.RoleId,
		},
		TX,
	)

	// 在 寫入 新的 進入 db
	for funcCode, itemObject := range *data {
		for itemCode, scopeObject := range itemObject {

			// 可編輯部門範圍
			scopeBanch := ""
			if scopeObject["scopeBanch"] == "all" || scopeObject["scopeBanch"] == "self" {
				scopeBanch = scopeObject["scopeBanch"].(string)
			} else {
				scopeBanchByte, _ := json.Marshal(scopeObject["scopeBanch"])
				scopeBanch = string(scopeBanchByte)
			}

			// 可編輯角色範圍
			scopeRole := ""
			if scopeObject["scopeRole"] == "all" || scopeObject["scopeRole"] == "self" {
				scopeRole = scopeObject["scopeRole"].(string)
			} else {
				scopeRoleByte, _ :=json.Marshal(scopeObject["scopeRole"])
				scopeRole = string(scopeRoleByte)
			}

			// 可編輯角色範圍
			scopeUser := ""
			if scopeObject["scopeUser"] == "all" || scopeObject["scopeUser"] == "self" {
				scopeUser = scopeObject["scopeUser"].(string)
			} else {
				scopeUserByte, _ :=json.Marshal(scopeObject["scopeUser"])
				scopeUser = string(scopeUserByte)
			}

			// 新增
			r.RoleStructRepo.SaveRoleStruct(
				&entities.RoleStruct{
					CompanyId: r.Role.CompanyId,
					RoleId: r.Role.RoleId,
					FuncCode: funcCode,
					ItemCode: itemCode,
					ScopeBanch: scopeBanch,
					ScopeRole: scopeRole,
					ScopeUser: scopeUser,
					CreateTime: &now,
					LastModify: &now,
				},
				TX,
			)
		}
	}
	return nil
}
