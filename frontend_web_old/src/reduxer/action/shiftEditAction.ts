import { shiftHistoryType, ShiftSocketType, shiftTotalType } from '../../type'
import { action, shiftEditType } from '../types'
import { store } from '../store'

class shiftEditAction {
    constructor () {
        this.setShiftEdit = this.setShiftEdit.bind(this)
    }

    setShiftEdit (shift: ShiftSocketType): action {
        // console.log(shift)
        return {
            type: shiftEditType.SET_SHIFT,
            payload: {
                ...store.getState().shiftEdit,
                ...shift
            }
        }
    }

    setShiftTotal (data: shiftTotalType[]): action {
        return {
            type: shiftEditType.SET_SHIFT,
            payload: {
                ...store.getState().shiftEdit,
                total: data || []
            }
        }
    }

    setShiftHistory (data: shiftHistoryType[]): action {
        return {
            type: shiftEditType.SET_SHIFT_HISTORY,
            payload: {
                ...store.getState().shiftEdit,
                history: data || []
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
