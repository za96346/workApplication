package entities

import "time"

type Log struct {
    CompanyId   int      `gorm:"column:companyId" json:"CompanyId"`     //type:BIGINT      comment:    version:2023-10-24 12:55
    LogId       int      `gorm:"column:logId" json:"LogId"`             //type:BIGINT      comment:    version:2023-10-24 12:55
    UserId      int      `gorm:"column:userId" json:"UserId"`           //type:BIGINT      comment:    version:2023-10-24 12:55
    Routes      string      `gorm:"column:routes" json:"Routes"`           //type:string      comment:    version:2023-10-24 12:55
    Ip          string      `gorm:"column:ip" json:"Ip"`                   //type:string      comment:    version:2023-10-24 12:55
    Params      *string      `gorm:"column:params" json:"Params"`           //type:string      comment:    version:2023-10-24 12:55
    Msg         *string      `gorm:"column:msg" json:"Msg"`                 //type:string      comment:    version:2023-10-24 12:55
    CreateTime  *time.Time   `gorm:"column:createTime" json:"CreateTime"`   //type:TIMESTAMP   comment:    version:2023-10-24 12:55
    LastModify  *time.Time   `gorm:"column:lastModify" json:"LastModify"`   //type:TIMESTAMP   comment:    version:2023-10-24 12:55
}
