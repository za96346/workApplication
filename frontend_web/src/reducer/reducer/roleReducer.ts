import { type action, roleType } from '../types'

const roleState = {
    all: [],
    single: {
        Permission: {},
        Role: {}
    },
    selector: []
}

export const roleReducer = (state = roleState, action: action): any => {
    const a = Object.values(roleType).filter((item) => item === action.type)
    if (a.length > 0) {
        return {
            ...state,
            ...action.payload
        }
    }
    return state
}
