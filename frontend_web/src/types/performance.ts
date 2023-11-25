import companyBanchTypes from './companyBanch'
import userTypes from './user'

declare namespace performanceTypes {
    interface TABLE {
        CompanyId: number
        UserId: number
        PerformanceId: number
        Year: number
        Month: number
        BanchId: number
        Goal: string
        Attitude: number
        Efficiency: number
        Professional: number
        Directions: string
        BeLate: number
        DayOffNotOnRule: number
        DeleteFlag: Flag
        DeleteTime: string
        CreateTime: string
        LastModify: string
    }

    interface year {
        UserName: userTypes.TABLE['UserName']
        Year: TABLE['Year']
        Score: number
    }

    interface reducerType {
        all: Array<
        TABLE & {
            UserName: userTypes.TABLE['UserName']
            RoleId: userTypes.TABLE['RoleId']
            BanchName: companyBanchTypes.TABLE['BanchName']
        }>
        year: year[]
    }
}
export default performanceTypes
