package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"errors"
)

type EntryApp struct {
	UserRepo repository.UserRepository
}

var _ EntryAppInterface = &EntryApp{}

type EntryAppInterface interface {
	Login(*entities.User) (*entities.User, *error)
}

func (u *EntryApp) Login(userEntity *entities.User) (*entities.User, *error) {
	user, _ := u.UserRepo.GetUserByAccount(userEntity)
	if !(
		(*user).Password == userEntity.Password && 
		(*user).Account == userEntity.Account &&
		(*user).Account != "" &&
		(*user).Password != "") {
		err := errors.New("帳號或密碼錯誤")
		return nil, &err
	}

	// 驗證是否是離職
	if (*user).IsQuitWorking() {
		err := errors.New("已離職")
		return nil, &err
	}

	return user, nil
}