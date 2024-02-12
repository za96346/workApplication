// Package models  操作代碼，
// author : http://www.liyang.love
// date : 2023-10-13 21:31
// desc : 操作代碼，
package Model

import "time"

// OperationItem  操作代碼，。
// 说明:undefined
// 表名:operation_item
// group: OperationItem
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.OperationItem
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.OperationItem
// version:2023-10-13 21:31
type OperationItem struct {
    OperationCode   string       `gorm:"column:operationCode;primaryKey" json:"OperationCode"`   //type:string       comment:操作代碼(edit)    version:2023-10-13 21:31
    OperationName   string       `gorm:"column:operationName" json:"OperationName"`     //type:string       comment:操作名稱          version:2023-10-13 21:31
    Sort            *int         `gorm:"column:sort" json:"Sort"`               //type:*int         comment:排序                 version:2024-00-14 15:12
    CreateTime      *time.Time   `gorm:"column:createTime" json:"CreateTime"`           //type:*time.Time   comment:創建時間          version:2023-10-13 21:31
    LastModify      *time.Time   `gorm:"column:lastModify" json:"LastModify"`           //type:*time.Time   comment:最近修改          version:2023-10-13 21:31
}

// TableName 表名:operation_item，操作代碼。
// 说明:undefined
func (o *OperationItem) TableName() string {
	return "operation_item"
}