package persistence

import (
	"backend/domain/repository"
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	Company                   repository.CompanyRepository
	CompanyBanch              repository.CompanyBanchRepository
	FunctionItem              repository.FunctionItemRepository
	FunctionRoleBanchRelation repository.FunctionRoleBanchRelationRepository
	Log                       repository.LogRepository
	OperationItem             repository.OperationItemRepository
	Performance               repository.PerformanceRepository
	QuitWorkUser              repository.QuitWorkUserRepository
	Role                      repository.RoleRepository
	RoleStruct                repository.RoleStructRepository
	User                      repository.UserRepository
	db                        *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	port, _ := strconv.Atoi(DbPort)
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser,
		DbPassword,
		DbHost,
		port,
		DbName,
	)
	db, err := gorm.Open(
		mysql.Open(
			DSN,
		),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	return &Repositories{
		Company: NewCompanyRepository(db),
		CompanyBanch: NewCompanyBanchRepository(db),
		FunctionItem: NewFunctionItemRepository(db),
		FunctionRoleBanchRelation: NewFunctionRoleBanchRelationRepository(db),
		Log: NewLogRepository(db),
		OperationItem: NewOperationItemRepository(db),
		Performance: NewPerformanceRepository(db),
		QuitWorkUser: NewQuitWorkUserRepository(db),
		Role: NewRoleRepository(db),
		RoleStruct: NewRoleStructRepository(db),
		User: NewUserRepository(db),
		db:      db,
	}, nil
}

// closes the  database connection
func (s *Repositories) Close() error {
	return nil
}
