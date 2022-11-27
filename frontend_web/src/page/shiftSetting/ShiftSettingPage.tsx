import { Tabs } from 'antd'
import React from 'react'
import { useParams } from 'react-router-dom'
import useReduceing from '../../Hook/useReducing'
// import BanchRule from './BanchRule'
import BanchStyle from './BanchStyle'

const ShiftSettingPage = (): JSX.Element => {
    const { banchId } = useParams()
    const { company } = useReduceing()
    const banchIdNumber = parseInt(banchId.replace('b', ''))
    const banch = company.banch.find((item) => item?.Id === banchIdNumber)

    return (
        <>

            <Tabs
                destroyInactiveTabPane
                className={styles.shiftSettingBlock}
            >
                <Tabs.TabPane tab={'圖標設定'} key={0}>
                    <div style={{ marginBottom: '20px', fontSize: '1.2rem' }}>{banch?.BanchName}</div>
                    <BanchStyle banchId={banchIdNumber} />
                </Tabs.TabPane>
                {/* <Tabs.TabPane tab="排班規則" key={1}>
                    <div style={{ marginBottom: '20px', fontSize: '1.2rem' }}>{banch?.BanchName}</div>
                    <BanchRule banchId={banchIdNumber} />
                </Tabs.TabPane> */}
            </Tabs>
        </>
    )
}
export default ShiftSettingPage
