import storage from 'redux-persist/lib/storage'
import { persistStore, persistReducer } from 'redux-persist'
import thunk from 'redux-thunk'
import { composeWithDevTools } from 'redux-devtools-extension'
import {
    legacy_createStore as createStore,
    combineReducers,
    applyMiddleware
} from 'redux'
import { loadingReducer, loadingReducerType } from './reducer/loadingReducer'
import { systemReducer, systemReducerType } from './reducer/systemReducer'
import { userReducer, userReducerType } from './reducer/userReducer'
import loadingAction from './action/loadingAction'
import systemAction from './action/systemAction'
import userAction from './action/userAction'
import { companyReducer, companyReducerType } from './reducer/companyReducer'
import companyAction from './action/companyAction'
import { companyBanchReducer, companyBanchReducerType } from './reducer/companyBanchReducer'
import companyBanchAction from './action/companyBanchAction'
import { roleReducer, roleReducerType } from './reducer/roleReducer'
import roleAction from './action/roleAction'
import { performanceReducer, performanceReducerType } from './reducer/performanceReducer'
import performanceAction from './action/performanceAction'

export const allReducer = {
    loading: loadingReducer,
    system: systemReducer,
    user: userReducer,
    company: companyReducer,
    companyBanch: companyBanchReducer,
    role: roleReducer,
    performance: performanceReducer
}
export const allAction = {
    loading: loadingAction,
    system: systemAction,
    user: userAction,
    company: companyAction,
    companyBanch: companyBanchAction,
    role: roleAction,
    performance: performanceAction
}

export interface RootState {
    loading: loadingReducerType
    system: systemReducerType
    user: userReducerType
    company: companyReducerType
    companyBanch: companyBanchReducerType
    role: roleReducerType
    performance: performanceReducerType
}

export const reducer = combineReducers<RootState>(allReducer)

// 持久化根reducers
const persistedReducer = persistReducer({
    key: 'root',
    storage,
    blacklist: []
},
reducer
)

export const store = createStore(
    persistedReducer,
    composeWithDevTools(applyMiddleware(...[thunk]))
)
export const persisStore = persistStore(store)

export type AppDispatch = typeof store.dispatch
