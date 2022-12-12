import { action, statusType } from '../types'

const statusState = {
    onEntry: false,

    onFetchCompany: false,
    onUpdateCompany: false,

    onFetchBanch: false,
    onCreateBanch: false,
    onUpdateBanch: false,
    onDeleteBanch: false,

    onFetchUserAll: false,
    onUpdateUser: false,

    onFetchSelfData: false,
    onUpdateSelfData: false,

    onFetchBanchStyle: false,
    onCreateBanchStyle: false,
    onUpdateBanchStyle: false,
    onDeleteBanchStyle: false,

    onFetchBanchRule: false,
    onCreateBanchRule: false,
    onUpdateBanchRule: false,
    onDeleteBanchRule: false,

    onFetchWeekendSetting: false,
    onCreateWeekendSetting: false,
    onDeleteWeekendSetting: false,

    onFetchWaitReply: false,
    onUpdateWaitReply: false,
    onCreateWaitReply: false,

    onFetchWorkTime: false
}

export interface statusReducerType {
    onEntry: boolean

    onFetchCompany: boolean
    onUpdateCompany: boolean

    onFetchBanch: boolean
    onCreateBanch: boolean
    onUpdateBanch: boolean
    onDeleteBanch: boolean

    onFetchUserAll: boolean
    onUpdateUser: boolean

    onFetchSelfData: boolean
    onUpdateSelfData: boolean

    onFetchBanchStyle: boolean
    onCreateBanchStyle: boolean
    onUpdateBanchStyle: boolean
    onDeleteBanchStyle: boolean

    onFetchBanchRule: boolean
    onCreateBanchRule: boolean
    onUpdateBanchRule: boolean
    onDeleteBanchRule: boolean

    onFetchWeekendSetting: boolean
    onCreateWeekendSetting: boolean
    onDeleteWeekendSetting: boolean

    onFetchWaitReply: boolean
    onUpdateWaitReply: boolean
    onCreateWaitReply: boolean

    onFetchWorkTime: boolean
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
