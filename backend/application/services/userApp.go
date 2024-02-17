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
	GetMine(*entities.User) (*entities.User, error)
	GetUsers(*entities.User) (*[]entities.User, error)
	GetUsersSelector(*entities.User) (*[]entities.User, error)

	UpdateUser(*entities.User) (*entities.User, *map[string]string)
	UpdatePassword(*entities.User) (*entities.User, *map[string]string)
	UpdateMine(*entities.User) (*entities.User, *map[string]string)

	SaveUser(*entities.User) (*entities.User, *map[string]string)
	DeleteUser(*entities.User) (*entities.User, *map[string]string)
}

func (u *UserApp) GetMine(*entities.User) (*entities.User, error) {

}

func (u *UserApp) GetUsers(*entities.User) (*[]entities.User, error) {

}

func (u *UserApp) GetUsersSelector(*entities.User) (*[]entities.User, error) {

}

func (u *UserApp) UpdateUser(*entities.User) (*entities.User, *map[string]string) {

}

func (u *UserApp) UpdatePassword(*entities.User) (*entities.User, *map[string]string) {

}

func (u *UserApp) UpdateMine(*entities.User) (*entities.User, *map[string]string) {
	
}

func (u *UserApp) SaveUser(*entities.User) (*entities.User, *map[string]string) {
	
}

func (u *UserApp) DeleteUser(*entities.User) (*entities.User, *map[string]string) {
	
}