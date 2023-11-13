enum funcCode {
    banchManage = 'banchManage',
    companyData = 'companyData',
    employeeManage = 'employeeManage',
    performance = 'performance',
    roleManage = 'roleManage',
    selfData = 'selfData',
    shift = 'shift',
    shiftSetting = 'shiftSetting',
    yearPerformance = 'yearPerformance'
}
declare namespace systemTypes {
    interface functionItem {
        CreateTime: string
        FuncName: string
        LastModify: string
        ScopeBanchEnable: Flag
        ScopeRoleEnable: Flag
        funcCode: funcCode
    }
    type operationItem = 'add' | 'edit' | 'delete' | 'inquire'
    interface permission {
        scopeBanch: 'self' | 'all' | number[]
        scopeRole: 'self' | 'all' | number[]
    }
    interface auth {
        menu: functionItem[]
        permission: Record<string, Record<string, permission>>
    }
}
export default systemTypes
export {
    funcCode
}
