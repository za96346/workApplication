import { modalType } from 'static'

declare global {
    type Flag = 'Y' | 'N'
    declare module '*.svg'
    declare module '*.png'
    declare module '*.jpg'
    declare module '*.jpeg'
    declare module '*.gif'
    declare module '*.bmp'
    declare module '*.tiff'

    interface BtnEventParams<T> {
        type?: modalType
        value?: T
        reload?: Function
    }
}
export {}
