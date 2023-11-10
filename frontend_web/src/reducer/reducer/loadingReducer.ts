import { type action, loadingType } from '../types'

const loadingState = {
    loading: false
}

export type loadingReducerType = typeof loadingState

export const loadingReducer = (state = loadingState, action: action): any => {
    const a = Object.values(loadingType).filter((item) => item === action.type)
    if (a.length > 0) {
        return {
            ...state,
            ...action.payload
        }
    }
    return state
}
