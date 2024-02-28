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
	GetMine(*entities.User) (*entities.User, *error)
	GetUsers(
		*entities.User,
		*[]int,
		*[]int,
	) (*[]entities.User, *error)
	GetUsersSelector(*entities.User) (*[]entities.User, *error)

	UpdateUser(*entities.User) (*entities.User, *error)
	UpdatePassword(*entities.User) (*entities.User, *error)
	UpdateMine(*entities.User) (*entities.User, *error)

	SaveUser(*entities.User) (*entities.User, *error)
	DeleteUser(*entities.User) (*entities.User, *error)
}

func (u *UserApp) GetMine(userEntity *entities.User) (*entities.User, *error) {
	user, err := u.userRepo.GetUser(userEntity)
	user.Password = ""
	return user, err
}

func (u *UserApp) GetUsers(
	userEntity *entities.User,
	scopeBanch *[]int,
	scopeRole *[]int,
) (*[]entities.User, *error) {
	return u.userRepo.GetUsers(
		userEntity,
		scopeBanch,
		scopeRole,
	)
}

func (u *UserApp) GetUsersSelector(userEntity *entities.User) (*[]entities.User, *error) {
	return u.userRepo.GetUsersSelector(userEntity)
}

func (u *UserApp) UpdateUser(userEntity *entities.User) (*entities.User, *error) {
	return u.userRepo.UpdateUser(userEntity)
}

func (u *UserApp) UpdatePassword(userEntity *entities.User) (*entities.User, *error) {
	return &entities.User{}, nil
}

func (u *UserApp) UpdateMine(userEntity *entities.User) (*entities.User, *error) {
	return u.userRepo.UpdateUser(userEntity)
}

func (u *UserApp) SaveUser(userEntity *entities.User) (*entities.User, *error) {
	return u.userRepo.SaveUser(userEntity)
}

func (u *UserApp) DeleteUser(userEntity *entities.User) (*entities.User, *error) {
	return u.userRepo.DeleteUser(userEntity)
}