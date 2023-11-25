declare namespace companyTypes {
    interface TABLE {
        CompanyId: number
        CompanyCode: string
        CompanyName: string
        CompanyLocation: string
        companyPhoneNumber: string
        BossId: number
        CreateTime: string
        LastModify: string
    }

    interface reducerType {
        mine: TABLE
    }
}
export default companyTypes
