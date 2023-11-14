import systemTypes from "./system"

declare namespace roleTypes {
    interface TABLE {
        CompanyId: number
        RoleId: number
        RoleName: string
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
}
export default roleTypes
