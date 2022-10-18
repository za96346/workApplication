import { Tabs } from 'antd'
import React from 'react'
import { useParams } from 'react-router-dom'
import BanchRule from './BanchRule'
import BanchStyle from './BanchStyle'

const ShiftSettingPage = (): JSX.Element => {
    const { banchId } = useParams()
    const banchIdNumber = parseInt(banchId)

    return (
        <>
            <Tabs
                destroyInactiveTabPane
                className={styles.shiftSettingBlock}
            >
                <Tabs.TabPane tab={'圖標設定'} key={0}>
                    <BanchStyle banchId={banchIdNumber} />
                </Tabs.TabPane>
                <Tabs.TabPane tab="排班規則" key={1}>
                    <BanchRule banchId={banchIdNumber} />
                </Tabs.TabPane>
            </Tabs>
        </>
    )
}
export default ShiftSettingPage
