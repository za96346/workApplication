
import { OnlineUserType } from '../../type'
import { action, shiftEditType } from '../types'

const shiftState = {
    onlineUser: []
}
export interface shiftReducerType {
    onlineUser: OnlineUserType[]
}

export const shiftEditReducer = (state = shiftState, action: action): any => {
    if (action.type === shiftEditType.SET_SHIFT_ONLINE_USER) {
        return {
            ...action.payload
        }
    } else {
        return state
    }
}
