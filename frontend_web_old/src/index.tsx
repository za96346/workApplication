import * as React from 'react'
import * as ReactDOM from 'react-dom/client'
import App from './App'
import { Provider } from 'react-redux'

import { PersistGate } from 'redux-persist/integration/react'
import { persisStore, store } from './reduxer/store'
import { ConfigProvider } from 'antd'
import Locale from 'antd/es/locale/zh_TW'
import '@vteam_components/cloud/dist/style.js'
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

const root = ReactDOM.createRoot(document.getElementById('root'))
root.render(
    <ConfigProvider locale={Locale}>
        <Provider store={store}>
            <PersistGate loading={null} persistor={persisStore}>
                <App />
            </PersistGate>
        </Provider>
    </ConfigProvider>
)
