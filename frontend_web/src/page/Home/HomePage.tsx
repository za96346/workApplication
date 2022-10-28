import { Tabs } from 'antd'
import React, { useEffect } from 'react'
import api from '../../api/api'
const HomePage = (): JSX.Element => {
    useEffect(() => {
        api.getSelfData()
    }, [])
    return (
        <div className={styles.HomeBlock}>
            <Tabs>
                <Tabs.TabPane tab='加入公司' key={1}>

                </Tabs.TabPane>
                <Tabs.TabPane tab='創建公司' key={2}>

                </Tabs.TabPane>
            </Tabs>
        </div>
    )
}
export default HomePage
