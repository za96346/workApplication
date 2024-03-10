package entities

import "time"

type FunctionItem struct {
    FuncCode   string       `gorm:"column:funcCode;primaryKey" json:"FuncCode"`   //type:string       comment:功能代碼(banchManager)    version:2023-10-04 12:55
    FuncName   string       `gorm:"column:funcName" json:"FuncName"`         //type:string       comment:功能名稱                  version:2023-10-04 12:55
    Sort             *int         `gorm:"column:sort" json:"Sort"`               //type:*int         comment:排序                 version:2024-00-14 15:12
    CreateTime *time.Time   `gorm:"column:createTime" json:"CreateTime"`     //type:*time.Time   comment:創建時間                  version:2023-10-04 12:55
    LastModify *time.Time   `gorm:"column:lastModify" json:"LastModify"`     //type:*time.Time   comment:最近修改                  version:2023-10-04 12:55
}