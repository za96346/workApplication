package entities

import (
    "time"
)

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
