import companyTypes from 'types/company'
import { type action, companyType } from '../types'

class companyAction {
    setMine (v: companyTypes.reducerType['mine']): action {
        return {
            type: companyType.SET_COMPANY_MINE,
            payload: {
                mine: v
            }
        }
    }
}
export default new companyAction()
