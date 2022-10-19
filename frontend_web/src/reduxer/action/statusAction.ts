import { store } from "../store"
import { action, statusType } from "../types"

class statusAcion {
    constructor () {
        this.onFetchBanch = this.onFetchBanch.bind(this)
        this.onFetchUserAll = this.onFetchUserAll.bind(this)
    }

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

    onFetchUserAll (s: boolean): action {
        return {
            type: s ? statusType.FETCH_USER_ALL_ON : statusType.FETCH_USER_ALL_OFF,
            payload: {
                ...store.getState().status,
                onFetchUserAll: s
            }
        }
    }

    onFetchSelfData (s: boolean): action {
        return {
            type: s ? statusType.FETCH_SELF_DATA_ON : statusType.FETCH_SELF_DATA_OFF,
            payload: {
                ...store.getState().status,
                onFetchSelfData: s
            }
        }
    }

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

    clearStatusAll (): action {
        return {
            type: statusType.CLEAR_STATUS_ALL,
            payload: {}
        }
    }
}
export default new statusAcion()
