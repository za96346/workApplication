/* eslint-disable @typescript-eslint/prefer-includes */
import { FullMessage } from 'method/notice'
import { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import shiftEditAction from 'reduxer/action/shiftEditAction'
import { ShiftSocketType } from '../../../type'
import { useWebsocket } from 'Hook/useWebsocket'
interface props {
    connectionStatus: string
    sendMessage: (v: string) => void
    close: Function
    lastJsonMessage: ShiftSocketType
}
const useShiftEditSocket = (banchId: number, token: string): props => {
    const dispatch = useDispatch()
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
        dispatch(shiftEditAction.setShiftEdit(lastJsonMessage))
        if (lastJsonMessage?.NewEntering?.length > 0) {
            FullMessage.info(`${lastJsonMessage?.NewEntering} 進入編輯室`)
        }
        if (lastJsonMessage?.NewLeaving?.length > 0) {
            FullMessage.info(`${lastJsonMessage?.NewLeaving} 離開編輯室`)
        }
        if (lastJsonMessage?.State?.errorMsg?.length > 0) {
            FullMessage.error(lastJsonMessage?.State?.errorMsg || '')
        }
    }, [lastJsonMessage])
    useEffect(() => {
        return () => {
            dispatch(shiftEditAction.clearShiftEdit())
        }
    }, [])
    return {
        connectionStatus, sendMessage, lastJsonMessage, close
    }
}
export default useShiftEditSocket