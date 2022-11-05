import { BanchRuleType, BanchStyleType, BanchType, CompanyType, UserType, WeekendSettingType } from '../../type'
import { action, companyType } from '../types'

export const companyState = {
    banch: [],
    employee: [],
    banchStyle: [],
    banchRule: [],
    weekendSetting: [],
    info: {}
}

export interface companyReducerType {
    banch: BanchType[]
    employee: UserType[]
    banchStyle: BanchStyleType[]
    banchRule: BanchRuleType[]
    weekendSetting: WeekendSettingType[]
    info: CompanyType
}

export const companyReducer = (state = companyState, action: action): any => {
    const a = Object.values(companyType).filter((item) => item === action.type)
    if (action.type === companyType.CLEAR_COMPANY_ALL) {
        return companyState
    } else if (a.length > 0) {
        return {
            ...action.payload
        }
    } else {
        return state
    }
}
