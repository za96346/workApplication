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
export interface LoginResponse {
    token: string
    message: string
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
    banch: string
    permession: string
}
