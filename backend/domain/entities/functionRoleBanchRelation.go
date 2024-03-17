package entities
import "time"

type FuncRoleBanchRelation struct {
    FuncCode        string       `gorm:"column:funcCode" json:"FuncCode"`               //type:string       comment:功能代碼( banchManage, shiftedit )         version:2024-00-16 21:06
    ItemCode        string       `gorm:"column:itemCode" json:"ItemCode"`               //type:string       comment:操作代碼(edit, delete...)                  version:2024-00-16 21:06
    HasScopeBanch   string       `gorm:"column:hasScopeBanch" json:"HasScopeBanch"`     //type:string       comment:操作部門範圍 ( [all, self, customize] )    version:2024-00-16 21:06
    HasScopeRole    string       `gorm:"column:hasScopeRole" json:"HasScopeRole"`       //type:string       comment:操作角色範圍 ( [all, self, customize] )    version:2024-00-16 21:06
    HasScopeUser    string       `gorm:"column:hasScopeUser" json:"HasScopeUser"`       //type:string       comment:操作使用者範圍 ( [all, self, customize] )    version:2024-00-16 21:06
    CreateTime      *time.Time   `gorm:"column:createTime" json:"CreateTime"`           //type:*time.Time   comment:創建時間                                   version:2024-00-16 21:06
    LastModify      *time.Time   `gorm:"column:lastModify" json:"LastModify"`           //type:*time.Time   comment:最近修改                                   version:2024-00-16 21:06
}
