// Package models  公司部門
// author : http://www.liyang.love
// date : 2023-10-04 16:45
// desc : 公司部門
package Model

import "time"

// CompanyBanch  公司部門。
// 说明:
// 表名:company_banch
// group: CompanyBanch
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.CompanyBanch
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.CompanyBanch
// version:2023-10-04 16:45
type CompanyBanch struct {
    CompanyId   *int         `gorm:"column:companyId;primaryKey" json:"CompanyId"`   //type:*int         comment:公司Id               version:2023-10-04 16:45
    BanchId     *int         `gorm:"column:banchId;primaryKey" json:"BanchId"`   //type:*int         comment:部門id               version:2023-10-04 16:45
    BanchName   string       `gorm:"column:banchName" json:"BanchName"`     //type:string       comment:部門名稱             version:2023-10-04 16:45
    DeleteFlag  string         `gorm:"column:deleteFlag" json:"DeleteFlag"`   //type:CHAR         comment:刪除旗標 ( N, Y )    version:2023-10-04 16:45
    DeleteTime  time.Time   `gorm:"column:deleteTime" json:"DeleteTime"`   //type:*time.Time   comment:刪除時間             version:2023-10-04 16:45
    CreateTime  time.Time   `gorm:"column:createTime" json:"CreateTime"`   //type:*time.Time   comment:創建時間             version:2023-10-04 16:45
    LastModify  time.Time   `gorm:"column:lastModify" json:"LastModify"`   //type:*time.Time   comment:最後更新時間         version:2023-10-04 16:45
}

// TableName 表名:company_banch，公司部門。
// 说明:
func (c CompanyBanch) TableName() string {
	return "company_banch"
}
