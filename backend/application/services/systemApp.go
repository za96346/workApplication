package application

import (
	"backend/domain/repository"
)

type SystemApp struct {
	roleRepo repository.RoleRepository
}

var _ SystemAppInterface = &SystemApp{}

type SystemAppInterface interface {
	GetAuth() (interface{}, error)
	GetFunc() (interface{}, error)
	GetRoleBanchList() (interface{}, error)
}