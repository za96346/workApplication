import storage from 'redux-persist/lib/storage'
import { persistStore, persistReducer } from 'redux-persist'
import thunk from 'redux-thunk'
import { composeWithDevTools } from 'redux-devtools-extension'
import {
    legacy_createStore as createStore,
    combineReducers,
    applyMiddleware
} from 'redux'
import { loadingReducer } from './reducer/loadingReducer'

export const allAction = {
    loading: loadingReducer
}
export const reducer = combineReducers(allAction)

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
