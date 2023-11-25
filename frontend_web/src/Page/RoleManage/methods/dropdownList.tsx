import React from 'react'
import { EditOutlined } from '@ant-design/icons'
import { modalType } from 'static'
import { DropdownProps } from 'shared/Dropdown'
import { type usePermissionTypes } from 'hook/usePermission'
import roleTypes from 'types/role'

export const dropdownList = (
    permission: usePermissionTypes.returnType,
    item: roleTypes.TABLE
): DropdownProps['menu'] => [
    permission.isEditable({ roleId: item?.RoleId }) && {
        icon: <EditOutlined />,
        key: modalType.edit,
        label: '編輯'
    },
    permission.isDeleteable({ roleId: item?.RoleId }) && {
        icon: <EditOutlined />,
        key: modalType.delete,
        label: '刪除'
    }
]
