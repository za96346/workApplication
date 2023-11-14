import { companyBanchReducerType } from 'reducer/reducer/companyBanchReducer'
import { type action, companyBanchType } from '../types'

class companyBanchAction {
    setAll (v: companyBanchReducerType['all']): action {
        return {
            type: companyBanchType.SET_BANCH_ALL,
            payload: {
                all: v
            }
        }
    }

    setSelector (v: companyBanchReducerType['selector']): action {
        return {
            type: companyBanchType.SET_BANCH_SELECTOR,
            payload: {
                selector: v
            }
        }
    }
}
export default new companyBanchAction()
