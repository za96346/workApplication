import { ModalProps } from 'antd'
import { ReactNode } from 'react'
import { useBreakPointProps } from 'hook/types'

declare namespace modal {
    type modalProps<T> = T
    & Omit<ModalProps, 'children' | 'width' | 'title' | 'footer'>
    & {
        type?: string
        value?: any
        noModal?: boolean
        width?: (
            isLess: useBreakPointProps.props['isLess'],
            isMore: useBreakPointProps.props['isMore']
        ) => string
        title?: (modalProps: modalProps<T>) => any
        footer?: (modalProps: modalProps<T>) => any
    }
    type modalInfoProps<T> = modalProps<Partial<T>> & {
        onClose?: () => Promise<void>
        onOpen?: () => Promise<void>
    }

    type indexProps<T, CV> = modalProps<Partial<T>> & {
        children: (v: CV) => JSX.Element
        uid?: string
    }
    interface indexReturnProps<T, H> {
        (v: Pick<indexProps<T, H>, 'noModal'> & H): JSX.Element
        open: (v: Partial<modalProps<T>>) => void
        close: (v: Partial<modalProps<T>>) => void
    }
    type realModalProps<T, H> = modalProps<Partial<T>> & {
        Children: React.ComponentType<H & { modalInfo?: modal.modalInfoProps<T> }>
        footer?: any
        modalEmmitter: any
        uid: string
        hocProps: Partial<H>
        children?: ReactNode
    }
}
export default modal
