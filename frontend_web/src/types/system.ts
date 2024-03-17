import companyBanchTypes from './companyBanch'
import roleTypes from './role'
import userTypes from './user'

enum FuncCodeEnum {
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
enum OperationCodeEnum {
    add = 'add',
    edit = 'edit',
    delete = 'delete',
    inquire = 'inquire',
    print = 'print',
    copy = 'copy',
    changeBanch = 'changeBanch'
}
enum ScopeEnum {
    all = 'all',
    self = 'self',
    customize = 'customize'
}

declare namespace systemTypes {
    interface functionItemTable {
        CreateTime: string
        FuncName: string
        LastModify: string
        ScopeBanchEnable: Flag
        ScopeRoleEnable: Flag
        FuncCode: FuncCodeEnum
        Sort: number
    }
    interface operationItemTable {
        OperationCode: OperationCodeEnum
        OperationName: string
        Sort: number
        CreateTime: string
        LastModify: string
    }
    interface functionRoleBanchRelationTable {
        FuncCode: string
        ItemCode: string
        HasScopeBanch: string
        HasScopeRole: string
        CreateTime: string
        LastModify: string
    }
    interface permission {
        scopeBanch: 'self' | 'all' | number[]
        scopeRole: 'self' | 'all' | number[]
        scopeUser: 'self' | 'all' | number[]
    }

    interface scope {
        scopeBanch: ScopeEnum[]
        scopeRole: ScopeEnum[]
        scopeUser: ScopeEnum[]
    }

    // api
    interface auth {
        menu: functionItemTable[]
        permission: Record<FuncCodeEnum, Record<OperationCodeEnum, permission>>
    }

    interface func {
        functionItem: functionItemTable[]
        operationItem: operationItemTable[]
        functionRoleBanchRelation: Record<FuncCodeEnum, Record<OperationCodeEnum, scope>>
    }

    interface roleBanchList {
        availableBanch: companyBanchTypes.TABLE[]
        availableRole: roleTypes.TABLE[]
        availableUser: userTypes.TABLE[]
        scopeBanch: Record<FuncCodeEnum, Record<OperationCodeEnum, number[]>>
        scopeRole: Record<FuncCodeEnum, Record<OperationCodeEnum, number[]>>
        scopeUser: Record<FuncCodeEnum, Record<OperationCodeEnum, number[]>>
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
    FuncCodeEnum,
    OperationCodeEnum,
    ScopeEnum
}
