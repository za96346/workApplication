import React, { useEffect } from 'react'
import { Modal } from 'shared/Modal/Index'
import modal from 'shared/Modal/types'
import Perofrmance from 'Page/Performance/Index'

interface modalInfo {
    userId?: number
    userName?: string
    year?: number
    onDestroy?: () => void
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const ModalPerformance = ({ modalInfo }: props): JSX.Element => {
    useEffect(() => {
        return () => {
            if (modalInfo?.onDestroy) modalInfo?.onDestroy()
        }
    }, [])
    return (
        <>
            <Perofrmance userId={modalInfo?.userId} year={modalInfo?.year} />
        </>
    )
}
export default Modal<modalInfo, any>({
    children: ModalPerformance,
    title: (v) => (
        <>
            檢視
            <span className='text-danger'>{v?.year}</span>
            年度
            <span className='text-danger'>{v?.userName}</span>
            每月績效
        </>
    ),
    width: (isLess) => '100vw'
})
