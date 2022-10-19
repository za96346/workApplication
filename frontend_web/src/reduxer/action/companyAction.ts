import { BanchRuleType, BanchStyleType, BanchType, UserType } from '../../type'
import { store } from '../store'
import { action, companyType } from '../types'

class companyAction {
    constructor () {
        this.setBanch = this.setBanch.bind(this)
        this.setEmployee = this.setEmployee.bind(this)
    }

    setBanch (banch: BanchType[]): action {
        return {
            type: companyType.SET_BANCH,
            payload: {
                ...store.getState().company,
                banch: banch || []
            }
        }
    }

    setEmployee (employee: UserType[]): action {
        return {
            type: companyType.SET_EMPLOYEE,
            payload: {
                ...store.getState().company,
                employee: employee || []
            }
        }
    }

    setBanchStyle (banchStyle: BanchStyleType[]): action {
        return {
            type: companyType.SET_BANCH_STYLE,
            payload: {
                ...store.getState().company,
                banchStyle: banchStyle || []
            }
        }
    }

    setBanchRule (banchRule: BanchRuleType[]): action {
        return {
            type: companyType.SET_BANCH_RULE,
            payload: {
                ...store.getState().company,
                banchRule: banchRule || []
            }
        }
    }

    clearCompanyAll (): action {
        return {
            type: companyType.CLEAR_COMPANY_ALL,
            payload: {}
        }
    }
}
export default new companyAction()
