import { action, userType } from '../types'

const userState = {
    token: '',
    selfData: {

    }
}

export const userReducer = (state = userState, action: action): any => {
    switch (action.type) {
        case userType.SET_TOKEN:
            return {
                ...action.payload
            }
        case userType.SET_SELF_DATA:
            return {
                ...action.payload
            }
        case userType.CLEAR_TOKEN:
            return {
                ...action.payload
            }
        case userType.CLEAR_USER_ALL:
            return userState
        default:
            return state
    }
}
