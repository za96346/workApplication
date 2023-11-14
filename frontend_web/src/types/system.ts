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
enum operationCode {
    add = 'add',
    edit = 'edit',
    delete = 'delete',
    inquire = 'inquire',
    print = 'print'
}
declare namespace systemTypes {
    interface functionItemTable {
        CreateTime: string
        FuncName: string
        LastModify: string
        ScopeBanchEnable: Flag
        ScopeRoleEnable: Flag
        funcCode: funcCode
    }
    interface operationItemTable {
        OperationCode: operationCode
        OperationName: string
        CreateTime: string
        LastModify: string
    }
    interface permission {
        scopeBanch: 'self' | 'all' | number[]
        scopeRole: 'self' | 'all' | number[]
    }
    interface auth {
        menu: functionItemTable[]
        permission: Record<funcCode, Record<operationCode, permission>>
    }

    interface func {
        functionItem: functionItemTable[]
        operationItem: operationItemTable[]
    }
}
export default systemTypes
export {
    funcCode,
    operationCode
}
