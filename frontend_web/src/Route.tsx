import React, { useEffect } from 'react'
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Navigate
} from 'react-router-dom'
import { useAppSelector } from 'hook/redux'
import Layout from 'shared/Layout'

import api from 'api/Index'
import { v4 } from 'uuid'

// page
import BanchManage from 'Page/BanchManage/Index'
import CompanyData from 'Page/CompanyData/Index'
import EmployeeManage from 'Page/EmployeeManage/Index'
import Performance from 'Page/Performance/Index'
import RoleManage from 'Page/RoleManage/Index'
import SelfData from 'Page/SelfData/Index'
import Shift from 'Page/Shift/Index'
import ShiftSetting from 'Page/ShiftSetting/Index'
import YearPerformance from 'Page/YearPerformance/Index'
import Login from 'Page/Login/Index'
import PrintForm from 'Page/Performance/PrintForm'
import PrintList from 'Page/Performance/PrintList'

// page func
const pageFuncCodePair = {
    banchManage: <BanchManage />,
    companyData: <CompanyData />,
    employeeManage: <EmployeeManage />,
    performance: <Performance />,
    roleManage: <RoleManage />,
    selfData: <SelfData />,
    shift: <Shift />,
    shiftSetting: <ShiftSetting />,
    yearPerformance: <YearPerformance />
}

const RouteIndex = (): JSX.Element => {
    const auth = useAppSelector((v) => v?.system?.auth)

    useEffect(() => {
        void api.system.auth()
        void api.system.getRoleBanchList()
    }, [])
    return (
        <Router>
            <Routes>
                <Route path='/' element={<Layout />}>
                    <Route path='entry/login' element={<Login />} />
                    {
                        auth?.menu?.map((item) => (
                            <Route
                                key={v4()}
                                path={item?.funcCode}
                                element={pageFuncCodePair?.[item?.funcCode]}
                            />
                        ))
                    }

                    {/* <Route path='/' element={<Navigate to={'entry/login'} />}/> */}
                    {/* <Route path='entry/:path' element={<Entry />} /> */}
                </Route>
                <Route path={'/performance/print/form'} element={<PrintForm />} />
                <Route path={'/performance/print/list'} element={<PrintList />} />
                <Route path='*' element={<Navigate to={'/'} />} />
            </Routes>
        </Router>
    )
}
export default RouteIndex
