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
    onCreateUser: false,

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

    onFetchPerformance: false,
    onUpdatePerformance: false,
    onDeletePerformance: false,
    onCreatePerformance: false,
    onCopyPerformance: false,
    onFetchYearPerformance: false,

    onFetchWorkTime: false,
    onFetchShiftMonth: false,

    onFetchShiftTotal: false,
    onFetchShiftHistory: false
}

export type statusReducerType = typeof statusState

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
