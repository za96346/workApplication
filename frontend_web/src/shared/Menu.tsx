import React, { useMemo } from 'react'
import { Menu as AntdMenu, Divider } from 'antd'
import { useAppSelector } from 'hook/redux'
import pic from 'asserts/logo192.png'
import { useLocation, useNavigate } from 'react-router-dom'

const Menu = (): JSX.Element => {
    const auth = useAppSelector((v) => v?.system?.auth)
    const sidebarOpen = useAppSelector((v) => v?.system?.sidebar)
    const navigate = useNavigate()
    const location = useLocation()

    const defaultSelectedKeys = location.pathname?.replace('/', '')

    const menuItem = useMemo(() => {
        return (
            [
                ...(auth?.menu || []),
                {
                    funcCode: 'logout',
                    FuncName: '登出'
                }
            ]
                ?.map((item) => ({
                    key: item?.funcCode,
                    label: item?.FuncName
                }))
        )
    }, [auth])

    return (
        <div
            style={{
                width: sidebarOpen
                    ? '256px'
                    : '0px'
            }}
            className='
                menu
                h-100
                d-flex
                flex-column
                align-items-center
            '
        >
            <img src={pic} className={`${!sidebarOpen && 'd-none'}`} />
            <Divider />
            <AntdMenu
                onClick={(v) => {
                    if (v?.key === 'logout') {
                        navigate('entry/login')
                    } else {
                        navigate(v?.key)
                    }
                }}
                defaultSelectedKeys={[defaultSelectedKeys]}
                mode="inline"
                items={menuItem}
                className={`${!sidebarOpen && 'd-none'}`}
            />
        </div>
    )
}
export default Menu
