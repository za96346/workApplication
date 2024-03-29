import { type action, companyBanchType } from '../types'

const companyBanchState = {
    all: [],
    selector: []
}

export const companyBanchReducer = (state = companyBanchState, action: action): any => {
    const a = Object.values(companyBanchType).filter((item) => item === action.type)
    if (a.length > 0) {
        return {
            ...state,
            ...action.payload
        }
    }
    return state
}
