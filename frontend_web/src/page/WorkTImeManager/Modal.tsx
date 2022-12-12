import React from 'react'
import { Modal as AntdModal, props as antdModalprops } from '../../component/Modal'

interface props extends antdModalprops {
}

const Modal = ({ ...attr }: props): JSX.Element => {
    return (
        <AntdModal
            {...attr}
            destroyOnClose
            forceRender
        >
            hi

        </AntdModal>
    )
}
export default Modal
