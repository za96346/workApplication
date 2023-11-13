import companyTypes from 'types/company'
import { type action, companyType } from '../types'

class companyAction {
    setMine (v: companyTypes.TABLE): action {
        return {
            type: companyType.SET_MINE,
            payload: {
                mine: v
            }
        }
    }
}
export default new companyAction()
