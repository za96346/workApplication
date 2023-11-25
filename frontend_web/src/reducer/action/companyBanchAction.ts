import companyBanchTypes from 'types/companyBanch'
import { type action, companyBanchType } from '../types'

class companyBanchAction {
    setAll (v: companyBanchTypes.reducerType['all']): action {
        return {
            type: companyBanchType.SET_BANCH_ALL,
            payload: {
                all: v
            }
        }
    }

    setSelector (v: companyBanchTypes.reducerType['selector']): action {
        return {
            type: companyBanchType.SET_BANCH_SELECTOR,
            payload: {
                selector: v
            }
        }
    }
}
export default new companyBanchAction()
