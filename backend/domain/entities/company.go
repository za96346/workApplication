package entities

import "time"

type Company struct {
    CompanyId            int         `gorm:"column:companyId;primaryKey" json:"CompanyId"`            //type:*int         comment:公司id          version:2023-10-02 14:21
    CompanyCode          string       `gorm:"column:companyCode" json:"CompanyCode" binding:"required"`                   //type:string       comment:公司代碼        version:2023-10-02 14:21
    CompanyName          string       `gorm:"column:companyName" json:"CompanyName" binding:"required"`                   //type:string       comment:公司名稱        version:2023-10-02 14:21
    CompanyLocation      string       `gorm:"column:companyLocation" json:"CompanyLocation" binding:"required"`           //type:string       comment:公司位置        version:2023-10-02 14:21
    CompanyPhonenumber   string       `gorm:"column:companyPhoneNumber" json:"companyPhoneNumber" binding:"required"`     //type:string       comment:公司電話        version:2023-10-02 14:21
    BossId               int         `gorm:"column:bossId" json:"BossId" binding:"required"`                             //type:*int         comment:負責人userId    version:2023-10-02 14:21
    CreateTime           *time.Time   `gorm:"column:createTime" json:"CreateTime"`                     //type:*time.Time   comment:創建時間        version:2023-10-02 14:21
    LastModify           *time.Time   `gorm:"column:lastModify" json:"LastModify"`                     //type:*time.Time   comment:最後更新時間    version:2023-10-02 14:21
}