import userTypes from 'types/user'
import { type action, userType } from '../types'

const userState = {
    mine: {},
    employee: [],
    selector: []
}

export interface userReducerType {
    mine: userTypes.TABLE
    employee: userTypes.TABLE[]
    selector: userTypes.TABLE[]
}

export const userReducer = (state = userState, action: action): any => {
    const a = Object.values(userType).filter((item) => item === action.type)
    if (a.length > 0) {
        return {
            ...state,
            ...action.payload
        }
    }
    return state
}
