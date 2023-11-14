import React, { useMemo, useRef, useState, createContext, useContext } from 'react'
import Session from 'func/Session'
import { useSessionProps } from './types'

const SessionContext = createContext(null)

/**
 * @params option 為可選參數,
 * @params allowStorage = true 的時候 會把每次的更改資料以及 歷史資料存到 sessionStorage 並且以 id 為key
*/
const session = <T, >(
    initValue: ConstructorParameters<typeof Session<T>>['0'],
    option?: ConstructorParameters<typeof Session<T>>['1']
): useSessionProps.InstanceReturnProps<T> => {
    const Instance = new Session<T>(initValue, option ?? {})
    return {
        // eslint-disable-next-line react/prop-types
        Provider: ({ children }) => (
            <SessionContext.Provider value={Instance}>
                {children}
            </SessionContext.Provider>
        ),
        Instance
    }
}

// eslint-disable-next-line no-empty-pattern
const useSession = <T, >({}): useSessionProps.returnProps<T> => {
    // 紀錄被使用的 session key
    const usedSessionKeyRef = useRef([])
    const [, forceUpdate] = useState(0)
    const providerSession = useContext<Session<T>>(SessionContext)

    const sessionInstance = useMemo(() => {
        let previousValue = providerSession?.get()
        providerSession.on(
            providerSession.id,
            () => {
                const listenedValue = usedSessionKeyRef.current?.map((item) => {
                    return providerSession?.get()?.[item]
                })

                if (JSON.stringify(previousValue) !== JSON.stringify(listenedValue)) {
                    forceUpdate((prev) => prev + 1)
                }
                previousValue = (
                    listenedValue === undefined ||
                    typeof listenedValue === 'function'
                )
                    ? listenedValue
                    : JSON.parse(JSON.stringify(listenedValue))
            }
        )
        return providerSession
    }, [])
    return {
        session: () => sessionInstance.get((target, prop) => {
            usedSessionKeyRef.current.push(prop)
            usedSessionKeyRef.current = [...new Set(usedSessionKeyRef.current)]
        }),
        setSession: sessionInstance.set.bind(sessionInstance),
        backward: sessionInstance.backward.bind(sessionInstance),
        forward: sessionInstance.forward.bind(sessionInstance),
        reset: sessionInstance.reset.bind(sessionInstance)
    }
}
export {
    useSession,
    session
}
