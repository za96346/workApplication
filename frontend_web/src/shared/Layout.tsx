import React from 'react'
import { Outlet, useLocation } from 'react-router-dom'
import Menu from './Menu'
import Header from './Header'
import { useAppSelector } from 'hook/redux'

const Layout = (): JSX.Element => {
    const location = useLocation()
    const menu = useAppSelector((v) => v?.system?.auth?.menu)

    const currentPage = menu?.find((item) => `/${item?.funcCode}` === location?.pathname)

    if (location.pathname === '/entry/login') {
        return (
            <div translate='no' className='layout login'>
                <Outlet />
            </div>
        )
    }

    return (
        <div translate='no' className='layout'>
            <Menu />
            <div className='w-100 h-100'>
                <Header />
                <div className={'article'}>
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
