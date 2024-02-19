package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)


type UserRepo struct {
	db *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db, "user"}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) GetUser(userEntity *entities.User) (*entities.User, *map[string]string) {
	var user entities.User
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("userId = ?", (*userEntity).UserId).
		Where("companyId = ?", (*userEntity).CompanyId).
		First(&user).
		Error

	return &user, persistenceErrorHandler(err)
}

func (r *UserRepo) GetUserByAccount(userEntity *entities.User) (*entities.User, *map[string]string) {
	var user entities.User
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("account = ?", (*userEntity).Account).
		Where("companyId = ?", (*userEntity).CompanyId).
		First(&user).
		Error

	return &user, persistenceErrorHandler(err)
}

func (r *UserRepo) GetUsers(
	userEntity *entities.User,
	roleScope *[]int,
	banchScope *[]int,
) (*[]entities.User, *map[string]string) {
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
		Where("roleId in (?)", roleScope).
		Where("banchId in (?)", banchScope).
		Where("deleteFlag", "N").
		Order("sort asc")

	// 使用者名稱
	if &userEntity.UserName != nil {
		searchQuery.Where("userName like ?", "%" + userEntity.UserName + "%")
	}

	// 員工編號
	if &userEntity.EmployeeNumber != nil {
		searchQuery.Where("employeeNumber like ?", "%" + userEntity.EmployeeNumber + "%")
	}

	// 離職狀態
	if &userEntity.QuitFlag != nil {
		searchQuery.Where("quitFlag = ?", userEntity.QuitFlag)
	}

	err := searchQuery.Find(&users).Error

	return &users, persistenceErrorHandler(err)
}

func (r *UserRepo) GetUsersSelector(
	userEntity *entities.User,
) (*[]entities.User, *map[string]string) {

	var users []entities.User

	searchQuery := r.db.
		Debug().
		Table(r.tableName).
		Where("companyId = ?", userEntity.CompanyId).
		Order("sort asc")

	// 使用者名稱
	if &userEntity.UserName != nil {
		searchQuery.Where("userName like ?", "%" + userEntity.UserName + "%")
	}

	// 員工編號
	if &userEntity.EmployeeNumber != nil {
		searchQuery.Where("employeeNumber like ?", "%" + userEntity.EmployeeNumber + "%")
	}

	// 部門查詢
	if &userEntity.BanchId != nil {
		searchQuery.Where("banchId = ?", userEntity.BanchId)
	}

	// 角色查詢
	if &userEntity.RoleId != nil {
		searchQuery.Where("roleId = ?", userEntity.RoleId)
	}

	err := searchQuery.Find(&users).Error

	return &users, persistenceErrorHandler(err)
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

func (r *UserRepo) SaveUser(userEntity *entities.User) (*entities.User, *map[string]string) {
	// 加入一些固定欄位
	now := time.Now()

	if (*r).IsAccountDuplicated(userEntity.Account) {
		return nil, persistenceErrorHandler(errors.New("帳號重複"))
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

	return userEntity, persistenceErrorHandler(err)
}


func (r *UserRepo) UpdateUser(userEntity *entities.User) (*entities.User, *map[string]string) {
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

	return userEntity, persistenceErrorHandler(err)
}

func (r *UserRepo) DeleteUser(userEntity *entities.User) (*entities.User, *map[string]string) {
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

	return userEntity, persistenceErrorHandler(err)
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