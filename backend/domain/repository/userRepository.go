package repository

import "backend/domain/entities"

type UserRepository interface {
	GetUser(*entities.User) (*entities.User, *map[string]string)
	GetUserByAccount(*entities.User) (*entities.User, *map[string]string)
	GetUsers(*entities.User, *[]int, *[]int) (*[]entities.User, *map[string]string)
	GetUsersSelector(*entities.User) (*[]entities.User, *map[string]string)

	UpdateUser(*entities.User) (*entities.User, *map[string]string)
	SaveUser(*entities.User) (*entities.User, *map[string]string)
	
	GetNewUserID(int) int
	IsAccountDuplicated(string) bool
}