// Package models  績效評核
// author : http://www.liyang.love
// date : 2023-10-08 14:35
// desc : 績效評核
package Model

import "time"

// Performance  績效評核。
// 说明:
// 表名:performance
// group: Performance
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.Performance
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.Performance
// version:2023-10-08 14:35
type Performance struct {
    CompanyId       int         `gorm:"column:companyId;primaryKey" json:"CompanyId"`      //type:*int         comment:公司id          version:2023-10-08 14:35
    UserId          int         `gorm:"column:userId;primaryKey" json:"UserId" binding:"required"`            //type:*int         comment:使用者id        version:2023-10-08 14:35
    PerformanceId   int         `gorm:"column:performanceId;primaryKey" json:"PerformanceId"`   //type:*int         comment:績效id          version:2023-10-08 14:35
    Year            int         `gorm:"column:year;primaryKey" json:"Year" binding:"required"`                //type:*int         comment:年分            version:2023-10-08 14:35
    Month           int         `gorm:"column:month;primaryKey" json:"Month" binding:"required"`              //type:*int         comment:月份            version:2023-10-08 14:35
    BanchId         int         `gorm:"column:banchId" json:"BanchId"`                     //type:*int         comment:部門id          version:2023-10-08 14:35
    Goal            string       `gorm:"column:goal" json:"Goal" binding:"required"`                           //type:string       comment:績效目標        version:2023-10-08 14:35
    Attitude        int         `gorm:"column:attitude" json:"Attitude" binding:"required"`                   //type:*int         comment:態度分數        version:2023-10-08 14:35
    Efficiency      int         `gorm:"column:efficiency" json:"Efficiency" binding:"required"`               //type:*int         comment:效率分數        version:2023-10-08 14:35
    Professional    int         `gorm:"column:professional" json:"Professional" binding:"required"`           //type:*int         comment:專業分數        version:2023-10-08 14:35
    Directions      string       `gorm:"column:directions" json:"Directions"`               //type:string       comment:                version:2023-10-08 14:35
    BeLate          int         `gorm:"column:beLate" json:"BeLate" binding:"required"`                       //type:*int         comment:遲到            version:2023-10-08 14:35
    DayOffNotOnRule int         `gorm:"column:dayOffNotOnRule" json:"DayOffNotOnRule" binding:"required"`     //type:*int         comment:未依規定請假    version:2023-10-08 14:35
    DeleteFlag      string         `gorm:"column:deleteFlag" json:"DeleteFlag"`               //type:CHAR         comment:刪除旗標 ( N, Y )    version:2023-10-08 16:10
    DeleteTime      *time.Time   `gorm:"column:deleteTime" json:"DeleteTime"`               //type:*time.Time   comment:刪除時間             version:2023-10-08 16:10
    CreateTime      *time.Time   `gorm:"column:createTime" json:"CreateTime"`               //type:*time.Time   comment:創建時間        version:2023-10-08 14:35
    LastModify      *time.Time   `gorm:"column:lastModify" json:"LastModify"`               //type:*time.Time   comment:最後更新時間    version:2023-10-08 14:35
}

// TableName 表名:performance，績效評核。
// 说明:
func (p *Performance) TableName() string {
	return "performance"
}


// 拿取新的 performance id ( max count + 1 )
func (p *Performance) GetNewPerformanceID(companyId int) int {
    var MaxCount int64
    DB.Model(&Performance{}).
        Where("companyId = ?", companyId).
        Count(&MaxCount)
    
    (*p).PerformanceId = int(MaxCount) + 1
    (*p).CompanyId = companyId

    return int(MaxCount) + 1
}