// Package models  使用者資料
// author : http://www.liyang.love
// date : 2023-10-01 13:06
// desc : 使用者資料
package models

import "time"

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
    Companyid        string      `gorm:"column:primaryKey;companyId" json:"Companyid"`    //type:string      comment:公司id                version:2023-10-01 13:06
    Userid           *int        `gorm:"column:primaryKey;userId" json:"Userid"`          //type:*int        comment:使用者id              version:2023-10-01 13:06
    Roleid           *int        `gorm:"column:roleId" json:"Roleid"`                     //type:*int        comment:使用者套用的角色id    version:2023-10-01 13:06
    Username         string      `gorm:"column:userName" json:"Username"`                 //type:string      comment:使用者名稱            version:2023-10-01 13:06
    Employeenumber   string      `gorm:"column:employeeNumber" json:"Employeenumber"`     //type:string      comment:使用者員工編號        version:2023-10-01 13:06
    Account          string      `gorm:"column:account" json:"Account"`                   //type:string      comment:使用者帳號            version:2023-10-01 13:06
    Password         string      `gorm:"column:password" json:"Password"`                 //type:string      comment:使用者密碼            version:2023-10-01 13:06
    Onworkday        time.Time   `gorm:"column:onWorkDay" json:"Onworkday"`               //type:TIMESTAMP   comment:開始工作日            version:2023-10-01 13:06
    Banchid          *int        `gorm:"column:banchId" json:"Banchid"`                   //type:*int        comment:部門id                version:2023-10-01 13:06
    Createtime       time.Time   `gorm:"column:createTime" json:"Createtime"`             //type:TIMESTAMP   comment:創建時間              version:2023-10-01 13:06
    Lastmodify       time.Time   `gorm:"column:lastModify" json:"Lastmodify"`             //type:TIMESTAMP   comment:最後更新時間          version:2023-10-01 13:06
}

// TableName 表名:user，使用者資料。
// 说明:
func (u User) TableName() string {
	return "user"
}