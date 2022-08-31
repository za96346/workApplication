package typeinterface

import (
	"sync"
)
var userTableMux sync.RWMutex
var userTableInstance *UserTable

type UserTable struct {
	UserId string
	CompanyCode string
	Account string
	Password string
	OnWorkDay string
	Banch string
	Permession string
	Work_state string
	CreateTime string
	LastModify string
	MonthSalary float64
	PartTimeSalary float64
}
type Response struct {
	Message string `binding:"required"`
	Data string	`binding:"required"`
	Status bool `binding:"required"`
}

func UserTableSingleton() *UserTable{
	if userTableInstance == nil {
		userTableMux.Lock()
		defer userTableMux.Unlock()
		if userTableInstance == nil {
			userTableInstance = &UserTable{}
		}
	}
	return userTableInstance
}

func NewResponse(message string, data string, status bool) *Response {
	res := new(Response)
	res.Message = message
	res.Data = data
	res.Status = status
	return res
}