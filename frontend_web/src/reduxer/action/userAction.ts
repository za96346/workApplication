import { SelfDataType } from '../../type'
import { store } from '../store'
import { userType, action } from '../types'

class userAction {
    constructor () {
        this.setToken = this.setToken.bind(this)
        this.clearToken = this.clearToken.bind(this)
    }

    setToken (token: string): action {
        return {
            type: userType.SET_TOKEN,
            payload: {
                ...store.getState().user,
                token
            }
        }
    }

    setSelfData (selfData: SelfDataType): action {
        return {
            type: userType.SET_SELF_DATA,
            payload: {
                ...store.getState().user,
                selfData
            }
        }
    }

    clearToken (): action {
        return {
            type: userType.CLEAR_TOKEN,
            payload: {
                ...store.getState().user,
                token: ''
            }
        }
    }

    clearUserAll (): action {
        return {
            type: userType.CLEAR_USER_ALL,
            payload: {}
        }
    }
}

export default new userAction()
