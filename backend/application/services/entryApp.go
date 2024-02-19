package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type EntryApp struct {
	userRepo repository.UserRepository
}

var _ EntryAppInterface = &EntryApp{}

type EntryAppInterface interface {
	Login(*entities.User) (*entities.User, bool)
}

func (u *EntryApp) Login(userEntity *entities.User) (*entities.User, bool) {
	user, _ := u.userRepo.GetUserByAccount(userEntity)
	if !(
		(*user).Password == userEntity.Password && 
		(*user).Account == userEntity.Account &&
		(*user).Account != "" &&
		(*user).Password != "") {

		return nil, false
	}

	// 驗證是否是離職
	if (*user).IsQuitWorking() {
		return nil, false
	}

	return user, true
}