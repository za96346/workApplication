import { useState } from 'react'

import { useEventListener, useIsomorphicLayoutEffect } from 'usehooks-ts'
import { useWindowSizeProps } from './types'

const useWindowSize = (): useWindowSizeProps.WindowSize => {
    const [windowSize, setWindowSize] = useState<useWindowSizeProps.WindowSize>({
        width: 0,
        height: 0
    })

    const handleSize = (): void => {
        setWindowSize({
            width: window.innerWidth,
            height: window.innerHeight
        })
    }

    useEventListener('resize', handleSize)

    // Set size at the first client-side load
    useIsomorphicLayoutEffect(() => {
        handleSize()
    }, [])

    return windowSize
}

export {
    useWindowSize,
    type useWindowSizeProps as useWindowSizeTypes
}
