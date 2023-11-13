import { companyBanchReducerType } from 'reducer/reducer/companyBanchReducer'
import { type action, companyBanchType } from '../types'

class companyBanchAction {
    setAll (v: companyBanchReducerType['all']): action {
        return {
            type: companyBanchType.SET_ALL,
            payload: {
                all: v
            }
        }
    }
}
export default new companyBanchAction()
