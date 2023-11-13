import { createContext, useContext } from 'react'
import modal from '../types'

const ModalContext = createContext<any>({})

const useModalContext = <T, >(): modal.modalProps<T> => {
    const data = useContext<modal.modalProps<T>>(ModalContext)
    return {
        ...data
    }
}
export {
    ModalContext,
    useModalContext
}
