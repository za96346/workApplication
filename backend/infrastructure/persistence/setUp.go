package persistence

import (
	"fmt"
	"backend/domain/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	Company repository.CompanyRepository
	CompanyBanch repository.CompanyBanchRepository
	FunctionItem repository.FunctionItemRepository
	FunctionRoleBanchRelation repository.FunctionRoleBanchRelationRepository
	Log repository.LogRepository
	OperationItem repository.OperationItemRepository
	Performance repository.PerformanceRepository
	QuitWorkUser repository.QuitWorkUserRepository
	Role repository.RoleRepository
	RoleStruct repository.RoleStructRepository
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		Company: NewCompanyRepository(db),
		db:   db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}
