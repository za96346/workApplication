import Menu from '../component/Menu'
import React, { useEffect } from 'react'
import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import { useSelector } from 'react-redux'

const Layout = (): JSX.Element => {
    const { pathname } = useLocation()
    const navigate = useNavigate()
    const { token } = useSelector((state: any) => state.user)

    useEffect(() => {
        if (!token) {
            navigate('/entry/login')
        }
    }, [token, navigate])

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
