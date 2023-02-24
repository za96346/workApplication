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
    BanchName: string // 部門名
    CompanyName: string // 公司名
    CompanyId: number // 公司id
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
    Icon: string
    RestTime: string
    BanchStyleId: number
    Date: string
    UserId: number
}
export interface ShiftState {
    disabledTable?: boolean
    submitAble?: boolean
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
    State: ShiftState
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

export interface performanceType {
    PerformanceId: number
    UserId: number
    Year: number
    Month: number
    BanchId: number
    Goal: string
    Attitude: number
    Efficiency: number
    Professional: number
    Directions: string
    BeLate: number
    DayOffNotOnRule: number
    BanchName: string
    CreateTime: string | Date
    LastModify: string | Date
    UserName: string
    CompanyId: number
}

export interface yearPerformanceType {
    UserId: number
    Year: number
    UserName: string
    Avg: number
}

export interface nullTime {
    Time: string
    Valid: boolean
}

export interface shiftMonthType {
    ShiftId: number // 班表的編號
    UserId: number // 使用者的編號
    BanchStyleId: number // 班表樣式id
    Year: number // 紀錄 年
    Month: number // 紀錄 月
    OnShiftTime: Date// 開始上班時間
    OffShiftTime: Date // 結束上班的時間
    RestTime: string // 休息時間
    PunchIn: nullTime // 上班卡
    PunchOut: nullTime// 下班卡
    CreateTime: Date // 創建的時間
    LastModify: Date // 上次修改的時間
    SpecifyTag: string // 特別的備註
    UserName: string // 使用者名稱
    Permission: number // 權限
    Banch: number // 部門編號
    EmployeeNumber: string // 員工編號
}

export interface shiftTotalType {
    UserId: number
    Year: number
    Month: number
    BanchId: number
    UserName: string
    Permession: number
    EmployeeNumber: string
    ChangeCocunt: number // 換班次數
    OverTimeCount: number // 加班次數
    ForgetPunchCount: number // 忘記打卡次數
    DayOffCount: number // 請假次數
    LateExcusedCount: number // 遲到次數
    TotalHours: number // 總計時數
}
