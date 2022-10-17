import { action, statusType } from '../types'

const statusState = {
    onFetchBanch: false,
    onFetchUserAll: false,
    onFetchSelfData: false,

    onFetchBanchStyle: false,
    onCreateBanchStyle: false,
    onUpdateBanchStyle: false,

    onFetchBanchRule: false,
    onCreateBanchRule: false,
    onUpdateBanchRule: false
}

export interface statusReducerType {
    onFetchBanch: boolean
    onFetchUserAll: boolean
    onFetchSelfData: boolean

    onFetchBanchStyle: boolean
    onCreateBanchStyle: boolean
    onUpdateBanchStyle: boolean

    onFetchBanchRule: boolean
    onCreateBanchRule: boolean
    onUpdateBanchRule: boolean
}

export const statusReducer = (state = statusState, action: action): any => {
    const a = Object.values(statusType).filter((item) => item === action.type)
    if (action.type === statusType.CLEAR_STATUS_ALL) {
        return statusState
    } else if (a.length > 0) {
        return {
            ...action.payload
        }
    } else {
        return state
    }
}
