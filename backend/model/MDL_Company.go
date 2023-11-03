// Package models  公司
// author : http://www.liyang.love
// date : 2023-10-02 14:21
// desc : 公司
package Model

import "time"

// Company  公司。
// 说明:
// 表名:company
// group: Company
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.Company
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.Company
// version:2023-10-02 14:21
type Company struct {
    CompanyId            int         `gorm:"column:companyId;primaryKey" json:"CompanyId"`            //type:*int         comment:公司id          version:2023-10-02 14:21
    CompanyCode          string       `gorm:"column:companyCode" json:"CompanyCode"`                   //type:string       comment:公司代碼        version:2023-10-02 14:21
    CompanyName          string       `gorm:"column:companyName" json:"CompanyName"`                   //type:string       comment:公司名稱        version:2023-10-02 14:21
    CompanyLocation      string       `gorm:"column:companyLocation" json:"CompanyLocation"`           //type:string       comment:公司位置        version:2023-10-02 14:21
    CompanyPhonenumber   string       `gorm:"column:companyPhoneNumber" json:"CompanyPhonenumber"`     //type:string       comment:公司電話        version:2023-10-02 14:21
    BossId               int         `gorm:"column:bossId" json:"BossId"`                             //type:*int         comment:負責人userId    version:2023-10-02 14:21
    CreateTime           time.Time   `gorm:"column:createTime" json:"CreateTime"`                     //type:*time.Time   comment:創建時間        version:2023-10-02 14:21
    LastModify           time.Time   `gorm:"column:lastModify" json:"LastModify"`                     //type:*time.Time   comment:最後更新時間    version:2023-10-02 14:21
}

// TableName 表名:company，公司。
// 说明:
func (c Company) TableName() string {
	return "company"
}
