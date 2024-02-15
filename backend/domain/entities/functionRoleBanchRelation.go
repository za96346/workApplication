// Package models  功能對應擁有的角色部門關聯
// author : http://www.liyang.love
// date : 2024-00-16 21:06
// desc : 功能對應擁有的角色部門關聯
package entities
import "time"

// FuncRoleBanchRelation  功能對應擁有的角色部門關聯。
// 说明:
// 表名:func_role_banch_relation
// group: FuncRoleBanchRelation
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.FuncRoleBanchRelation
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.FuncRoleBanchRelation
// version:2024-00-16 21:06
type FuncRoleBanchRelation struct {
    FuncCode        string       `gorm:"column:funcCode" json:"FuncCode"`               //type:string       comment:功能代碼( banchManage, shiftedit )         version:2024-00-16 21:06
    ItemCode        string       `gorm:"column:itemCode" json:"ItemCode"`               //type:string       comment:操作代碼(edit, delete...)                  version:2024-00-16 21:06
    HasScopeBanch   string       `gorm:"column:hasScopeBanch" json:"HasScopeBanch"`     //type:string       comment:操作部門範圍 ( [all, self, customize] )    version:2024-00-16 21:06
    HasScopeRole    string       `gorm:"column:hasScopeRole" json:"HasScopeRole"`       //type:string       comment:操作角色範圍 ( [all, self, customize] )    version:2024-00-16 21:06
    CreateTime      *time.Time   `gorm:"column:createTime" json:"CreateTime"`           //type:*time.Time   comment:創建時間                                   version:2024-00-16 21:06
    LastModify      *time.Time   `gorm:"column:lastModify" json:"LastModify"`           //type:*time.Time   comment:最近修改                                   version:2024-00-16 21:06
}

// TableName 表名:func_role_banch_relation，功能對應擁有的角色部門關聯。
// 说明:
func (f *FuncRoleBanchRelation) TableName() string {
	return "func_role_banch_relation"
}
