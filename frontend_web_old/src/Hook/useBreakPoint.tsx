import { createBreakpoint, useWindowSize } from 'react-use'

/**
 * @returns
  xs: '480px',
  sm: '576px',
  md: '768px',
  lg: '992px',
  xl: '1200px',
  xxl: '1600px',
 */

const breakP: any = {
    xxs: 330,
    xs: 480,
    sm: 576,
    md: 768,
    lg: 992,
    xl: 1200,
    xxl: 1600
}

type WidthPropsType = 'xxs' | 'xs' | 'sm' | 'md' | 'lg' | 'xl' | 'xxl'

interface props {
    breakPoint: string
    isLess: (widthProps: WidthPropsType) => boolean
    isMore: (widthProps: WidthPropsType) => boolean
}

export const useBreakPoint = (): props => {
    const { width } = useWindowSize()
    const breakPointFn = createBreakpoint({
        xxs: 330,
        xs: 480,
        sm: 576,
        md: 768,
        lg: 992,
        xl: 1200,
        xxl: 1600
    })

    const isMore = (widthProps: WidthPropsType): boolean => {
        if (width > breakP[widthProps]) {
            return true
        }
        return false
    }

    const isLess = (widthProps: WidthPropsType): boolean => {
        if (width <= breakP[widthProps]) {
            return true
        }
        return false
    }

    const breakPoint = breakPointFn()

    return { breakPoint, isLess, isMore }
}
