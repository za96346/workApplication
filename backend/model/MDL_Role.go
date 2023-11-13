// Package models  角色表
// author : http://www.liyang.love
// date : 2023-10-02 14:10
// desc : 角色表
package Model

import "time"

// Role  角色表。
// 说明:
// 表名:role
// group: Role
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.Role
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.Role
// version:2023-10-02 14:10
type Role struct {
    CompanyId   int         `gorm:"column:companyId;primaryKey" json:"CompanyId"`   //type:*int         comment:公司id               version:2023-10-02 14:18
    RoleId      int         `gorm:"column:roleId;primaryKey" json:"RoleId"`   //type:*int         comment:角色id               version:2023-10-02 14:18
    RoleName    string       `gorm:"column:roleName" json:"RoleName" binding:"required"`       //type:string       comment:角色名稱             version:2023-10-02 14:18
    StopFlag    string         `gorm:"column:stopFlag" json:"StopFlag"`       //type:CHAR         comment:停用旗標 ( N, Y )    version:2023-10-02 14:18
    DeleteFlag  string       `gorm:"column:deleteFlag" json:"DeleteFlag"`   //type:string       comment:刪除旗標 ( N, Y )    version:2023-10-02 19:26
    DeleteTime  *time.Time   `gorm:"column:deleteTime" json:"DeleteTime"`   //type:*time.Time   comment:刪除時間             version:2023-10-02 19:28
    CreateTime  *time.Time   `gorm:"column:createTime" json:"CreateTime"`   //type:*time.Time   comment:創建時間             version:2023-10-02 14:18
    LastModify  *time.Time   `gorm:"column:lastModify" json:"LastModify"`   //type:*time.Time   comment:最近修改             version:2023-10-02 14:18
}

// TableName 表名:role，角色表。
// 说明:
func (r *Role) TableName() string {
	return "role"
}

// 拿取新的 role id ( max count + 1 )
func (r *Role) GetNewRoleID(companyId int) int {
    var MaxCount int64
    DB.Model(&Role{}).
        Where("companyId = ?", companyId).
        Count(&MaxCount)
    
    (*r).RoleId = int(MaxCount) + 1
    (*r).CompanyId = companyId

    return int(MaxCount) + 1
}