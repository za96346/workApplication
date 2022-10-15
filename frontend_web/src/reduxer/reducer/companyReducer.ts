import { action, companyType } from '../types'

const companyState = {
    banch: [],
    employee: []
}

export const companyReducer = (state = companyState, action: action): any => {
    switch (action.type) {
        case companyType.SET_BANCH:
            return {
                ...action.payload
            }
        case companyType.SET_EMPLOYEE:
            return {
                ...action.payload
            }
        case companyType.CLEAR_COMPANY_ALL:
            return companyState
        default:
            return state
    }
}
