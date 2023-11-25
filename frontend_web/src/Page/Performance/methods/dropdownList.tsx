import React from 'react'
import { EditOutlined } from '@ant-design/icons'
import { modalType } from 'static'
import { DropdownProps } from 'shared/Dropdown'
import { type usePermissionTypes } from 'hook/usePermission'
import performanceTypes from 'types/performance'

export const dropdownList = (
    permission: usePermissionTypes.returnType,
    item: inferFirstArray<performanceTypes.reducerType['all']>
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
