import userTypes from 'types/user'
import { type action, userType } from '../types'

class userAction {
    setMine (v: userTypes.TABLE): action {
        return {
            type: userType.SET_MINE,
            payload: {
                mine: v
            }
        }
    }

    setEmployee (v: userTypes.TABLE[]): action {
        return {
            type: userType.SET_EMPLOYEE,
            payload: {
                employee: v
            }
        }
    }
}
export default new userAction()
