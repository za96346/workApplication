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

    type inferFirstArray<T> = T extends Array<infer U> ? U : never

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
