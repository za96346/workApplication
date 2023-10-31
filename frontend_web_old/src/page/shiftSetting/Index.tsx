import { Tabs } from 'antd'
import React from 'react'
import useReduceing from 'Hook/useReducing'
import BanchStyle from './component/banchIcon/Index'
import EditShift from './component/editShift/Index'

const ShiftSettingPage = (): JSX.Element => {
    const { company, state } = useReduceing()
    const banch = company.banch.find((item) => item?.Id === state.banchId)

    return (
        <>

            <Tabs
                destroyInactiveTabPane
                className={window.styles.shiftSettingBlock}
            >
                <Tabs.TabPane tab={'圖標設定'} key={0}>
                    <div style={{ marginBottom: '20px', fontSize: '1.2rem' }}>{banch?.BanchName}</div>
                    <BanchStyle banchId={state.banchId} />
                </Tabs.TabPane>
                <Tabs.TabPane tab={'排班設定'} key={1}>
                    <EditShift />
                </Tabs.TabPane>
            </Tabs>
        </>
    )
}
export default ShiftSettingPage
