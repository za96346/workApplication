package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepo struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db, "user"}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) GetUser(userEntity *entities.User) (*entities.User, error) {
	var user entities.User
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("userId = ?", (*userEntity).UserId).
		Where("companyId = ?", (*userEntity).CompanyId).
		First(&user).
		Error

	return &user, err
}

func (r *UserRepo) GetUserByAccount(userEntity *entities.User) (*entities.User, error) {
	var user entities.User
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("account = ?", (*userEntity).Account).
		First(&user).
		Error

	return &user, err
}

func (r *UserRepo) GetUsers(
	userEntity *entities.User,
	roleScope *[]int,
	banchScope *[]int,
) (*[]entities.User, error) {
	var users []entities.User

	searchQuery := r.db.
		Debug().
		Table(r.tableName).
		Select(`
			companyId,
			userId,
			roleId,
			banchId,
			userName,
			employeeNumber,
			account,
			onWorkDay,
			quitFlag,
			deleteFlag,
			deleteTime,
			createTime,
			lastModify
		`).
		Where("companyId = ?", userEntity.CompanyId).
		Where("deleteFlag", "N").
		Order("sort asc")

	// 角色範圍查詢
	if roleScope != nil {
		searchQuery.Where("roleId in (?)", *roleScope)
	}

	// 部門範圍查詢
	if banchScope != nil {
		searchQuery.Where("banchId in (?)", *banchScope)
	}

	// 使用者名稱
	if userEntity.UserName != "" {
		searchQuery.Where("userName like ?", "%"+userEntity.UserName+"%")
	}

	// 員工編號
	if userEntity.EmployeeNumber != "" {
		searchQuery.Where("employeeNumber like ?", "%"+userEntity.EmployeeNumber+"%")
	}

	// 離職狀態
	if userEntity.QuitFlag != "" {
		searchQuery.Where("quitFlag = ?", userEntity.QuitFlag)
	}

	err := searchQuery.Find(&users).Error

	return &users, err
}

func (r *UserRepo) GetUsersSelector(
	userEntity *entities.User,
) (*[]entities.User, error) {

	var users []entities.User

	searchQuery := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", userEntity.CompanyId).
		Order("sort asc")

	// 使用者名稱
	if userEntity.UserName != "" {
		searchQuery.Where("userName like ?", "%"+userEntity.UserName+"%")
	}

	// 員工編號
	if userEntity.EmployeeNumber != "" {
		searchQuery.Where("employeeNumber like ?", "%"+userEntity.EmployeeNumber+"%")
	}

	// 部門查詢
	if userEntity.BanchId != nil {
		searchQuery.Where("banchId = ?", userEntity.BanchId)
	}

	// 角色查詢
	if userEntity.RoleId != 0 {
		searchQuery.Where("roleId = ?", userEntity.RoleId)
	}

	err := searchQuery.Find(&users).Error

	return &users, err
}

// 拿取新的 user id ( max count + 1 )
func (r *UserRepo) GetNewUserID(companyId int) int {
	var MaxCount int64

	r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", companyId).
		Select("max(userId)").
		Row().
		Scan(&MaxCount)

	return int(MaxCount) + 1
}

func (r *UserRepo) SaveUser(userEntity *entities.User) (*entities.User, error) {
	// 加入一些固定欄位
	now := time.Now()

	if (*r).IsAccountDuplicated(userEntity.Account) {
		err := errors.New("帳號重複")
		return nil, err
	}

	(*userEntity).UserId = (*r).GetNewUserID((*userEntity).CompanyId)

	(*userEntity).CreateTime = &now
	(*userEntity).LastModify = &now
	(*userEntity).DeleteTime = nil
	(*userEntity).DeleteFlag = "N"
	(*userEntity).QuitFlag = "N"

	err := r.db.
		Debug().
		Table(r.tableName).
		Create(userEntity).
		Error

	return userEntity, err
}

func (r *UserRepo) UpdateUser(userEntity *entities.User) (*entities.User, error) {
	// 加入一些固定欄位
	now := time.Now()

	(*userEntity).LastModify = &now
	(*userEntity).DeleteTime = nil
	(*userEntity).DeleteFlag = "N"

	// 更新
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", userEntity.CompanyId).
		Where("userId = ?", userEntity.UserId).
		Updates(&userEntity).
		Error

	return userEntity, err
}

func (r *UserRepo) DeleteUser(userEntity *entities.User) (*entities.User, error) {
	// 加入一些固定欄位
	now := time.Now()

	(*userEntity).LastModify = &now
	(*userEntity).DeleteTime = &now
	(*userEntity).DeleteFlag = "Y"

	// 更新
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", userEntity.CompanyId).
		Where("userId = ?", userEntity.UserId).
		Updates(&userEntity).
		Error

	return userEntity, err
}

// 帳號是否重複
func (r *UserRepo) IsAccountDuplicated(account string) bool {
	var MaxCount int64

	r.db.
		Debug().
		Table(r.tableName).
		Where("account = ?", account).
		Count(&MaxCount)

	return MaxCount > 0
}
