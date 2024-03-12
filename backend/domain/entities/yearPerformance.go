package entities

type YearPerformance struct {
	Year int `gorm:"column:year" json:"Year"`
	UserName string  `gorm:"column:userName" json:"UserName"`
	UserId int `gorm:"column:userId" json:"UserId"`
	Score float32 `gorm:"column:score" json:"Score"`
}
