package entities

import "time"

type CompanyBanch struct {
    CompanyId   int         `gorm:"column:companyId;primaryKey" json:"CompanyId"`   //type:*int         comment:公司Id               version:2023-10-04 16:45
    BanchId     int         `gorm:"column:banchId;primaryKey" json:"BanchId"`   //type:*int         comment:部門id               version:2023-10-04 16:45
    BanchName   string       `gorm:"column:banchName" json:"BanchName" binding:"required"`     //type:string       comment:部門名稱             version:2023-10-04 16:45
    Sort        *int         `gorm:"column:sort" json:"Sort"`               //type:*int         comment:排序                 version:2024-00-14 15:12
    DeleteFlag  string         `gorm:"column:deleteFlag" json:"DeleteFlag"`   //type:CHAR         comment:刪除旗標 ( N, Y )    version:2023-10-04 16:45
    DeleteTime  *time.Time   `gorm:"column:deleteTime" json:"DeleteTime"`   //type:*time.Time   comment:刪除時間             version:2023-10-04 16:45
    CreateTime  *time.Time   `gorm:"column:createTime" json:"CreateTime"`   //type:*time.Time   comment:創建時間             version:2023-10-04 16:45
    LastModify  *time.Time   `gorm:"column:lastModify" json:"LastModify"`   //type:*time.Time   comment:最後更新時間         version:2023-10-04 16:45
}


func (c *CompanyBanch) TableName() string {
	return "company_banch"
}

// 拿取新的 banch id ( max count + 1 )
func (cb *CompanyBanch) GetNewBanchID(companyId int) int {
    var MaxCount int64
    DB.Model(&CompanyBanch{}).
        Where("companyId = ?", companyId).
        Select("max(banchId)").
        Row().
        Scan(&MaxCount)
    
    (*cb).BanchId = int(MaxCount) + 1
    (*cb).CompanyId = companyId

    return int(MaxCount) + 1
}

// 查詢是否有重複banch name
func (b *CompanyBanch) IsBanchNameDuplicated() bool {
    var MaxCount int64
    DB.Model(&CompanyBanch{}).
        Where("companyId = ?", (*b).CompanyId).
        Where("banchName = ?", (*b).BanchName).
        Where("banchId != ?", (*b).BanchId).
        Where("deleteFlag = ?", "N").
        Count(&MaxCount)

    return int(MaxCount) > 0
}
