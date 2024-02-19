package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)


type RoleRepo struct {
	db *gorm.DB
	tableName string
}

func NewRoleRepository(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db, "role"}
}

var _ repository.RoleRepository = &RoleRepo{}

func (r *RoleRepo) SaveRole(roleEntity *entities.Role, TX *gorm.DB) (*entities.Role, *map[string]string) {
	roleEntity.RoleId = r.GetNewRoleID(roleEntity.CompanyId)

	err := TX.
		Debug().
		Table(r.tableName).
		Create(&roleEntity).
		Error

	return roleEntity, persistenceErrorHandler(err)
}

func (r *RoleRepo) GetRole(roleEntity *entities.Role) (*entities.Role, *map[string]string) {
	var role entities.Role

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleEntity.CompanyId).
		Where("roleId = ?", roleEntity.RoleId).
		Where("deleteFlag = ?", "N").
		First(role).
		Error

	return &role, persistenceErrorHandler(err)
}

func (r *RoleRepo) GetRoles(roleEntity *entities.Role) (*[]entities.Role, *map[string]string) {
	var roles []entities.Role

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleEntity.CompanyId).
		Where("deleteFlag = ?", "N").
		Where("roleName like ?", "%" + roleEntity.RoleName + "%").
		Order("sort asc").
		Find(&roles).
		Error

	return &roles, persistenceErrorHandler(err)
}

func (r *RoleRepo) GetRolesSelector(roleEntity *entities.Role) (*[]entities.Role, *map[string]string) {
	var roles []entities.Role

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleEntity.CompanyId).
		Order("sort asc").
		Find(&roles).
		Error

	return &roles, persistenceErrorHandler(err)
}


func (r *RoleRepo) UpdateRole(roleEntity *entities.Role, TX *gorm.DB) (*entities.Role, *map[string]string) {
	if r.IsRoleNameDuplicated(roleEntity) {
		return roleEntity, persistenceErrorHandler(errors.New("角色名稱重複"))
	}
	err := TX.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleEntity.CompanyId).
		Where("roleId = ?", roleEntity.RoleId).
		Updates(&roleEntity).
		Error

	return roleEntity, persistenceErrorHandler(err)
}

func (r *RoleRepo) DeleteRole(roleEntity *entities.Role) (*entities.Role, *map[string]string) {
	now := time.Now()
	roleEntity.DeleteFlag = "Y"
	roleEntity.DeleteTime = &now
	roleEntity.LastModify = &now

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", roleEntity.CompanyId).
		Where("roleId = ?", roleEntity.RoleId).
		Updates(&roleEntity).
		Error

	return roleEntity, persistenceErrorHandler(err)
}

// 拿取新的 role id ( max count + 1 )
func (r *RoleRepo) GetNewRoleID(companyId int) int {
    var MaxCount int64

	r.db.
		Debug().
        Where("companyId = ?", companyId).
        Select("max(roleId)").
        Row().
        Scan(&MaxCount)

    return int(MaxCount) + 1
}

// 查詢是否有重複role name
func (r *RoleRepo) IsRoleNameDuplicated(roleEntity *entities.Role) bool {
    var MaxCount int64

	r.db.
		Debug().
        Where("companyId = ?", (*roleEntity).CompanyId).
        Where("roleName = ?", (*roleEntity).RoleName).
        Where("roleId != ?", (*roleEntity).RoleId).
        Where("deleteFlag = ?", "N").
        Count(&MaxCount)

    return int(MaxCount) > 0
}

func (r *RoleRepo) GetRolesId(roleEntity *entities.Role) *[]int {
    var rolesIdArr []int

	r.db.
		Debug().
		Table(r.tableName).
		Select("roleId").
		Where("companyId = ?", roleEntity.CompanyId).
		Where("deleteFlag = ?", "N").
		Find(&entities.Role{}).
		Pluck("roleId", &rolesIdArr)

    return &rolesIdArr
}

func (r *RoleRepo) GetRolesIdByScopeRole(roleEntity *entities.Role, scopeRole *[]int) *[]int {
    var rolesIdArr []int

	r.db.
		Debug().
		Table(r.tableName).
		Select("roleId").
		Where("companyId = ?", roleEntity.CompanyId).
		Where("deleteFlag = ?", "N").
		Where("roleId in (?)", *scopeRole).
		Find(&entities.Role{}).
		Pluck("roleId", &rolesIdArr)

    return &rolesIdArr
}

func (r *RoleRepo) Begin() *gorm.DB {
	return r.db.Begin()
}