import { BanchType, UserType } from '../../type'
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
                banch
            }
        }
    }

    setEmployee (employee: UserType[]): action {
        return {
            type: companyType.SET_EMPLOYEE,
            payload: {
                ...store.getState().company,
                employee
            }
        }
    }
}
export default new companyAction()
