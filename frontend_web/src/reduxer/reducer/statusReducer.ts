import { action, statusType } from '../types'

const statusState = {
    onFetchCompany: false,
    onUpdateCompany: false,

    onFetchBanch: false,
    onCreateBanch: false,
    onUpdateBanch: false,

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
    onFetchCompany: boolean
    onUpdateCompany: boolean

    onFetchBanch: boolean
    onCreateBanch: boolean
    onUpdateBanch: boolean

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
