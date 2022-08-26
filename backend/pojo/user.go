package pojo

type User struct {
	Name string  `json:"name" form:"id" binding:"required"`
	Banch string `json:"banch" form:"id" binding:"required"`
	// UserLogin
}
type UserLogin struct {
	Acocout string `json:"account"`
	Password string `json:"password"`
}