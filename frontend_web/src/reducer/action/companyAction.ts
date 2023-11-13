import { type action, companyType } from '../types'
import { companyReducerType } from 'reducer/reducer/companyReducer'

class companyAction {
    setMine (v: companyReducerType['mine']): action {
        return {
            type: companyType.SET_MINE,
            payload: {
                mine: v
            }
        }
    }
}
export default new companyAction()
