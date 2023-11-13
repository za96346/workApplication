import systemTypes from 'types/system'
import { type action, systemType } from '../types'

class systemAction {
    setAuth (v: systemTypes.auth): action {
        return {
            type: systemType.SET_AUTH,
            payload: {
                auth: v
            }
        }
    }

    setSidebar (v: boolean): action {
        return {
            type: systemType.SET_AUTH,
            payload: {
                sidebar: v
            }
        }
    }
}
export default new systemAction()
