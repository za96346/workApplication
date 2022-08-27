import * as React from 'react';
import * as ReactDOM from 'react-dom/client';
import App from './App.tsx';
import { Provider } from 'react-redux';
import { composeWithDevTools } from 'redux-devtools-extension';
import { PersistGate } from "redux-persist/integration/react";
import storage from 'redux-persist/lib/storage';
import { persistStore, persistReducer } from 'redux-persist';
import thunk from "redux-thunk";
import {
    legacy_createStore as createStore,
    combineReducers,
	applyMiddleware
} from 'redux';
import { userReducer } from './reduxer/reducer/userReducer';

// npm i -D @types/node-sass
// npm i -D node-sass

// npm install --save-dev webpack webpack-cli
// npm install --save react react-dom
// npm install --save-dev @types/react @types/react-dom
// npm install --save-dev typescript ts-loader source-map-loader
// npm install --save typescript @types/node @types/react @types/react-dom @types/jest
// npm install typed-scss-modules --save-dev

// npm install style-loader css-loader sass-loader --save-dev

// npm install mini-css-extract-plugin --save-dev

const reducer = combineReducers({
	user: userReducer
})

// 持久化根reducers
const persistedReducer = persistReducer({
		key: 'root',
		storage,
		whitelist: ['loginData']
	},
	reducer
);
  
export const store = createStore(
  persistedReducer,
  composeWithDevTools(applyMiddleware(...[thunk])),
);
const persisStore = persistStore(store);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Provider store={store}>
		<PersistGate loading={null} persistor={persisStore}>
			<App />
		</PersistGate>
    </Provider>

  </React.StrictMode>
);
