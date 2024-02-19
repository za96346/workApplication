package persistence

import (
	"backend/domain/repository"
	"fmt"
	"strings"

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


func persistenceErrorHandler(err error) *map[string]string {
	dbErr := map[string]string{}

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			dbErr["notFound"] = "recorder not found"
		}
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
		}
	}
	return  &dbErr
}