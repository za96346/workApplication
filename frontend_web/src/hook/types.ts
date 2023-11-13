declare namespace useBreakPointProps {
    type WidthPropsType = 'xxs' | 'xs' | 'sm' | 'md' | 'lg' | 'xl' | 'xxl' | 'sxl'
    interface props {
        breakPoint?: string
        isLess: (widthProps: WidthPropsType) => boolean
        isMore: (widthProps: WidthPropsType) => boolean
        width: number
    }
}

declare namespace useWindowSizeProps {
    interface WindowSize {
        width: number
        height: number
    }
}

declare namespace useCallBackStateProps {
    type callbackRefProps<T> = ((v: T) => void)
    interface options {
        allowCallbackCombined: boolean
    }
    // eslint-disable-next-line no-shadow
    type indexProps = <T>(initValue: T, options?: options) => (
        [
            T, (
                (
                    v: T | ((va: T) => T),
                    callback?: callbackRefProps<T>,
                    a?: any
                ) => void
            )
        ]
    )
}

export {
    useBreakPointProps,
    useWindowSizeProps,
    useCallBackStateProps
}
