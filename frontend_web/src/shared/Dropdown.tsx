import React from 'react'
import { DownOutlined } from '@ant-design/icons'
import { Space, Dropdown as AntdDropdown, MenuProps } from 'antd'

interface props {
    menu: MenuProps['items']
}

const Dropdown = ({ menu = [] }: props): JSX.Element => {
    return (
        <AntdDropdown menu={{ items: menu }}>
            <a onClick={(e) => e.preventDefault()}>
                <Space>
                    . . .
                    <DownOutlined />
                </Space>
            </a>
        </AntdDropdown>
    )
}
export default Dropdown
