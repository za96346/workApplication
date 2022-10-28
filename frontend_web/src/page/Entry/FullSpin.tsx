import { Spin } from 'antd'
import React from 'react'
import useReduceing from '../../Hook/useReducing'

const FullSpin = (): JSX.Element => {
    const { loading } = useReduceing()
    if (loading.onEntry) {
        return (
            <>
                <div style={{
                    position: 'fixed',
                    top: 0,
                    bottom: 0,
                    left: 0,
                    right: 0,
                    backgroundColor: 'rgba(255,255,255,0.4)',
                    zIndex: 1000,
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center'
                }}>
                    <Spin />
                </div>
            </>
        )
    }
    return <></>
}
export default FullSpin
