import { modalType } from 'static'

declare global {
    const styles: Record<string, string>
    type Flag = 'Y' | 'N'
    declare module'*.scss'
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

    interface Window {
        styles: Record<string, string>
    }
}
export {}
