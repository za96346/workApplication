import useWebSocket, { ReadyState } from 'react-use-websocket'

const DEFAULT_URL = `${process.env.REACT_APP_SOCKET}/shift`

interface Websocket {
    options?: any
    BASE_URL?: string
}

export const useWebsocket = ({ options, BASE_URL = DEFAULT_URL }: Websocket): any => {
    const { readyState, sendMessage, lastJsonMessage } = useWebSocket(`${BASE_URL}`, {
        ...options,
        shouldReconnect: () => true,
        reconnectAttempts: 10,
        reconnectInterval: 3000,
        onReconnectStop: (e) => console.log('==== websocket stop ====', e)
    // onOpen: () => {},
    // onError: () => {},
    // onClose: () => {},
    })

    const connectionStatus = {
        [ReadyState.CONNECTING]: 'Connecting',
        [ReadyState.OPEN]: 'Open',
        [ReadyState.CLOSING]: 'Closing',
        [ReadyState.CLOSED]: 'Closed',
        [ReadyState.UNINSTANTIATED]: 'Uninstantiated'
    }[readyState]

    return { connectionStatus, sendMessage, lastJsonMessage }
}
