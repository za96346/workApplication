
import { ShiftSocketType, shiftTotalType } from '../../type'
import { action, shiftEditType } from '../types'

const shiftState = {
    BanchId: -1,
    EditUser: [],
    OnlineUser: [],
    ShiftData: [],
    BanchStyle: [],
    WeekendSetting: [],
    Status: 1, // 1 開放編輯、 2 主管審核 3 確認發布 ,
    StartDay: '',
    EndDay: '',
    total: []
}

export interface shiftReducerType extends ShiftSocketType {
    total: shiftTotalType[]
}

export const shiftEditReducer = (state = shiftState, action: action): any => {
    const a = Object.values(shiftEditType).filter((item) => item === action.type)
    if (action.type === shiftEditType.CLEAR_SHIFT_ALL) {
        return shiftState
    } else if (a.length > 0) {
        return {
            ...action.payload
        }
    } else {
        return state
    }
}
