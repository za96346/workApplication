package table

import (
	"time"

)

// log
type Log struct {
	LogId int64
	UserId int64
	UserName string
	CompanyId int64
	CompanyCode string
	Permession int
	Routes string
	Ip string
	Params string
	Msg string
	CreateTime time.Time
	LastModify time.Time
}

