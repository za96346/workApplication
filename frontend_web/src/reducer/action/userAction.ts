import userTypes from 'types/user'
import { type action, userType } from '../types'

class userAction {
    setMine (v: userTypes.reducerType['mine']): action {
        return {
            type: userType.SET_USER_MINE,
            payload: {
                mine: v
            }
        }
    }

    setEmployee (v: userTypes.reducerType['employee']): action {
        return {
            type: userType.SET_EMPLOYEE,
            payload: {
                employee: v
            }
        }
    }

    setSelector (v: userTypes.reducerType['selector']): action {
        return {
            type: userType.SET_USER_SELECTOR,
            payload: {
                selector: v
            }
        }
    }
}
export default new userAction()
