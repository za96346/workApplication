import React from 'react'
import { Modal as AntdModal, ModalProps } from 'antd'

export const Modal = ({ ...rest }: ModalProps): JSX.Element => {
    return (
        <AntdModal cancelText="取消" okText="確定" centered {...rest}>
            {rest.children}
        </AntdModal>
    )
}
