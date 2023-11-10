import api from 'api/Index'
import React, { useEffect } from 'react'
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Navigate
} from 'react-router-dom'
import Layout from 'shared/Layout'

const RouteIndex = (): JSX.Element => {
    useEffect(() => {
        api.auth.get()
    }, [])
    return (
        <Router>
            <Routes>
                <Route path='/' element={<Layout />}>

                    {/* <Route path='/' element={<Navigate to={'entry/login'} />}/> */}
                    {/* <Route path='entry/:path' element={<Entry />} /> */}
                </Route>
                {/* <Route path='*' element={<ErrorPage/>} /> */}
            </Routes>
        </Router>
    )
}
export default RouteIndex
