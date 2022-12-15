import { ReactNode } from 'react'

export interface languageType {
    a: string
}

export interface ResMessage {
    message: string
}

export interface LoginType {
    Account: string
    Password: string
}

export interface GetEmailCaptcha {
    Email: string
}

export interface RegisterType {
    CompanyCode: string // 公司碼
    Account: string // 帳號(email)
    Password: string // 密碼
    PasswordConfirm: string // 再次確認密碼
    Captcha: number // 驗證碼
}

export interface EmpManagerCellType {
    key: string
    EmployeeNumber: string | number
    UserName: string
    OnWorkDay: string
    WorkState: 'on' | 'off'
    Banch: string
    Permession: string
}

export interface BanchType {
    Id: number
    CompanyId: string
    BanchName: string
    BanchShiftStyle: string
    CreateTime: string | Date
    LastModify: string | Date
}

export interface UserType {
    UserId: number // 使用者id
    CompanyCode: string // 公司碼
    UserName: string // 員工姓名
    EmployeeNumber: string // 員工編號
    OnWorkDay: string | Date // 到職日
    Banch: number // 部門編號
    Permession: number // 權限(0 admin , 1 manager, 2 personal)
    WorkState: string // 工作狀態
}

export interface SelfDataType extends UserType {
    Account: string // 帳號
    Password: string // 密碼
    MonthSalary: number // 月薪
    PartTimeSalary: number // 時薪
    CreateTime: Date | string // 創建的時間
    LastModify: Date | string // 上次修改的時間
}

export interface ResType<T> extends ResMessage {
    data?: T
    status: boolean
}

export interface BanchStyleType {
    StyleId?: number // 樣式id
    BanchId?: number // 部門id
    Icon: string // 圖標
    RestTime: string // 休息時間
    TimeRangeName: string // 時段名稱
    OnShiftTime: string // 開始上班時段
    OffShiftTime: string // 結束上班時段
    CreateTime?: number // 創建時間
    LastModify?: number // 上次修改時間
}

export interface BanchRuleType {
    RuleId?: number // 規則Id
    BanchId?: number // 部門id
    MinPeople: number // 最少上班人數
    MaxPeople: number // 最多上班人數
    WeekDay: number // 星期幾 1, 2, 3, 4, 5, 6, 7
    WeekType: number // 平假日
    OnShiftTime: string // "18:00:00"
    OffShiftTime: string// "18:00:00"
    CreateTime?: Date | string // 創建時間
    LastModify?: Date | string // 上次修改時間
}

export interface CompanyType {
    CompanyId?: number
    CompanyCode?: string
    CompanyName: string
    CompanyLocation: string
    CompanyPhoneNumber: string
    BossId: string
    SettlementDate: number
    TermStart: string | Date
    TermEnd: string | Date
    CreateTime?: string | Date
    LastModify?: string | Date
}

export interface ShiftSettingListType {
    title: string
    icons: string
    time: ReactNode
    id: string
}

export interface BanchRuleListType {
    title: string
    id: string
    time: ReactNode
    weekType: number
}

export interface OnlineUserType {
    Color: string
    Online: number
    Pic: string
    Position: number
    UserId: number
    UserName: string
}

export interface ShiftEditType {
    OffShiftTime: string
    OnShiftTime: string
    RestTime: string
    BanchStyleId: number
    Date: string
    UserId: number
}

export interface ShiftSocketType {
    BanchId: number
    EditUser: UserType[]
    OnlineUser: OnlineUserType[]
    ShiftData: ShiftEditType[]
    BanchStyle: BanchStyleType[]
    WeekendSetting: WeekendSettingType[]
    Status: number // 1 開放編輯、 2 主管審核 3 確認發布 ,
    StartDay: string
    EndDay: string
}

export interface WeekendSettingType {
    WeekendId: number
    CompanyId: number
    Date: string
    CreateTime: string | Date
    LastModify: string | Date
}

export interface WaitReplyType {
    WaitId: number
    UserId: number
    UserName: string
    CompanyId: number
    SpecifyTag: string
    IsAccept: number
    CreateTime: string | Date
    LastModify: string | Date
}

export interface workTimeType {
    WorkTimeId?: number
    UserId: number
    Year: number
    Month: number
    WorkHours: number
    TimeOff: number
    UsePaidVocation: number
    CreateTime?: string | Date
    LastModify?: string | Date
}
