import { OnlineUserType } from '../../type'
import { action, shiftEditType } from '../types'

class shiftEditAction {
    constructor () {
        this.setShiftEditOnlineUser = this.setShiftEditOnlineUser.bind(this)
    }

    setShiftEditOnlineUser (onlineUser: OnlineUserType[]): action {
        return {
            type: shiftEditType.SET_SHIFT_ONLINE_USER,
            payload: {
                onlineUser
            }
        }
    }
}
export default new shiftEditAction()
