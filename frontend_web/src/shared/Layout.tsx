import React, { useEffect } from 'react'
import { Outlet, useLocation } from 'react-router-dom'
import Menu from './Menu'
import Header from './Header'
import { useAppSelector } from 'hook/redux'
import { Loading } from './Loading/Index'
import api from 'api/Index'

const Layout = (): JSX.Element => {
    const location = useLocation()
    const menu = useAppSelector((v) => v?.system?.auth?.menu)
    const sidebarOpen = useAppSelector((v) => v?.system.sidebar)

    const currentPage = menu?.find((item) => `/${item?.FuncCode}` === location?.pathname)

    useEffect(() => {
        // 這邊先這樣打 兩次 api, 之後換後端
        if (location.pathname !== '/entry/login') {
            void Promise.all([
                api.system.auth(),
                api.system.getRoleBanchList()
            ]).then(() => {
                void api.system.auth()
                void api.system.getRoleBanchList()
            })
        }
    }, [location.pathname])

    if (location.pathname === '/entry/login') {
        return (
            <div translate='no' className='layout login'>
                <Outlet />
            </div>
        )
    }

    return (
        <div translate='no' className='layout'>
            <Loading />
            <Menu />
            <div
                className='w-100 h-100'
            >
                <Header />
                <div
                    className={'article'}
                    style={!sidebarOpen
                        ? {
                            width: '100%',
                            left: 0
                        }
                        : {}}
                >
                    <h4 className='text-secondary'>
                        {currentPage?.FuncName || ''}
                    </h4>
                    <Outlet />
                </div>
            </div>
        </div>
    )
}
export default Layout
