package application

import (
	"backend/domain/repository"
)

type EntryApp struct {
	userRepo repository.UserRepository
}

var _ EntryAppInterface = &EntryApp{}

type EntryAppInterface interface {
	Login() (bool)
}

func (u *EntryApp) Login() (bool) {

}