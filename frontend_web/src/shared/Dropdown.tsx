import React from 'react'
import { Space, Dropdown as AntdDropdown, MenuProps, Typography } from 'antd'
import { modalType } from 'static'

interface props {
    menu?: MenuProps['items']
    onSelect?: (v: modalType) => void
}

const Dropdown = ({ menu = [], onSelect }: props): JSX.Element => {
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
                • • •
                </Space>
            </Typography.Link>
        </AntdDropdown>
    )
}
export default Dropdown
export {
    type props as DropdownProps
}
