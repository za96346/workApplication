import React from 'react'

import {
    BrowserRouter as Router,
    Routes,
    Route,
    Navigate
} from 'react-router-dom'
import styles from './index.module.scss'
import 'antd/dist/antd.min.css'

import Layout from './page/Layout'
import ShiftSettingPage from './page/shiftSetting/ShiftSettingPage'
import ShiftPage from './page/Shift/ShiftPage'
import SettingPage from './page/Setting/SettingPage'
import ShiftSearchPage from './page/ShiftSearch/ShiftSearchPage'
import language from './language'
import ErrorPage from './page/ErrorPage'
import Entry from './page/Entry/EntryPage'
import statics from './statics'

// global init
window.styles = styles
window.language = language
window.statics = statics

const App = (): JSX.Element => {
    console.log(window.language)
    return (
        <Router>
            <Routes>
                <Route path='/' element={<Layout />}>
                    <Route path='/' element={<Navigate to={'entry/login'} />}/>
                    <Route path='entry/:path' element={<Entry />} />

                    <Route path='shift' element={<ShiftPage />} />
                    <Route path='shift/:banch' element={<ShiftPage />} />

                    <Route path='shiftSetting/:banch' element={<ShiftSettingPage />} />
                    <Route path='shiftSearch' element={<ShiftSearchPage />} />

                    <Route path='setting/:types' element={<SettingPage />} />
                </Route>
                <Route path='*' element={<ErrorPage/>} />
            </Routes>
        </Router>
    )
}

export default App
