import * as React from 'react'
import * as ReactDOM from 'react-dom/client'
import { Provider } from 'react-redux'

import { store } from './reducer/store'
import { ConfigProvider } from 'antd'
import Locale from 'antd/es/locale/zh_TW'
import RouteIndex from 'Route'

const root = ReactDOM.createRoot(document.getElementById('root') as NonNullable<Element>)
root.render(
    <ConfigProvider locale={Locale}>
        <Provider store={store} children={''}>
            <RouteIndex />
        </Provider>
    </ConfigProvider>
)
