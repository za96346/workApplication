import React, { useState, useEffect } from 'react'
import { MutatingDots } from 'react-loader-spinner'
import { useAppSelector } from 'hook/redux'

/**
 * --mainColor = 主顏色
 * --secondaryColor = 副顏色
*/
const Loading = (): JSX.Element => {
    const [isLoading, setIsLoading] = useState(false)
    const loading = useAppSelector((v) => v?.loading)
    const mainColor = '#00FF80'
    const secondaryColor = '#3399FF'

    // 這邊要處理 loading 狀態
    useEffect(() => {
        const isFindLoading = Object.values(loading)
            .filter((item) => item && typeof item === 'boolean')
            ?.length > 0
        let timeOutRef: NodeJS.Timeout
        if (isFindLoading && !isLoading) {
            setIsLoading(true)
        } else {
            timeOutRef = setTimeout(() => {
                setIsLoading(false)
            }, 100)
        }
        return () => {
            if (timeOutRef) {
                clearTimeout(timeOutRef)
            }
        }
    }, [loading])

    return (
        <>
            {
                isLoading && (
                    <div
                        style={{
                            position: 'fixed',
                            top: '0px',
                            left: '0px',
                            right: '0px',
                            bottom: '0px',
                            zIndex: 1001,
                            backgroundColor: 'rgba(0,0,0,0)',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center'
                        }}
                    >
                        <MutatingDots
                            height="100"
                            width="100"
                            radius="12.5"
                            color={mainColor}
                            secondaryColor={secondaryColor}
                            ariaLabel="mutating-dots-loading"
                            wrapperStyle={{}}
                            wrapperClass=""
                            visible
                        />
                    </div>
                )
            }
        </>
    )
}
export { Loading }
