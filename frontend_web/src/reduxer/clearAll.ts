import companyAction from './action/companyAction'
import statusAction from './action/statusAction'
import userAction from './action/userAction'
import { store } from './store'

export const clearAll = (): void => {
    store.dispatch(companyAction.clearCompanyAll())
    store.dispatch(statusAction.clearStatusAll())
    store.dispatch(userAction.clearUserAll())
}
