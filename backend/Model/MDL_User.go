// Package models  使用者資料
// author : http://www.liyang.love
// date : 2023-10-01 13:06
// desc : 使用者資料
package Model

import (
	"time"
)

// User  使用者資料。
// 说明:
// 表名:user
// group: User
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.User
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.User
// version:2023-10-01 13:06
type User struct {
    CompanyId        int       `gorm:"column:companyId;primaryKey" json:"CompanyId"`    //type:string       comment:公司id                version:2023-10-02 14:15
    UserId           int         `gorm:"column:userId;primaryKey" json:"UserId"`          //type:*int         comment:使用者id              version:2023-10-02 14:15
    RoleId           int         `gorm:"column:roleId" json:"RoleId" binding:"required"`                     //type:*int         comment:使用者套用的角色id    version:2023-10-02 14:15
    BanchId          *int         `gorm:"column:banchId" json:"BanchId"`                   //type:*int         comment:部門id                version:2023-10-02 14:15
    UserName         string       `gorm:"column:userName" json:"UserName" binding:"required"`                 //type:string       comment:使用者名稱            version:2023-10-02 14:15
    EmployeeNumber   string       `gorm:"column:employeeNumber" json:"EmployeeNumber" binding:"required"`     //type:string       comment:使用者員工編號        version:2023-10-02 14:15
    Account          string       `gorm:"column:account" json:"Account" binding:"required"`                   //type:string       comment:使用者帳號            version:2023-10-02 14:15
    Password         string       `gorm:"column:password" json:"Password"`                 //type:string       comment:使用者密碼            version:2023-10-02 14:15
    OnWorkDay        time.Time   `gorm:"column:onWorkDay" json:"OnWorkDay" binding:"required"`               //type:*time.Time   comment:開始工作日            version:2023-10-02 14:15
    Sort             *int         `gorm:"column:sort" json:"Sort"`               //type:*int         comment:排序                 version:2024-00-14 15:12
    QuitFlag         string         `gorm:"column:quitFlag" json:"QuitFlag"`                 //type:CHAR         comment:離職旗標              version:2024-00-06 14:30
    DeleteFlag       string         `gorm:"column:deleteFlag" json:"DeleteFlag"`             //type:CHAR         comment:刪除旗標 ( N, Y )     version:2023-10-02 19:31
    DeleteTime       *time.Time   `gorm:"column:deleteTime" json:"DeleteTime"`             //type:*time.Time   comment:刪除時間              version:2023-10-02 19:31
    CreateTime       *time.Time   `gorm:"column:createTime" json:"CreateTime"`             //type:*time.Time   comment:創建時間              version:2023-10-02 14:15
    LastModify       *time.Time   `gorm:"column:lastModify" json:"LastModify"`             //type:*time.Time   comment:最後更新時間          version:2023-10-02 14:15
}

// TableName 表名:user，使用者資料。
// 说明:
func (u *User) TableName() string {
	return "user"
}

// 拿取新的 user id ( max count + 1 )
func (u *User) GetNewUserID(companyId ...int) int {
    var MaxCount int64

    validateField := (*u).CompanyId

    if len(companyId) > 0 {
        validateField = companyId[0]
    }

    DB.Model(&User{}).
        Where("companyId = ?", validateField).
        Select("max(userId)").
        Row().
        Scan(&MaxCount)
    
    (*u).UserId = int(MaxCount) + 1
    (*u).CompanyId = validateField

    return int(MaxCount) + 1
}

// 帳號是否重複
func (u *User) IsAccountDuplicated(account ...string) bool {
    var MaxCount int64

    validateField := (*u).Account

    if len(account) > 0 {
        validateField = account[0]
    }

    DB.Model(&User{}).
        Where("account = ?", validateField).
        Count(&MaxCount)

    return MaxCount > 0
}

// 是否離職
func (u *User) IsQuitWorking() bool {
    return  (*u).QuitFlag == "Y"
}