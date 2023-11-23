import * as React from 'react'
import * as ReactDOM from 'react-dom/client'
import { Provider } from 'react-redux'

import { store } from './reducer/store'
import { ConfigProvider } from 'antd'
import Locale from 'antd/es/locale/zh_TW'
import RouteIndex from 'Route'

import 'bootstrap/dist/css/bootstrap.min.css'
import stylees from './index.scss'

window.styles = stylees

const root = ReactDOM.createRoot(document.getElementById('root') as NonNullable<Element>)
root.render(
    <ConfigProvider locale={Locale}>
        <Provider store={store}>
            <RouteIndex />
        </Provider>
    </ConfigProvider>
)
