import React, { useEffect, Suspense, lazy } from 'react'

import {
    BrowserRouter as Router,
    Routes,
    Route,
    Navigate
} from 'react-router-dom'
import stylees from './index.module.scss'
import 'antd/dist/reset.css'
import 'bootstrap/dist/css/bootstrap.min.css'

// import WeekendSettingPage from './page/WeekendSetting/WeekendSettingPage'
import { store } from './reduxer/store'
import statusAction from './reduxer/action/statusAction'
import { Spin } from 'antd'
import useReduceing from 'Hook/useReducing'

const Entry = lazy((): any => import('./page/Entry/EntryPage'))
const Layout = lazy((): any => import('./page/Layout'))
const HomePage = lazy((): any => import('./page/Home/HomePage'))
const EmployeeManager = lazy((): any => import('./page/EmployeeManager/EmployeeManager'))
const ShiftPage = lazy((): any => import('./page/Shift/ShiftPage'))
const BanchManager = lazy((): any => import('./page/BanchManager/BanchManager'))
const SignPage = lazy((): any => import('./page/Sign/Sign'))
const ShiftSettingPage = lazy((): any => import('./page/shiftSetting/ShiftSettingPage'))
const ShiftSearchPage = lazy((): any => import('./page/ShiftSearch/ShiftSearchPage'))
const SettingPage = lazy((): any => import('./page/Setting/Index'))
const ErrorPage = lazy((): any => import('./page/ErrorPage'))
const WorkTimeManagerPage = lazy((): any => import('./page/WorkTImeManager/WorkTimeManagerPage'))
const PerformancePage = lazy((): any => import('./page/Performance/Index'))

// global init f
window.styles = stylees
// console.log = () => {}
const App = (): JSX.Element => {
    const { loading } = useReduceing()
    console.log('process env =>', process.env)
    const hasLoading = Object.values(loading).filter((i) => i === true).length > 0
    useEffect(() => {
        // 每次 重整 reset 狀態
        store.dispatch(statusAction.clearStatusAll())
    }, [])
    return (
        <>
            {
                hasLoading && (
                    <div className={window.styles.spinner}>
                        <Spin size='large' />
                    </div>
                )
            }
            <Suspense fallback={<Spin/>}>
                <Router>
                    <Routes>
                        <Route path='/' element={<Layout />}>
                            <Route path='/' element={<Navigate to={'entry/login'} />}/>
                            <Route path='entry/:path' element={<Entry />} />

                            <Route path='home' element={<HomePage />} />

                            <Route path='employeeManager' element={<EmployeeManager />} />

                            <Route path='shift/:banchId' element={<ShiftPage />} />

                            <Route path='banchManager' element={<BanchManager />} />
                            <Route path='sign' element={<SignPage/>}/>
                            {/* <Route path='weekendSetting' element={<WeekendSettingPage />}/> */}
                            <Route path='workTimeManager' element={<WorkTimeManagerPage />} />

                            <Route path='shiftSetting/:banchId' element={<ShiftSettingPage />} />
                            <Route path='shiftSearch' element={<ShiftSearchPage />} />

                            <Route path='setting/:types' element={<SettingPage />} />
                            <Route path='performance/:banchId' element={<PerformancePage />}/>
                        </Route>
                        <Route path='*' element={<ErrorPage/>} />
                    </Routes>
                </Router>
            </Suspense>
        </>

    )
}

export default App
