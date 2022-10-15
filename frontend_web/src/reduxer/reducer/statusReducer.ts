import { action, statusType } from '../types'

const statusState = {
    onFetchBanch: false,
    onFetchUserAll: false
}

export const statusReducer = (state = statusState, action: action): any => {
    switch (action.type) {
        case statusType.FETCH_BANCH_ON:
            return {
                ...action.payload
            }
        case statusType.FETCH_BANCH_OFF:
            return {
                ...action.payload
            }
        case statusType.FETCH_USER_ALL_ON:
            return {
                ...action.payload
            }
        case statusType.FETCH_USER_ALL_OFF:
            return {
                ...action.payload
            }
        case statusType.CLEAR_STATUS_ALL:
            return statusState
        default:
            return state
    }
}
