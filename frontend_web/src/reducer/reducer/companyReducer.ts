import companyTypes from 'types/company'
import { type action, companyType } from '../types'

const companyState = {
    mine: {}
}

export interface companyReducerType {
    mine: companyTypes.TABLE
}

export const companyReducer = (state = companyState, action: action): any => {
    const a = Object.values(companyType).filter((item) => item === action.type)
    if (a.length > 0) {
        return {
            ...state,
            ...action.payload
        }
    }
    return state
}
