
// Package models  離職使用者，因為一個員工有可能離職在復職，故設計此table，來記錄斷斷續續的年資。
// author : http://www.liyang.love
// date : 2024-00-02 21:56
// desc : 離職使用者，因為一個員工有可能離職在復職，故設計此table，來記錄斷斷續續的年資。
package entities

import (
    "time"
)

// QuitWorkUser  離職使用者，因為一個員工有可能離職在復職，故設計此table，來記錄斷斷續續的年資。。
// 说明:因為一個員工有可能離職在復職，故設計此table，來記錄斷斷續續的年資。
// 表名:quit_work_user
// group: QuitWorkUser
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.QuitWorkUser
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.QuitWorkUser
// version:2024-00-02 21:56
type QuitWorkUser struct {
    CompanyId      *int        `gorm:"column:companyId;primaryKey" json:"CompanyId"`   //type:*int        comment:公司id            version:2024-00-02 21:56
    UserId         *int        `gorm:"column:userId;primaryKey" json:"UserId"`      //type:*int        comment:使用者id          version:2024-00-02 21:56
    QuitId         *int        `gorm:"column:quitId" json:"QuitId"`                 //type:*int        comment:離職id            version:2024-00-02 21:56
    StartTime      *time.Time   `gorm:"column:startTime" json:"StartTime"`           //type:TIMESTAMP   comment:開始離職日        version:2024-00-02 21:56
    EndTime        *time.Time   `gorm:"column:endTime" json:"EndTime"`               //type:TIMESTAMP   comment:結束離職日        version:2024-00-02 21:56
    IsDuringQuit   string        `gorm:"column:isDuringQuit" json:"IsDuringQuit"`     //type:CHAR        comment:是否在離職期間    version:2024-00-02 21:56
    Sort           *int         `gorm:"column:sort" json:"Sort"`               //type:*int         comment:排序                 version:2024-00-14 15:12
    CreateTime     *time.Time   `gorm:"column:createTime" json:"CreateTime"`         //type:TIMESTAMP   comment:創建時間          version:2024-00-02 21:56
    LastModify     *time.Time   `gorm:"column:lastModify" json:"LastModify"`         //type:TIMESTAMP   comment:最後更新時間      version:2024-00-02 21:56
}

// TableName 表名:quit_work_user，離職使用者。
// 说明:因為一個員工有可能離職在復職，故設計此table，來記錄斷斷續續的年資。
func (q *QuitWorkUser) TableName() string {
	return "quit_work_user"
}
