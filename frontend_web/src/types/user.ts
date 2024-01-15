declare namespace userTypes {
    interface TABLE {
        CompanyId: number
        UserId: number
        RoleId: number
        BanchId: number
        UserName: string
        EmployeeNumber: string
        Account: string
        Password: string
        OnWorkDay: string
        Sort: number
        QuitFlag: Flag
        DeleteFlag: Flag
        DeleteTime: string
        CreateTime: string
        LastModify: string
    }

    interface reducerType {
        mine: TABLE
        employee: TABLE[]
        selector: TABLE[]
    }
}
export default userTypes
