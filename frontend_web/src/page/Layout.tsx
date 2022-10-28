import Menu from '../component/Menu'
import React, { useEffect } from 'react'
import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import useReduceing from '../Hook/useReducing'

const Layout = (): JSX.Element => {
    const { pathname } = useLocation()
    const navigate = useNavigate()
    const { user } = useReduceing()

    useEffect(() => {
        if (!user.token) {
            navigate('/entry/login')
        }
    }, [user.token])

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
