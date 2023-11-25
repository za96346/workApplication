import { type action, performanceType } from '../types'

const performanceState = {
    all: [],
    year: []
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
