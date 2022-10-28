import storage from 'redux-persist/lib/storage'
import { persistStore, persistReducer } from 'redux-persist'
import thunk from 'redux-thunk'
import { composeWithDevTools } from 'redux-devtools-extension'
import {
    legacy_createStore as createStore,
    combineReducers,
    applyMiddleware
} from 'redux'
import { userReducer } from './reducer/userReducer'
import { companyReducer } from './reducer/companyReducer'
import { statusReducer } from './reducer/statusReducer'
import { shiftEditReducer } from './reducer/shiftEditReducer'

export const reducer = combineReducers({
    user: userReducer,
    company: companyReducer,
    status: statusReducer,
    shiftEdit: shiftEditReducer
})

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
export type RootState = ReturnType<typeof store.getState>
