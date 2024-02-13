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
                    FuncCode: 'logout',
                    FuncName: '登出',
                    Sort: 100
                }
            ]
                ?.sort((a, b) => a?.Sort - b?.Sort)
                ?.map((item) => ({
                    key: item?.FuncCode,
                    label: item?.FuncName
                }))
        )
    }, [auth])

    return (
        <div
            style={{
                left: sidebarOpen
                    ? '0px'
                    : '-256px'
            }}
            className='
                menu
                h-100
                d-flex
                flex-column
                align-items-center
            '
        >
            <img src={pic} />
            <Divider />
            <AntdMenu
                onClick={(v) => {
                    if (v?.key === 'logout') {
                        navigate('entry/login')
                    } else {
                        navigate(v?.key)
                    }
                }}
                selectedKeys={[defaultSelectedKeys]}
                mode="inline"
                items={menuItem}
            />
        </div>
    )
}
export default Menu
