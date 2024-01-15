declare namespace companyBanchTypes {
    interface TABLE {
        CompanyId: number
        BanchId: number
        BanchName: string
        Sort: number
        DeleteFlag: Flag
        DeleteTime: string
        CreateTime: string
        LastModify: string
    }

    interface reducerType {
        all: TABLE[]
        selector: TABLE[]
    }
}
export default companyBanchTypes
