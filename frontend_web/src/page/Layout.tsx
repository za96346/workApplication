import Menu from '../component/Menu'
import React, { useEffect } from 'react'
import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import useReduceing from '../Hook/useReducing'
import { useBreakPoint } from '../Hook/useBreakPoint'
import { Drawer } from 'antd'

const Layout = (): JSX.Element => {
    const { isLess } = useBreakPoint()
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
                {
                    isLess('md')
                        ? <Drawer open>
                            <Menu />
                        </Drawer>
                        : <Menu />
                }
                <div className={styles.rightBlock}>
                    <Outlet />
                </div>
            </div>
        </>
    )
}
export default Layout
