import Menu from '../component/Menu'
import React, { useEffect, useState } from 'react'
import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import useReduceing from '../Hook/useReducing'
import { useBreakPoint } from '../Hook/useBreakPoint'
import { Drawer } from 'antd'
import Header from '../component/Header'

const Layout = (): JSX.Element => {
    const [show, setShow] = useState(false)
    const { isLess } = useBreakPoint()
    const { pathname } = useLocation()
    const navigate = useNavigate()
    const { user } = useReduceing()

    useEffect(() => {
        if (!user.token) {
            navigate('/entry/login')
        }
    }, [user.token])

    useEffect(() => {
        setShow(false)
    }, [navigate])

    if (pathname.includes('entry')) {
        return (
            <div className={styles.entryLayOut}>
                <Outlet />
            </div>
        )
    }
    return (
        <>
            {
                isLess('md') && (
                    <Header setShow={setShow} />
                )
            }
            <div className={styles.article}>
                {
                    isLess('md')
                        ? <Drawer onClose={() => setShow(false)} open={show}>
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
