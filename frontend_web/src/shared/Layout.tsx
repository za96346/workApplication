import React, { Suspense, useEffect, useState } from 'react'
import { Outlet, useLocation, useNavigate, useSearchParams } from 'react-router-dom'
import { Drawer, Spin } from 'antd'

const Layout = (): JSX.Element => {
    return (
        <>
            rgwweffwf
            <div translate='no' className={''}>
                <div className={''}>
                    <Outlet />
                </div>
            </div>
        </>
    )
}
export default Layout
