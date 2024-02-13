import React from 'react'
import { Space, Dropdown as AntdDropdown, MenuProps, Typography } from 'antd'
import { modalType } from 'static'

interface props {
    menu?: MenuProps['items']
    onSelect?: (v: modalType) => void
    children?: any
}

const Dropdown = ({ menu = [], onSelect, children }: props): JSX.Element => {
    return (
        <AntdDropdown
            menu={{
                items: menu,
                onSelect: (v) => { onSelect(v?.key as modalType) },
                selectable: true,
                selectedKeys: []
            }}
            trigger={['click']}
        >
            <Typography.Link
                onClick={(e) => {
                    e.stopPropagation()
                }}
            >
                <Space>
                    {children ?? '• • •'}
                </Space>
            </Typography.Link>
        </AntdDropdown>
    )
}
export default Dropdown
export {
    type props as DropdownProps
}
