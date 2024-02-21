package dtos

import "backend/domain/entities"

type PerformanceDetailDto struct {
	entities.Performance
	BanchName string  `gorm:"column:banchName" json:"BanchName"`
	UserName string  `gorm:"column:userName" json:"UserName"`
}