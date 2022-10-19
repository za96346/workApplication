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
import ErrorPage from './page/ErrorPage'
import Entry from './page/Entry/EntryPage'
import EmployeeManager from './page/EmployeeManager/EmployeeManager'
import BanchManager from './page/BanchManager/BanchManager'

// global init
window.styles = styles

const App = (): JSX.Element => {
    console.log('process env =>', process.env)
    return (
        <Router>
            <Routes>
                <Route path='/' element={<Layout />}>
                    <Route path='/' element={<Navigate to={'entry/login'} />}/>
                    <Route path='entry/:path' element={<Entry />} />

                    <Route path='home' element={<></>}/>

                    <Route path='employeeManager' element={<EmployeeManager />} />

                    <Route path='shift/:banchId' element={<ShiftPage />} />

                    <Route path='banchManager' element={<BanchManager />} />

                    <Route path='shiftSetting/:banchId' element={<ShiftSettingPage />} />
                    <Route path='shiftSearch' element={<ShiftSearchPage />} />

                    <Route path='setting/:types' element={<SettingPage />} />
                </Route>
                <Route path='*' element={<ErrorPage/>} />
            </Routes>
        </Router>
    )
}

export default App
