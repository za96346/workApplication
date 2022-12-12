import { BanchRuleType, BanchStyleType, BanchType, CompanyType, UserType, WaitReplyType, WeekendSettingType } from '../../type'
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

    setCompany (info: CompanyType): action {
        return {
            type: companyType.SET_COMPANY,
            payload: {
                ...store.getState().company,
                info: info || {}
            }
        }
    }

    setWeekendSetting (weekendSetting: WeekendSettingType[]): action {
        return {
            type: companyType.SET_WEEKEND_SETTING,
            payload: {
                ...store.getState().company,
                weekendSetting: weekendSetting || []
            }
        }
    }

    setWaitReply (waitReply: WaitReplyType[]): action {
        return {
            type: companyType.SET_WAIT_REPLY,
            payload: {
                ...store.getState().company,
                waitReply: waitReply || []
            }
        }
    }

    setWorkTime (workTime): action {
        return {
            type: companyType.SET_WORK_TIME,
            payload: {
                ...store.getState().company,
                workTime: workTime || []
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
