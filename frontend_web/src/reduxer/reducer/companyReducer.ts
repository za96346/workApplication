import { action, companyType } from '../types'

const companyState = {
    banch: [],
    employee: [],
    banchStyle: [],
    banchRule: []
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
