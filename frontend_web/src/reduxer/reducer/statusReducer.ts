import { action, statusType } from '../types'

const statusState = {
    onFetchBanch: false,
    onFetchUserAll: false,
    onFetchSelfData: false,
    onFetchBanchStyle: false,
    onFetchBanchRule: false
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
