package entities

import "time"

type Role struct {
    CompanyId   int         `gorm:"column:companyId;primaryKey" json:"CompanyId"`   //type:*int         comment:公司id               version:2023-10-02 14:18
    RoleId      int         `gorm:"column:roleId;primaryKey" json:"RoleId"`   //type:*int         comment:角色id               version:2023-10-02 14:18
    RoleName    string       `gorm:"column:roleName" json:"RoleName" binding:"required"`       //type:string       comment:角色名稱             version:2023-10-02 14:18
    StopFlag    string         `gorm:"column:stopFlag" json:"StopFlag"`       //type:CHAR         comment:停用旗標 ( N, Y )    version:2023-10-02 14:18
    DeleteFlag  string       `gorm:"column:deleteFlag" json:"DeleteFlag"`   //type:string       comment:刪除旗標 ( N, Y )    version:2023-10-02 19:26
    DeleteTime  *time.Time   `gorm:"column:deleteTime" json:"DeleteTime"`   //type:*time.Time   comment:刪除時間             version:2023-10-02 19:28
    Sort        *int         `gorm:"column:sort" json:"Sort"`               //type:*int         comment:排序                 version:2024-00-14 15:12
    CreateTime  *time.Time   `gorm:"column:createTime" json:"CreateTime"`   //type:*time.Time   comment:創建時間             version:2023-10-02 14:18
    LastModify  *time.Time   `gorm:"column:lastModify" json:"LastModify"`   //type:*time.Time   comment:最近修改             version:2023-10-02 14:18
}
