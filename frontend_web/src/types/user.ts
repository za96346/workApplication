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
        DeleteFlag: Flag
        DeleteTime: string
        CreateTime: string
        LastModify: string
    }
}
export default userTypes
