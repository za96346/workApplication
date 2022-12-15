import { ShiftSocketType } from '../../type'
import { store } from '../store'
import { action, shiftEditType } from '../types'

class shiftEditAction {
    constructor () {
        this.setShiftEdit = this.setShiftEdit.bind(this)
    }

    setShiftEdit (shift: ShiftSocketType): action {
        return {
            type: shiftEditType.SET_SHIFT,
            payload: {
                ...store.getState().shiftEdit,
                ...shift
            }
        }
    }
}
export default new shiftEditAction()
