import { roleReducerType } from 'reducer/reducer/roleReducer'
import { type action, roleType } from '../types'

class roleAction {
    setAll (v: roleReducerType['all']): action {
        return {
            type: roleType.SET_ROLE_ALL,
            payload: {
                all: v
            }
        }
    }

    setSingle (v: roleReducerType['single']): action {
        return {
            type: roleType.SET_ROLE_SINGLE,
            payload: {
                single: v
            }
        }
    }

    setSelector (v: roleReducerType['selector']): action {
        return {
            type: roleType.SET_ROLE_SELECTOR,
            payload: {
                selector: v
            }
        }
    }

    clearSingle (): action {
        return {
            type: roleType.CLEAR_ROLE_SINGLE,
            payload: {
                single: {}
            }
        }
    }
}
export default new roleAction()
