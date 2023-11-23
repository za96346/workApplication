import performanceTypes from 'types/performance'
import { type action, performanceType } from '../types'
import userTypes from 'types/user'
import companyBanchTypes from 'types/companyBanch'

const performanceState = {
    all: []
}

export interface performanceReducerType {
    all: Array<
    performanceTypes.TABLE & {
        UserName: userTypes.TABLE['UserName']
        BanchName: companyBanchTypes.TABLE['BanchName']
    }>
}

export const performanceReducer = (state = performanceState, action: action): any => {
    const a = Object.values(performanceType).filter((item) => item === action.type)
    if (a.length > 0) {
        return {
            ...state,
            ...action.payload
        }
    }
    return state
}
