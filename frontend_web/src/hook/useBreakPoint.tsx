import { useWindowSize } from './useWindowSize'
import { useBreakPointProps } from './types'
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
    sxl: 1100,
    xl: 1200,
    xxl: 1600
}

const useBreakPoint = (): useBreakPointProps.props => {
    const { width } = useWindowSize()
    // console.log(width)
    const isMore = (widthProps: useBreakPointProps.WidthPropsType): boolean => {
        if (width > breakP[widthProps]) {
            return true
        }
        return false
    }

    const isLess = (widthProps: useBreakPointProps.WidthPropsType): boolean => {
        if (width <= breakP[widthProps]) {
            return true
        }
        return false
    }

    return { isLess, isMore, width }
}

export {
    useBreakPoint,
    type useBreakPointProps as useBreakPointTypes
}
