import React, {
    useEffect, useMemo, useRef
} from 'react'
import { Modal as AntdModal } from 'antd'
import { useBreakPoint } from 'hook/useBreakPoint'
import modal from '../types'
import { modalEvent } from '../method'
import { ModalContext, useModalContext } from '../context/useModalContext'
import { useCallbackState } from 'hook/useCallBackState'

const Footer = ({
    children
}: { children: (v: any) => JSX.Element }): JSX.Element => {
    const { open, ...other } = useModalContext()
    return (
        open && (
            <div className="ant-modal-footer">
                {children(other)}
            </div>
        )
    )
}

const realModal = <T, H>({
    Children,
    modalEmmitter,
    uid,
    hocProps,
    ...attr
}: modal.realModalProps<T, H>): JSX.Element => {
    const [open, setOpen] = useCallbackState(false)
    const { isLess, isMore, width: breakPointWidth } = useBreakPoint()
    const showAttrRef = useRef<modal.modalInfoProps<T>>()

    const onClose = async (): Promise<void> => {
        return await new Promise<void>((resolve, reject) => {
            setOpen(false, () => {
                resolve()
            })
        })
    }
    const onOpen = async (): Promise<void> => {
        return await new Promise<void>((resolve, reject) => {
            setOpen(true, () => {
                resolve()
            })
        })
    }
    const emmiterHandler = (v: modal.modalProps<T>): void => {
        if (v) {
            showAttrRef.current = {
                ...v,
                onClose,
                onOpen
            }
        }
        if (v?.open) void onOpen()
        if (!v?.open) void onClose()
    }

    // modal props
    const modalProps = useMemo(() => {
        let width = isLess('md') ? '95vw' : '500px'
        let title = ''
        let footer = null
        // 寬度
        if (attr?.width) width = attr?.width(isLess, isMore)
        if (showAttrRef.current?.width) width = showAttrRef.current?.width(isLess, isMore)

        // title
        if (attr?.title) title = attr?.title(showAttrRef.current)
        if (showAttrRef.current?.title) title = showAttrRef.current?.title(showAttrRef.current)

        // footer
        if (attr?.footer) footer = attr?.footer(showAttrRef.current)
        if (showAttrRef.current?.footer) footer = showAttrRef.current?.footer(showAttrRef.current)

        return {
            ...attr,
            ...showAttrRef.current,
            title,
            footer,
            onClose,
            onOpen,
            open,
            width,
            bodyStyle: {
                maxWidth: '100vw',
                minWidth: '250px',
                ...(showAttrRef.current?.bodyStyle || {}),
                ...(attr?.bodyStyle || {})
            }
        }
    }, [open, JSON.stringify(attr), JSON.stringify(showAttrRef.current), breakPointWidth])

    useEffect((): any => {
        modalEmmitter.on(modalEvent.show + uid, emmiterHandler)
        return () => {
            modalEmmitter.off(modalEvent.show + uid, onClose)
        }
    }, [])

    if (attr?.noModal) {
        return (
            <ModalContext.Provider value={modalProps}>
                <Children {...(hocProps) as H} />
            </ModalContext.Provider>
        )
    }

    return (
        <AntdModal
            {...modalProps}
            destroyOnClose
            mask
            onCancel={onClose}
            footer={null}
        >
            <ModalContext.Provider value={modalProps}>

                <Children {...(hocProps) as H} modalInfo={modalProps} />
                {
                    modalProps.footer && (
                        <div className="ant-modal-footer">
                            {modalProps.footer}
                        </div>
                    )
                }

            </ModalContext.Provider>
        </AntdModal>
    )
}

export {
    realModal,
    Footer
}
