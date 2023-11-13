import React from 'react'
import EventEmitter from 'eventemitter3'
import { v4 } from 'uuid'
import modal from './types'
import { modalEvent } from './method'
import { Footer, realModal } from './component/RealModal'
import { useModalContext } from './context/useModalContext'

/**
 * @T 為 open props
 * @CV 為 children type
*/
const Modal = <T, CV>({
    children,
    uid = v4(),
    ...attr
}: modal.indexProps<T, CV>): modal.indexReturnProps<T, CV> => {
    const modalEmmitter = new EventEmitter()
    type H = React.ComponentProps<typeof children>
    const fun = ({
        noModal,
        ...hocProps
    }: Pick<modal.indexProps<T, CV>, 'noModal'> & H): JSX.Element => realModal<any, H>({
        ...attr,
        Children: children as (v: H) => JSX.Element,
        modalEmmitter,
        uid,
        noModal,
        hocProps: hocProps as H
    })
    fun.open = (attrs?: modal.modalProps<T>) => {
        modalEmmitter.emit(modalEvent.show + uid, {
            ...(attrs || {}),
            open: true
        })
    }
    fun.close = (attrs?: modal.modalProps<T>) => {
        modalEmmitter.emit(modalEvent.show + uid, {
            ...(attrs || {}),
            open: false
        })
    }
    fun.Footer = Footer
    return fun
}

Modal.Footer = Footer

export {
    Modal,
    useModalContext as useModal,
    type modal as modalTypes
}
