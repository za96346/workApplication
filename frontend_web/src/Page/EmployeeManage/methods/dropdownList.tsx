import React from 'react'
import { EditOutlined } from '@ant-design/icons'
import { modalType } from 'static'
import { DropdownProps } from 'shared/Dropdown'
import { type usePermissionTypes } from 'hook/usePermission'

export const dropdownList = (permission: usePermissionTypes.returnType): DropdownProps['menu'] => [
    permission.isEditable && {
        icon: <EditOutlined />,
        key: modalType.edit,
        label: '編輯'
    },
    permission.isDeleteable && {
        icon: <EditOutlined />,
        key: modalType.delete,
        label: '刪除'
    }
]
