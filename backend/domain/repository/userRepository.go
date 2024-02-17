package repository

import "backend/domain/entities"

type UserRepository interface {
	GetUser(*entities.User) (*entities.User, error)
	GetUsers() (*[]entities.User, error)
	UpdateUser(*entities.User) (*entities.User, *map[string]string)
	SaveUser(*entities.User) (*entities.User, *map[string]string)
	
	GetNewUserID(companyId int) int
	IsAccountDuplicated(account string) bool
}