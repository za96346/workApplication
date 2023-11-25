import systemTypes from 'types/system'
import { type action, systemType } from '../types'

class systemAction {
    setAuth (v: systemTypes.reducerType['auth']): action {
        return {
            type: systemType.SET_AUTH,
            payload: {
                auth: v
            }
        }
    }

    setSidebar (v: systemTypes.reducerType['sidebar']): action {
        return {
            type: systemType.SET_AUTH,
            payload: {
                sidebar: v
            }
        }
    }

    setFunc (v: systemTypes.reducerType['func']): action {
        return {
            type: systemType.SET_FUNC,
            payload: {
                func: v
            }
        }
    }

    setRoleBanchList (v: systemTypes.reducerType['roleBanchList']): action {
        return {
            type: systemType.SET_ROLE_BANCH_LIST,
            payload: {
                roleBanchList: v
            }
        }
    }
}
export default new systemAction()
