import {
    useRef, useEffect, useState
} from 'react'

import { useCallBackStateProps } from './types'

const useCallbackState: useCallBackStateProps.indexProps = <T, >(
    initValue: T,
    options: useCallBackStateProps.options
) => {
    const [data, setData] = useState<T>(initValue)
    const callbackRef = useRef<Array<useCallBackStateProps.callbackRefProps<T>>>([])
    const setDataFunc = (
        v: T | ((va: T) => T),
        callback?: useCallBackStateProps.callbackRefProps<T>
    ): void => {
        // 這裡使用 不使用 非同步來 防止 react batch, 會影響效能
        // 目前發現只有在 兩層 的時候會 嘴後一個 callback 會被 吃掉
        callbackRef.current = [...callbackRef.current, callback]
        // console.log(callback)
        setData(v)
    }
    useEffect(() => {
        // 這邊加上 call back 是否只要執行一次
        if (options?.allowCallbackCombined) {
            if (callbackRef.current?.length > 0) {
                const actionFunc = callbackRef.current?.[callbackRef.current.length - 1]
                if (actionFunc) actionFunc(data)
            }
            return
        }
        callbackRef.current.forEach((item) => {
            // 這邊使用非同步 來防止 react batch
            setTimeout(() => {
                if (item) item(data)
            }, 0)
        })
        callbackRef.current = []
    }, [data])
    return [data, setDataFunc]
}
export {
    useCallbackState,
    type useCallBackStateProps as useCallBackStateTypes
}
