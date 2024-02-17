package persistence

import (
	"errors"
	"backend/domain/entities"
	"backend/domain/repository"
	"github.com/jinzhu/gorm"
	"strings"
)


type RoleRepo struct {
	db *gorm.DB
	tableName string
}

func NewRoleRepository(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db, "role"}
}

var _ repository.RoleRepository = &RoleRepo{}

func (r *RoleRepo) SaveRole(roleEntity *entities.Role) (*entities.Role, *map[string]string) {
	dbErr := map[string]string{}
	err := r.db.
		Debug().
		Table(r.tableName).
		Create(&roleEntity).
		Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["email_taken"] = "email already taken"
			return nil, &dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, &dbErr
	}
	return roleEntity, nil
}

func (r *RoleRepo) GetRole(roleEntity *entities.Role) (*entities.Role, error) {
	var role entities.Role
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("roleId = ?", (*roleEntity).RoleId).
		Take(&roleEntity).
		Error

	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("role not found")
	}
	return &role, nil
}

func (r *RoleRepo) GetRoles() (*[]entities.Role, error) {
	var roles []entities.Role
	err := r.db.
		Debug().
		Table(r.tableName).
		Find(&roles).
		Error

	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("role not found")
	}
	return &roles, nil
}


func (r *RoleRepo) UpdateRole(roleEntity *entities.Role) (*entities.Role, *map[string]string) {
	dbErr := map[string]string{}
	err := r.db.
		Debug().
		Table(r.tableName).
		Save(&roleEntity).
		Error

	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, &dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, &dbErr
	}
	return roleEntity, nil
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
