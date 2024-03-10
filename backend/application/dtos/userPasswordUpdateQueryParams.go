package dtos

type UserPasswordUpdateQueryParams struct {
	OldPassword string `json:"OldPassword" binding:"required"`
	NewPassword string `json:"NewPassword" binding:"required"`
	NewPasswordAgain string `json:"NewPasswordAgain" binding:"required"`
	UserId int `gorm:"column:userId;primaryKey" json:"UserId" binding:"required"`
}