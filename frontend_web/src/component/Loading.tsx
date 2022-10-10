import React from 'react'
import { Spin } from 'antd'

const Loading = (): JSX.Element => {
    return (
        <div style={{ width: '100%', height: '100%', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
            <Spin size="large" tip={'數據加載中'} />
        </div>
    )
}

export default Loading
