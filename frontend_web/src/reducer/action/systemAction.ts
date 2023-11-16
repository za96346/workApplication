import { type action, systemType } from '../types'
import { systemReducerType } from 'reducer/reducer/systemReducer'

class systemAction {
    setAuth (v: systemReducerType['auth']): action {
        return {
            type: systemType.SET_AUTH,
            payload: {
                auth: v
            }
        }
    }

    setSidebar (v: systemReducerType['sidebar']): action {
        return {
            type: systemType.SET_AUTH,
            payload: {
                sidebar: v
            }
        }
    }

    setFunc (v: systemReducerType['func']): action {
        return {
            type: systemType.SET_FUNC,
            payload: {
                func: v
            }
        }
    }

    setRoleBanchList (v: systemReducerType['roleBanchList']): action {
        return {
            type: systemType.SET_ROLE_BANCH_LIST,
            payload: {
                roleBanchList: v
            }
        }
    }
}
export default new systemAction()
