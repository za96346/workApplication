import { type action, userType } from '../types'
import { userReducerType } from 'reducer/reducer/userReducer'

class userAction {
    setMine (v: userReducerType['mine']): action {
        return {
            type: userType.SET_USER_MINE,
            payload: {
                mine: v
            }
        }
    }

    setEmployee (v: userReducerType['employee']): action {
        return {
            type: userType.SET_EMPLOYEE,
            payload: {
                employee: v
            }
        }
    }
}
export default new userAction()
