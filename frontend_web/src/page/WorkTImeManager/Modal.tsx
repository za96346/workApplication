import React, { useEffect } from 'react'
import api from '../../api/api'
import { Modal as AntdModal, props as antdModalprops } from '../../Share/Modal'

interface props extends antdModalprops {
}

const Modal = ({ ...attr }: props): JSX.Element => {
    useEffect(() => {
        api.createWorkTime({
            UserId: 2,
            Year: 2022,
            Month: 7,
            WorkHours: 8,
            TimeOff: 8,
            UsePaidVocation: 1
        })
    }, [])
    return (
        <AntdModal
            {...attr}
            destroyOnClose
            forceRender
        >

        </AntdModal>
    )
}
export default Modal
