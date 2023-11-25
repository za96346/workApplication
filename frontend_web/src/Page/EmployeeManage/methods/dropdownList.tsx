import React from 'react'
import { EditOutlined } from '@ant-design/icons'
import { modalType } from 'static'
import { DropdownProps } from 'shared/Dropdown'
import { type usePermissionTypes } from 'hook/usePermission'
import userTypes from 'types/user'

export const dropdownList = (
    permission: usePermissionTypes.returnType,
    item: userTypes.TABLE
): DropdownProps['menu'] => [
    permission.isEditable({ banchId: item?.BanchId, roleId: item?.RoleId }) && {
        icon: <EditOutlined />,
        key: modalType.edit,
        label: '編輯'
    },
    permission.isDeleteable({ banchId: item?.BanchId, roleId: item?.RoleId }) && {
        icon: <EditOutlined />,
        key: modalType.delete,
        label: '刪除'
    }
]
