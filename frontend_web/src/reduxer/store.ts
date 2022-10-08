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

export const reducer = combineReducers({
    user: userReducer
})

// 持久化根reducers
const persistedReducer = persistReducer({
    key: 'root',
    storage,
    whitelist: ['user']
},
reducer
)

export const store = createStore(
    persistedReducer,
    composeWithDevTools(applyMiddleware(...[thunk]))
)
export const persisStore = persistStore(store)
