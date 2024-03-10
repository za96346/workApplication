package aggregates

import "backend/domain/repository"

var roleRepo repository.RoleRepository
var banchRepo repository.CompanyBanchRepository

func NewAggregateRepo(roleRepoProp repository.RoleRepository, banchRepoProp repository.CompanyBanchRepository) {
	roleRepo = roleRepoProp
	banchRepo = banchRepoProp
}