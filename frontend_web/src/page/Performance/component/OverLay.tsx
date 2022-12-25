import { DeleteOutlined, EditOutlined, SwitcherOutlined } from '@ant-design/icons'
import { Menu } from 'antd'
import React from 'react'
import { v4 } from 'uuid'
const list = [
    {
        title: '編輯',
        key: v4(),
        icon: <EditOutlined />,
        color: 'text-primary'
    },
    {
        title: '刪除',
        key: v4(),
        icon: <DeleteOutlined/>,
        color: 'text-danger'
    },
    {
        title: '更換組別',
        key: v4(),
        icon: <SwitcherOutlined />,
        color: 'text-secondary'
    }
]

const OverLay = (): JSX.Element => {
    return (
        <>
            <Menu className='list-group'>
                {list.map(category => (
                    <Menu.Item key={category.key}>
                        <a
                            className={`list-group-item list-group-item-action border-0 w-100 ${category.color}`}
                            href="#"
                            onClick={(event) => { event.preventDefault() }}
                        >
                            {category.icon}
                            {category.title}
                        </a>
                    </Menu.Item>
                ))}
            </Menu>
        </>
    )
}
export default OverLay
