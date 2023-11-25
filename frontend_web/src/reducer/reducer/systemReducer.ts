import { type action, systemType } from '../types'

const systemState = {
    auth: {},
    sidebar: true,
    func: {},
    roleBanchList: {}
}

export const systemReducer = (state = systemState, action: action): any => {
    const a = Object.values(systemType).filter((item) => item === action.type)
    if (a.length > 0) {
        return {
            ...state,
            ...action.payload
        }
    }
    return state
}
