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
}
export default performanceTypes
