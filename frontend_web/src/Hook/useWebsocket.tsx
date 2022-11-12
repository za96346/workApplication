import useWebSocket, { ReadyState } from 'react-use-websocket'
import { WebSocketLike } from 'react-use-websocket/dist/lib/types'

const DEFAULT_URL = `${process.env.REACT_APP_SOCKET}/shift`

interface Websocket {
    options?: any
    BASE_URL?: string
}

interface props {
    connectionStatus: string
    sendMessage: Function
    lastJsonMessage: any
    socket: WebSocketLike
}

export const useWebsocket = ({ options, BASE_URL = DEFAULT_URL }: Websocket): props => {
    const { readyState, sendMessage, lastJsonMessage, getWebSocket } = useWebSocket(`${BASE_URL}`, {
        ...options,
        shouldReconnect: () => true,
        reconnectAttempts: 10,
        reconnectInterval: 3000,
        onReconnectStop: (e) => console.log('==== websocket stop ====', e),
        onClose: () => console.log('socket close')

    // onOpen: () => {},
    // onError: () => {},
    })
    const socket = getWebSocket()
    const connectionStatus = {
        [ReadyState.CONNECTING]: 'Connecting',
        [ReadyState.OPEN]: 'Open',
        [ReadyState.CLOSING]: 'Closing',
        [ReadyState.CLOSED]: 'Closed',
        [ReadyState.UNINSTANTIATED]: 'Uninstantiated'
    }[readyState]

    return { connectionStatus, sendMessage, lastJsonMessage, socket }
}
