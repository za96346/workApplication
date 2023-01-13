import Menu from 'Share/Menu'
import React, { useEffect, useState } from 'react'
import { Outlet, useLocation, useNavigate, useSearchParams } from 'react-router-dom'
import useReduceing from 'Hook/useReducing'
import { useBreakPoint } from 'Hook/useBreakPoint'
import { Drawer } from 'antd'
import Header from 'Share/Header'

const Layout = (): JSX.Element => {
    const [show, setShow] = useState(false)
    const { isLess } = useBreakPoint()
    const { pathname } = useLocation()
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [searchP, setSearchP] = useSearchParams()
    const navigate = useNavigate()
    const { user } = useReduceing()

    useEffect(() => {
        if (!user.token) {
            navigate('/entry/login')
        }
    }, [user.token])
    console.log('state => ', searchP.get('state'))
    console.log('scope => ', searchP.get('scope'))
    console.log('code => ', searchP.get('code'))
    console.log('error => ', searchP.get('error'))
    useEffect(() => {
        setShow(false)
    }, [navigate])

    if (pathname.includes('entry')) {
        return (
            <div className={window.styles.entryLayOut}>
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
            <div translate='no' className={window.styles.article}>
                {
                    isLess('md')
                        ? <Drawer onClose={() => setShow(false)} open={show}>
                            <Menu />
                        </Drawer>
                        : <Menu />
                }
                <div className={window.styles.rightBlock}>
                    <Outlet />
                </div>
            </div>
        </>
    )
}
export default Layout
