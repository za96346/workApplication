import { ShiftSocketType } from '../../type'
import { action, shiftEditType } from '../types'

class shiftEditAction {
    constructor () {
        this.setShiftEdit = this.setShiftEdit.bind(this)
    }

    setShiftEdit (shift: ShiftSocketType): action {
        // console.log(shift)
        return {
            type: shiftEditType.SET_SHIFT,
            payload: {
                ...shift
            }
        }
    }

    clearShiftEdit (): action {
        // console.log(shift)
        return {
            type: shiftEditType.CLEAR_SHIFT_ALL,
            payload: {}
        }
    }
}
export default new shiftEditAction()
