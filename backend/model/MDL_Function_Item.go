// Package models  功能項目，
// author : http://www.liyang.love
// date : 2023-10-04 12:55
// desc : 功能項目，
package Model

import "time"

// FunctionItem  功能項目，。
// 说明:undefined
// 表名:function_item
// group: FunctionItem
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.FunctionItem
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.FunctionItem
// version:2023-10-04 12:55
type FunctionItem struct {
    FuncCode   string       `gorm:"column:funcCode;primaryKey" json:"FuncCode"`   //type:string       comment:功能代碼(banchManager)    version:2023-10-04 12:55
    FuncName   string       `gorm:"column:funcName" json:"FuncName"`         //type:string       comment:功能名稱                  version:2023-10-04 12:55
    ScopeRoleEnable   string  `gorm:"column:scopeRoleEnable" json:"ScopeRoleEnable"`     //type:CHAR         comment:可編輯角色範圍            version:2023-10-04 13:02
    ScopeBanchEnable  string  `gorm:"column:scopeBanchEnable" json:"ScopeBanchEnable"`   //type:CHAR         comment:可編輯部門範圍            version:2023-10-04 13:02
    Sort             *int         `gorm:"column:sort" json:"Sort"`               //type:*int         comment:排序                 version:2024-00-14 15:12
    CreateTime *time.Time   `gorm:"column:createTime" json:"CreateTime"`     //type:*time.Time   comment:創建時間                  version:2023-10-04 12:55
    LastModify *time.Time   `gorm:"column:lastModify" json:"LastModify"`     //type:*time.Time   comment:最近修改                  version:2023-10-04 12:55
}

// TableName 表名:function_item，功能項目。
// 说明:undefined
func (f *FunctionItem) TableName() string {
	return "function_item"
}