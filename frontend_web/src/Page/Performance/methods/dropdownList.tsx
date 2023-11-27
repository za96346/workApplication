import React from 'react'
import { CopyOutlined, EditOutlined, ExceptionOutlined } from '@ant-design/icons'
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
    },
    permission?.isCopyable({ banchId: item?.BanchId, roleId: item?.RoleId }) && {
        icon: <CopyOutlined />,
        key: modalType.copy,
        label: '複製'
    },
    permission?.isChangeBanchable({ banchId: item?.BanchId }) && {
        icon: <ExceptionOutlined />,
        key: modalType.changeBanch,
        label: '更換部門'
    }
]
