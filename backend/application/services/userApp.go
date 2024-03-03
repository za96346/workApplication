package application

import (
	"backend/application/dtos"
	"backend/domain/aggregates"
	"backend/domain/entities"
	"backend/domain/repository"
	"backend/enum"
	"backend/interfaces/method"
	"errors"
)

type UserApp struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
	companyBanchRepo repository.CompanyBanchRepository
}

var _ UserAppInterface = &UserApp{}

type UserAppInterface interface {
	GetMine(
		userEntity *entities.User,
		sessionStruct *method.SessionStruct,
	) (*entities.User, *error)

	GetUsers(
		userEntity *entities.User,
		sessionStruct *method.SessionStruct,
	) (*[]entities.User, *error)

	GetUsersSelector(
		userEntity *entities.User,
		sessionStruct *method.SessionStruct,
	) (*[]entities.User, *error) 

	UpdateUser(
		userEntity *entities.User,
		sessionStruct *method.SessionStruct,
	) (*entities.User, *error)

	UpdatePassword(
		updatePwdDto *dtos.UserPasswordUpdateQueryParams,
		sessionStruct *method.SessionStruct,
	) (*entities.User, *error)

	UpdateMine(
		userEntity *entities.User,
		sessionStruct *method.SessionStruct,
	) (*entities.User, *error)

	SaveUser(
		userEntity *entities.User,
		sessionStruct *method.SessionStruct,
	) (*entities.User, *error)

	DeleteUser(
		userEntity *entities.User,
		sessionStruct *method.SessionStruct,
	) (*entities.User, *error)
}

func (u *UserApp) GetMine(userEntity *entities.User, sessionStruct *method.SessionStruct) (*entities.User, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		u.roleRepo,
		u.companyBanchRepo,
		false,
		"",
		"",
	)
	
	if err != nil {
		return nil, err
	}

	userEntity = &authAggregate.User

	user, err := u.userRepo.GetUser(userEntity)
	user.Password = ""
	return user, err
}

func (u *UserApp) GetUsers(
	userEntity *entities.User,
	sessionStruct *method.SessionStruct,
) (*[]entities.User, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		u.roleRepo,
		u.companyBanchRepo,
		true,
		string(enum.EmployeeManage),
		string(enum.Inquire),
	)
	
	if err != nil {
		return nil, err
	}

	userEntity.CompanyId = authAggregate.User.CompanyId

	return u.userRepo.GetUsers(
		userEntity,
		authAggregate.GetScopeBanchWithCustomize(userEntity.BanchId),
		authAggregate.GetScopeRolehWithCustomize(&userEntity.RoleId),
	)
}

func (u *UserApp) GetUsersSelector(
	userEntity *entities.User,
	sessionStruct *method.SessionStruct,
) (*[]entities.User, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		u.roleRepo,
		u.companyBanchRepo,
		false,
		"",
		"",
	)
	
	if err != nil {
		return nil, err
	}

	userEntity.CompanyId = authAggregate.User.CompanyId
	return u.userRepo.GetUsersSelector(userEntity)
}

func (u *UserApp) UpdateUser(
	userEntity *entities.User,
	sessionStruct *method.SessionStruct,
) (*entities.User, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		u.roleRepo,
		u.companyBanchRepo,
		true,
		string(enum.EmployeeManage),
		string(enum.Edit),
	)
	
	if err != nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeBanchValidation(*(*userEntity).BanchId); err != nil {
		return nil, &err
	}

	if err := authAggregate.CheckScopeRoleValidation((*userEntity).RoleId); err != nil {
		return nil, &err
	}

	// 檢驗欄位
	if userEntity.UserId == 0 {
		err := errors.New("更新失敗，UserId is nil.")
		return nil, &err
	}

	userEntity.CompanyId = authAggregate.User.CompanyId
	return u.userRepo.UpdateUser(userEntity)
}

func (u *UserApp) UpdatePassword(
	updatePwdDto *dtos.UserPasswordUpdateQueryParams,
	sessionStruct *method.SessionStruct,
) (*entities.User, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		u.roleRepo,
		u.companyBanchRepo,
		true,
		string(enum.EmployeeManage),
		string(enum.Edit),
	)
	
	if err != nil {
		return nil, err
	}

	userEntity, _ := u.userRepo.GetUser(&entities.User{
		CompanyId: authAggregate.User.CompanyId,
		UserId: updatePwdDto.UserId,
	})

	if err := authAggregate.CheckScopeBanchValidation(*(*userEntity).BanchId); err != nil {
		return nil, &err
	}

	if err := authAggregate.CheckScopeRoleValidation((*userEntity).RoleId); err != nil {
		return nil, &err
	}

	// 驗證 密碼
	if updatePwdDto.OldPassword != userEntity.Password ||
		updatePwdDto.NewPassword != updatePwdDto.NewPasswordAgain {
		
		err := errors.New("舊密碼不相符, 或 新密碼不相符")
		return nil, &err
	}

	return u.userRepo.UpdateUser(userEntity)
}

func (u *UserApp) UpdateMine(
	userEntity *entities.User,
	sessionStruct *method.SessionStruct,
) (*entities.User, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		u.roleRepo,
		u.companyBanchRepo,
		true,
		string(enum.SelfData),
		string(enum.Edit),
	)
	
	if err != nil {
		return nil, err
	}

	// 只能更新名字
	authAggregate.User.UserName = userEntity.UserName

	return u.userRepo.UpdateUser(&authAggregate.User)
}

func (u *UserApp) SaveUser(
	userEntity *entities.User,
	sessionStruct *method.SessionStruct,
) (*entities.User, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		u.roleRepo,
		u.companyBanchRepo,
		true,
		string(enum.EmployeeManage),
		string(enum.Add),
	)
	
	if err != nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeBanchValidation(*(*userEntity).BanchId); err != nil {
		return nil, &err
	}
	if err := authAggregate.CheckScopeRoleValidation((*userEntity).RoleId); err != nil {
		return nil, &err
	}

	userEntity.CompanyId = authAggregate.User.CompanyId

	return u.userRepo.SaveUser(userEntity)
}

func (u *UserApp) DeleteUser(
	userEntity *entities.User,
	sessionStruct *method.SessionStruct,
) (*entities.User, *error) {
	authAggregate, err := aggregates.NewAuthAggregate(
		sessionStruct,
		u.roleRepo,
		u.companyBanchRepo,
		true,
		string(enum.EmployeeManage),
		string(enum.Delete),
	)
	
	if err != nil {
		return nil, err
	}

	if err := authAggregate.CheckScopeBanchValidation(*(*userEntity).BanchId); err != nil {
		return nil, &err
	}
	if err := authAggregate.CheckScopeRoleValidation((*userEntity).RoleId); err != nil {
		return nil, &err
	}

	userEntity.CompanyId = authAggregate.User.CompanyId

	return u.userRepo.DeleteUser(userEntity)
}