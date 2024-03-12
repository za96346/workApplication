import React from 'react'
import { EyeFilled } from '@ant-design/icons'
import { modalType } from 'static'
import { DropdownProps } from 'shared/Dropdown'
import { type usePermissionTypes } from 'hook/usePermission'

export const dropdownList = (
    permission: usePermissionTypes.returnType
): DropdownProps['menu'] => [
    permission.isInquirable && {
        icon: <EyeFilled />,
        key: modalType.inquire,
        label: '檢視'
    }
]
