import companyBanchTypes from './companyBanch'
import roleTypes from './role'
import userTypes from './user'

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
    print = 'print',
    copy = 'copy',
    changeBanch = 'changeBanch'
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

    // api
    interface auth {
        menu: functionItemTable[]
        permission: Record<funcCode, Record<operationCode, permission>>
    }

    interface func {
        functionItem: functionItemTable[]
        operationItem: operationItemTable[]
    }

    interface roleBanchList {
        availableBanch: companyBanchTypes.TABLE[]
        availableRole: roleTypes.TABLE[]
        availableUser: userTypes.TABLE[]
        scopeBanch: Record<funcCode, Record<operationCode, number[]>>
        scopeRole: Record<funcCode, Record<operationCode, number[]>>
        scopeUser: Record<funcCode, Record<operationCode, number[]>>
    }

    interface reducerType {
        auth: auth
        sidebar: boolean
        func: func
        roleBanchList: roleBanchList
    }
}
export default systemTypes
export {
    funcCode,
    operationCode
}
