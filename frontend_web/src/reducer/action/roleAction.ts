import { type action, roleType } from '../types'
import roleTypes from 'types/role'

class roleAction {
    setAll (v: roleTypes.reducerType['all']): action {
        return {
            type: roleType.SET_ROLE_ALL,
            payload: {
                all: v
            }
        }
    }

    setSingle (v: roleTypes.reducerType['single']): action {
        return {
            type: roleType.SET_ROLE_SINGLE,
            payload: {
                single: v
            }
        }
    }

    setSelector (v: roleTypes.reducerType['selector']): action {
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
