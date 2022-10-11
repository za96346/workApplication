import { userType, action } from '../types'

class userAction {
    constructor () {
        this.setToken = this.setToken.bind(this)
        this.clearToken = this.clearToken.bind(this)
    }

    setToken (token: string): action {
        return {
            type: userType.SET_TOKEN,
            payload: token
        }
    }

    clearToken (): action {
        return {
            type: userType.CLEAR_TOKEN,
            payload: ''
        }
    }
}

export default new userAction()