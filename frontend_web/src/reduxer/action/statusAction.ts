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

    onFetchUserAll (s: boolean): action {
        return {
            type: s ? statusType.FETCH_USER_ALL_ON : statusType.FETCH_USER_ALL_OFF,
            payload: {
                ...store.getState().status,
                onFetchUserAll: s
            }
        }
    }
}
export default new statusAcion()
