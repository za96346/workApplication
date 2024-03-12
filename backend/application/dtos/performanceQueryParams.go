package dtos

type PerformanceQueryParams struct {
	BanchId int `json:"BanchId"`
	RoleId int `json:"RoleId"`
	UserId int `json:"UserId"`
	UserName string `json:"UserName"`
	StartDate string `json:"StartDate"`
	EndDate string `json:"EndDate"`
	StartYear string `json:"StartYear"`
	EndYear string `json:"EndYear"`
}