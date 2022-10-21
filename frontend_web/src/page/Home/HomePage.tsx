import { Tabs } from 'antd'
import React from 'react'
const HomePage = (): JSX.Element => {
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
