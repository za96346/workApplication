import Menu from '../component/Menu'
import React, { useEffect } from 'react'
import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import { useSelector } from 'react-redux'
import api from '../api/api'
import { RootState } from '../reduxer/store'

const Layout = (): JSX.Element => {
    const { pathname } = useLocation()
    const navigate = useNavigate()
    const { token } = useSelector((state: RootState) => state.user)

    useEffect(() => {
        if (!token) {
            navigate('/entry/login')
        } else {
            void api.getBanch()
            void api.getBanchStyle(1)
            void api.getBanchRule(1)
        }
    }, [token])

    if (pathname.includes('entry')) {
        return (
            <div className={styles.entryLayOut}>
                <Outlet />
            </div>
        )
    }
    return (
        <>
            <div className={styles.article}>
                <Menu />
                <div className={styles.rightBlock}>
                    <Outlet />
                </div>
            </div>
        </>
    )
}
export default Layout
