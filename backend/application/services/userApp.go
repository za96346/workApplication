package application

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type UserApp struct {
	userRepo repository.UserRepository
}

var _ UserAppInterface = &UserApp{}

type UserAppInterface interface {
	GetMine(*entities.User) (*entities.User, *map[string]string)
	GetUsers(
		*entities.User,
		*[]int,
		*[]int,
	) (*[]entities.User, *map[string]string)
	GetUsersSelector(*entities.User) (*[]entities.User, *map[string]string)

	UpdateUser(*entities.User) (*entities.User, *map[string]string)
	UpdatePassword(*entities.User) (*entities.User, *map[string]string)
	UpdateMine(*entities.User) (*entities.User, *map[string]string)

	SaveUser(*entities.User) (*entities.User, *map[string]string)
	DeleteUser(*entities.User) (*entities.User, *map[string]string)
}

func (u *UserApp) GetMine(userEntity *entities.User) (*entities.User, *map[string]string) {
	user, err := u.userRepo.GetUser(userEntity)
	user.Password = ""
	return user, err
}

func (u *UserApp) GetUsers(
	userEntity *entities.User,
	scopeBanch *[]int,
	scopeRole *[]int,
) (*[]entities.User, *map[string]string) {
	return u.userRepo.GetUsers(
		userEntity,
		scopeBanch,
		scopeRole,
	)
}

func (u *UserApp) GetUsersSelector(userEntity *entities.User) (*[]entities.User, *map[string]string) {
	return u.userRepo.GetUsersSelector(userEntity)
}

func (u *UserApp) UpdateUser(userEntity *entities.User) (*entities.User, *map[string]string) {
	return u.userRepo.UpdateUser(userEntity)
}

func (u *UserApp) UpdatePassword(userEntity *entities.User) (*entities.User, *map[string]string) {
	return &entities.User{}, &map[string]string{}
}

func (u *UserApp) UpdateMine(userEntity *entities.User) (*entities.User, *map[string]string) {
	return u.userRepo.UpdateUser(userEntity)
}

func (u *UserApp) SaveUser(userEntity *entities.User) (*entities.User, *map[string]string) {
	return u.userRepo.SaveUser(userEntity)
}

func (u *UserApp) DeleteUser(*entities.User) (*entities.User, *map[string]string) {
	return &entities.User{}, &map[string]string{}
}