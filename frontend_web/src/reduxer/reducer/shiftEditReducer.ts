
import { OnlineUserType } from '../../type'
import { action, shiftEditType } from '../types'

const shiftState = {
    onlineUser: [],
    status: 1
}
export interface shiftReducerType {
    onlineUser: OnlineUserType[]
    status: number
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
