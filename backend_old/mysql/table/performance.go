package table

import (
	"time"

)

type Performance struct {
	PerformanceId int64 `json:"PerformanceId"`
	UserId int64 `json:"UserId"`
	Year int `json:"Year"`
	Month int `json:"Month"`
	BanchId int64 `json:"BanchId"`
	Goal string `json:"Goal"`
	Attitude int `json:"Attitude"`
	Efficiency int `json:"Efficiency"`
	Professional int `json:"Professional"`
	Directions string `json:"Directions"`
	BeLate int `json:"BeLate"`
	DayOffNotOnRule int `json:"DayOffNotOnRule"`
	BanchName string `json:"BanchName"`
	CreateTime time.Time `json:"CreateTime"`//創建的時間
	LastModify time.Time `json:"LastModify"`// 上次修改的時間
}

type PerformanceExtend struct {
	Performance
	UserName string `json:"UserName"` // 名字
	CompanyId int64 `json:"CompanyId"` // 公司編號
}

type YearPerformance struct {
	UserId int64
	Year int
	UserName string
	Avg float32
}