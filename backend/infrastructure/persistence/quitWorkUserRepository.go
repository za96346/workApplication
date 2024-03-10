package persistence

import (
	"backend/domain/repository"

	"gorm.io/gorm"
)

type QuitWorkUserRepo struct {
	db        *gorm.DB
	tableName string
}

func NewQuitWorkUserRepository(db *gorm.DB) *QuitWorkUserRepo {
	return &QuitWorkUserRepo{db, "quit_work_user"}
}

var _ repository.QuitWorkUserRepository = &QuitWorkUserRepo{}
