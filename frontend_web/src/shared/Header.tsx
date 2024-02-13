import { MenuOutlined, UserOutlined } from '@ant-design/icons'
import { useAppDispatch, useAppSelector } from 'hook/redux'
import React from 'react'
import Dropdown from './Dropdown'
import { useNavigate } from 'react-router-dom'

const userDropdown = [
    {
        key: 'selfData',
        label: '基本資料'
    },
    {
        key: 'logout',
        label: '登出'
    }
]

const Header = (): JSX.Element => {
    const navigate = useNavigate()
    const { dispatch, action } = useAppDispatch()
    const sidebarOpen = useAppSelector((v) => v?.system.sidebar)
    return (
        <div
            style={!sidebarOpen
                ? {
                    width: '100%',
                    left: 0
                }
                : {}}
            className={`
                header
                d-flex
                justify-content-between
                align-item-center
            `}
        >
            <MenuOutlined
                style={{ fontSize: 20, cursor: 'pointer' }}
                onClick={() => {
                    dispatch(action.system.setSidebar(!sidebarOpen))
                }}
            />
            <Dropdown
                onSelect={(v: any) => {
                    console.log(v)
                    if (v === 'logout') {
                        navigate('/entry/login')
                        return
                    }
                    navigate(v)
                }}
                menu={userDropdown}
            >
                <UserOutlined
                    style={{ fontSize: 20, cursor: 'pointer' }}
                />
            </Dropdown>
        </div>
    )
}
export default Header
