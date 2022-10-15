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
    OnWorkDay: Date // 到職日
    Banch: number // 部門
    Permession: string // 權限 admin 100 ,, personal 2 ,, manager 1
    WorkState: string // 工作狀態 (到職on or 離職off)
    MonthSalary: number // 月薪
    PartTimeSalary: number // 時薪
    Captcha: number // 驗證碼
}

export interface EmpManagerCellType {
    key: string
    empIdx: string | number
    name: string
    onWorkDay: string
    workState: string
    banch: number
    permession: string
}

export interface BanchType {
    Id: number
    CompanyId: string
    BanchName: string
    BanchShiftStyle: string
    CreateTime: string | Date
    LastModify: string | Date
}

export interface SelfDataType {
    UserId: number // 使用者的編號
    CompanyCode: string // 公司碼
    Account: string // 帳號
    Password: string // 密碼
    OnWorkDay: string | Date // 到職日
    Banch: number // 部門
    Permession: number // 權限  (0 admin , 1 manager, 2 personal)
    WorkState: string // 工作狀態 (到職on or 離職off)
    MonthSalary: number // 月薪
    PartTimeSalary: number // 時薪
    CreateTime: Date | string // 創建的時間
    LastModify: Date | string // 上次修改的時間
}

export interface UserType {
    UserId: number // 使用者id
    CompanyCode: string // 公司碼
    UserName: string // 員工姓名
    EmployeeNumber: string // 員工編號
    OnWorkDay: string | Date // 到職日
    Banch: number // 部門編號
    Permession: number // 權限
    WorkState: string // 工作狀態
}

export interface ResType<T> extends ResMessage {
    data?: T
    status: boolean
}
