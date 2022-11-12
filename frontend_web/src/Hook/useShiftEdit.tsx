/* eslint-disable @typescript-eslint/prefer-includes */
import { message } from 'antd'
import { useEffect, useRef } from 'react'
import { useDispatch } from 'react-redux'
import shiftEditAction from '../reduxer/action/shiftEditAction'
import { OnlineUserType, ShiftSocketType } from '../type'
import { useWebsocket } from './useWebsocket'
interface props {
    connectionStatus: string
    sendMessage: Function
    close: Function
    lastJsonMessage: ShiftSocketType
}
const useShiftEditSocket = (banchId: number, token: string): props => {
    const dispatch = useDispatch()
    const record = useRef<OnlineUserType[]>([])
    const { connectionStatus, sendMessage, lastJsonMessage, socket } = useWebsocket({
        options: {
            onClose: (event: any) => {
                console.log(event)
            },
            queryParams: {
                banchId,
                token
            }
        }
    })
    const close = (): void => {
        socket?.close()
    }
    useEffect(() => {
        dispatch(shiftEditAction.setShiftEditOnlineUser(lastJsonMessage?.OnlineUser))
        if (!record.current) return
        // 找到新進的使用者
        lastJsonMessage?.OnlineUser?.forEach((item) => {
            const a = record.current.find((i) => i?.UserId === item?.UserId)
            if (!a) {
                message.info(`${item.UserName} 進入編輯室`)
            }
        })

        record.current?.forEach((item) => {
            const a = lastJsonMessage?.OnlineUser?.find((i) => i.UserId === item.UserId)
            if (!a) {
                message.info(`${item.UserName} 離開編輯室`)
            }
        })

        record.current = lastJsonMessage?.OnlineUser || []
    }, [lastJsonMessage])
    // useEffect(() => {
    //     firstEnter.current = true
    // }, [])
    return {
        connectionStatus, sendMessage, lastJsonMessage, close
    }
}
export default useShiftEditSocket
