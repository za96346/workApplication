package persistence

import (
	"errors"
	"backend/domain/entities"
	"backend/domain/repository"
	"github.com/jinzhu/gorm"
	"strings"
)


type UserRepo struct {
	db *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db, "user"}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(userEntity *entities.User) (*entities.User, *map[string]string) {
	dbErr := map[string]string{}
	err := r.db.
		Debug().
		Table(r.tableName).
		Create(&userEntity).
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
	return userEntity, nil
}

func (r *UserRepo) GetUser(userEntity *entities.User) (*entities.User, error) {
	var user entities.User
	err := r.db.
		Debug().
		Table(r.tableName).
		Where("userId = ?", (*userEntity).UserId).
		Take(&userEntity).
		Error

	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *UserRepo) GetUsers() (*[]entities.User, error) {
	var users []entities.User
	err := r.db.
		Debug().
		Table(r.tableName).
		Find(&users).
		Error

	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &users, nil
}


func (r *UserRepo) UpdateUser(userEntity *entities.User) (*entities.User, *map[string]string) {
	dbErr := map[string]string{}
	err := r.db.
		Debug().
		Table(r.tableName).
		Save(&userEntity).
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
	return userEntity, nil
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