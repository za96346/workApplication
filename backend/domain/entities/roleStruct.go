package entities

import (
    "time"
)

type RoleStruct struct {
    CompanyId   int         `gorm:"column:companyId;primaryKey" json:"CompanyId"`   //type:*int         comment:公司id                                version:2023-10-02 15:48
    RoleId      int         `gorm:"column:roleId;primaryKey" json:"RoleId"`   //type:*int         comment:角色id                                version:2023-10-02 15:48
    FuncCode    string       `gorm:"column:funcCode;primaryKey" json:"FuncCode"`   //type:string       comment:功能代碼( banchManage, shiftedit )    version:2023-10-02 15:48
    ItemCode    string       `gorm:"column:itemCode;primaryKey" json:"ItemCode"`   //type:string       comment:操作代碼(edit, delete...)             version:2023-10-02 15:48
    ScopeBanch  string       `gorm:"column:scopeBanch" json:"ScopeBanch"`   //type:string       comment:可操作部門範圍 ( 部門ID[] )           version:2023-10-03 23:02
    ScopeRole   string       `gorm:"column:scopeRole" json:"ScopeRole"`     //type:string       comment:可操作角色範圍 ( 角色ID[] )           version:2023-10-02 15:48
    CreateTime  *time.Time   `gorm:"column:createTime" json:"CreateTime"`   //type:*time.Time   comment:創建時間                              version:2023-10-02 15:48
    LastModify  *time.Time   `gorm:"column:lastModify" json:"LastModify"`   //type:*time.Time   comment:最近修改                              version:2023-10-02 15:48
}
