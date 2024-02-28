package repository

import "backend/domain/entities"

type UserRepository interface {
	GetUser(userEntity *entities.User) (*entities.User, *error)
	GetUserByAccount(userEntity *entities.User) (*entities.User, *error)
	GetUsers(
		userEntity *entities.User,
		roleScope *[]int,
		banchScope *[]int,
	) (*[]entities.User, *error)
	GetUsersSelector(
		userEntity *entities.User,
	) (*[]entities.User, *error)

	UpdateUser(userEntity *entities.User) (*entities.User, *error)
	SaveUser(userEntity *entities.User) (*entities.User, *error)
	DeleteUser(userEntity *entities.User) (*entities.User, *error)
	
	GetNewUserID(int) int
	IsAccountDuplicated(string) bool
}