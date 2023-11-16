declare namespace companyBanchTypes {
    interface TABLE {
        CompanyId: number
        BanchId: number
        BanchName: string
        DeleteFlag: Flag
        DeleteTime: string
        CreateTime: string
        LastModify: string
    }
}
export default companyBanchTypes
