import { MenuOutlined, UserOutlined } from '@ant-design/icons'
import { useAppDispatch, useAppSelector } from 'hook/redux'
import React from 'react'

const Header = (): JSX.Element => {
    const { dispatch, action } = useAppDispatch()
    const sidebarOpen = useAppSelector((v) => v?.system.sidebar)
    return (
        <div className='header d-flex justify-content-between align-item-center'>
            <MenuOutlined
                style={{ fontSize: 20, cursor: 'pointer' }}
                onClick={() => {
                    dispatch(action.system.setSidebar(!sidebarOpen))
                }}
            />
            <UserOutlined style={{ fontSize: 20, cursor: 'pointer' }} />
        </div>
    )
}
export default Header
