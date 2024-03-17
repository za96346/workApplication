package aggregates

import "backend/domain/repository"

var roleRepo repository.RoleRepository
var banchRepo repository.CompanyBanchRepository
var userRepo repository.UserRepository

func NewAggregateRepo(roleRepoProp repository.RoleRepository, banchRepoProp repository.CompanyBanchRepository, userRepoProp repository.UserRepository) {
	roleRepo = roleRepoProp
	banchRepo = banchRepoProp
	userRepo = userRepoProp
}