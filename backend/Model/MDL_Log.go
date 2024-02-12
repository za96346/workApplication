// Package models  Log紀錄
// author : http://www.liyang.love
// date : 2023-10-24 12:55
// desc : Log紀錄
package Model

import "time"

// Log  Log紀錄。
// 说明:
// 表名:log
// group: Log
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.Log
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.Log
// version:2023-10-24 12:55
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

// TableName 表名:log，Log紀錄。
// 说明:
func (l *Log) TableName() string {
	return "log"
}
// 拿取新的 log id ( max count + 1 )
func (l *Log) GetNewLogId(companyId int) int {
    var MaxCount int64
    DB.Model(&Log{}).
        Where("companyId = ?", companyId).
        Count(&MaxCount)
    
    (*l).LogId = int(MaxCount) + 1
    (*l).CompanyId = companyId

    return int(MaxCount) + 1
}