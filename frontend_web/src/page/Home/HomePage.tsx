import { Card } from 'antd'
import React from 'react'

const HomePage = (): JSX.Element => {
    return (
        <div style={{
            height: '100%',
            width: '100%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
        }}>
            <div style={{ display: 'flex', width: '100%', justifyContent: 'space-evenly' }}>
                <Card title="創建公司" bordered={false} style={{ width: 300 }}>
                </Card>
                <Card title="加入公司" bordered={false} style={{ width: 300 }}>
                </Card>
            </div>
        </div>
    )
}
export default HomePage
