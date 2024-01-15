import systemTypes from "./system"

declare namespace roleTypes {
    interface TABLE {
        CompanyId: number
        RoleId: number
        RoleName: string
        Sort: number
        StopFlag: string
        DeleteFlag: Flag
        DeleteTime: string
        CreateTime: string
        LastModify: string
    }

    interface single {
        Permission: systemTypes.auth['permission']
        Role: TABLE
    }

    interface reducerType {
        all: TABLE[]
        single: single
        selector: TABLE[]
    }
}
export default roleTypes
