import Session from 'func/Session'

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

declare namespace usePermissionProps {
    interface returnType {
        isEditable: ({ banchId, roleId }: { banchId?: number, roleId?: number }) => boolean
        isDeleteable: ({ banchId, roleId }: { banchId?: number, roleId?: number }) => boolean
        isCopyable: ({ banchId, roleId }: { banchId?: number, roleId?: number }) => boolean
        isChangeBanchable: ({ banchId }: { banchId?: number }) => boolean

        isInquirable: boolean
        isAddable: boolean
        isPrintable: boolean

    }
}

declare namespace useSessionProps {
    interface returnProps<T> {
        session: () => ReturnType<Session<T>['get']>
        setSession: Session<T>['set']
        backward: Session<T>['backward']
        forward: Session<T>['forward']
        reset: Session<T>['reset']
    }
    interface InstanceReturnProps<T> {
        Provider: React.FunctionComponent<{ children: any }>
        Instance: Session<T>
    }
}

declare namespace useFlagProps {
    interface returnProps {
        flagToDom: (v: {
            flag: Flag
            flagYText: any
            flagNText: any
        }) => JSX.Element
    }
}

export {
    type useBreakPointProps,
    type useWindowSizeProps,
    type useCallBackStateProps,
    type usePermissionProps,
    type useSessionProps,
    type useFlagProps
}
