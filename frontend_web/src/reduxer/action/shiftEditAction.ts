import { OnlineUserType } from '../../type'
import { store } from '../store'
import { action, shiftEditType } from '../types'

class shiftEditAction {
    constructor () {
        this.setShiftEditOnlineUser = this.setShiftEditOnlineUser.bind(this)
    }

    setShiftEditOnlineUser (onlineUser: OnlineUserType[]): action {
        return {
            type: shiftEditType.SET_SHIFT_ONLINE_USER,
            payload: {
                ...store.getState().shiftEdit,
                onlineUser
            }
        }
    }

    setShiftStatus (status: number): action {
        return {
            type: shiftEditType.SET_SHIFT_STATUS,
            payload: {
                ...store.getState().shiftEdit,
                status
            }
        }
    }
}
export default new shiftEditAction()
