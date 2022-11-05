import { store } from "../store"
import { action, statusType } from "../types"

class statusAcion {
    constructor () {
        this.onFetchBanch = this.onFetchBanch.bind(this)
        this.onFetchUserAll = this.onFetchUserAll.bind(this)
    }

    // 登入 註冊
    onEntry (s: boolean): action {
        return {
            type: s ? statusType.ENTRY_ON : statusType.ENTRY_OFF,
            payload: {
                ...store.getState().status,
                onEntry: s
            }
        }
    }

    // 部門
    onFetchBanch (s: boolean): action {
        return {
            type: s ? statusType.FETCH_BANCH_ON : statusType.FETCH_BANCH_OFF,
            payload: {
                ...store.getState().status,
                onFetchBanch: s
            }
        }
    }

    onUpdateBanch (s: boolean): action {
        return {
            type: s ? statusType.UPDATE_BANCH_ON : statusType.UPDATE_BANCH_OFF,
            payload: {
                ...store.getState().status,
                onUpdateBanch: s
            }
        }
    }

    onCreateBanch (s: boolean): action {
        return {
            type: s ? statusType.CREATE_BANCH_ON : statusType.CREATE_BANCH_OFF,
            payload: {
                ...store.getState().status,
                onCreateBanch: s
            }
        }
    }

    // 所有員工
    onFetchUserAll (s: boolean): action {
        return {
            type: s ? statusType.FETCH_USER_ALL_ON : statusType.FETCH_USER_ALL_OFF,
            payload: {
                ...store.getState().status,
                onFetchUserAll: s
            }
        }
    }

    onUpdateUser (s: boolean): action {
        return {
            type: s ? statusType.UPDATE_USER_ON : statusType.UPDATE_USER_OFF,
            payload: {
                ...store.getState().status,
                onUpdateUser: s
            }
        }
    }

    // 自己資料
    onFetchSelfData (s: boolean): action {
        return {
            type: s ? statusType.FETCH_SELF_DATA_ON : statusType.FETCH_SELF_DATA_OFF,
            payload: {
                ...store.getState().status,
                onFetchSelfData: s
            }
        }
    }

    onUpdateSelfData (s: boolean): action {
        return {
            type: s ? statusType.UPDATE_SELF_DATA_ON : statusType.UPDATE_SELF_DATA_OFF,
            payload: {
                ...store.getState().status,
                onUpdateSelfData: s
            }
        }
    }

    // 部門樣式
    onFetchBanchStyle (s: boolean): action {
        return {
            type: s ? statusType.FETCH_BANCH_STYLE_ON : statusType.FETCH_BANCH_OFF,
            payload: {
                ...store.getState().status,
                onFetchBanchStyle: s
            }
        }
    }

    onUpdateBanchStyle (s: boolean): action {
        return {
            type: s ? statusType.UPDATE_BANCH_STYLE_ON : statusType.UPDATE_BANCH_STYLE_OFF,
            payload: {
                ...store.getState().status,
                onUpdateBanchStyle: s
            }
        }
    }

    onCreateBanchStyle (s: boolean): action {
        return {
            type: s ? statusType.CREATE_BANCH_STYLE_ON : statusType.CREATE_BANCH_STYLE_OFF,
            payload: {
                ...store.getState().status,
                onCreateBanchStyle: s
            }
        }
    }

    onDeleteBanchStyle (s: boolean): action {
        return {
            type: s ? statusType.DELETE_BANCH_STYLE_ON : statusType.DELETE_BANCH_STYLE_OFF,
            payload: {
                ...store.getState().status,
                onDeleteBanchStyle: s
            }
        }
    }

    // 部門規則
    onFetchBanchRule (s: boolean): action {
        return {
            type: s ? statusType.FETCH_BANCH_RULE_ON : statusType.FETCH_BANCH_RULE_OFF,
            payload: {
                ...store.getState().status,
                onFetchBanchRule: s
            }
        }
    }

    onUpdateBanchRule (s: boolean): action {
        return {
            type: s ? statusType.UPDATE_BANCH_RULE_ON : statusType.UPDATE_BANCH_RULE_OFF,
            payload: {
                ...store.getState().status,
                onUpdateBanchRule: s
            }
        }
    }

    onCreateBanchRule (s: boolean): action {
        return {
            type: s ? statusType.CREATE_BANCH_RULE_ON : statusType.CREATE_BANCH_RULE_OFF,
            payload: {
                ...store.getState().status,
                onCreateBanchRule: s
            }
        }
    }

    onDeleteBanchRule (s: boolean): action {
        return {
            type: s ? statusType.DELETE_BANCH_RULE_ON : statusType.DELETE_BANCH_RULE_OFF,
            payload: {
                ...store.getState().status,
                onDeleteBanchRule: s
            }
        }
    }

    // 公司
    onFetchCompany (s: boolean): action {
        return {
            type: s ? statusType.FETCH_COMPANY_ON : statusType.FETCH_COMPANY_OFF,
            payload: {
                ...store.getState().status,
                onFetchCompany: s
            }
        }
    }

    onUpdateCompany (s: boolean): action {
        return {
            type: s ? statusType.UPDATE_COMPANY_ON : statusType.UPDATE_COMPANY_OFF,
            payload: {
                ...store.getState().status,
                onUpdateCompany: s
            }
        }
    }

    // 假日設定
    onFetchWeekendSetting (s: boolean): action {
        return {
            type: s ? statusType.FETCH_WEEKEND_SETTING_ON : statusType.FETCH_WEEKEND_SETTING_OFF,
            payload: {
                ...store.getState().status,
                onFetchWeekendSetting: s
            }
        }
    }

    onCreateWeekendSetting (s: boolean): action {
        return {
            type: s ? statusType.CREATE_WEEKEND_SETTING_ON : statusType.CREATE_WEEKEND_SETTING_OFF,
            payload: {
                ...store.getState().status,
                onCreateWeekendSetting: s
            }
        }
    }

    onDeleteWeekendSetting (s: boolean): action {
        return {
            type: s ? statusType.DELETE_WEEKEND_SETTING_ON : statusType.DELETE_WEEKEND_SETTING_OFF,
            payload: {
                ...store.getState().status,
                onDeleteWeekendSetting: s
            }
        }
    }

    // 清除
    clearStatusAll (): action {
        return {
            type: statusType.CLEAR_STATUS_ALL,
            payload: {}
        }
    }
}
export default new statusAcion()
