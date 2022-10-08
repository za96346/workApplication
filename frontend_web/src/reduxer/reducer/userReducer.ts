import { action, userType } from '../types'

const userState = {
    token: ''
}

export const userReducer = (state = userState, action: action): any => {
    switch (action.type) {
        case userType.SET_TOKEN:
            return {
                token: action.payload
            }
        case userType.CLEAR_TOKEN:
            return {
                token: action.payload
            }
        default:
            return state
    }
}
