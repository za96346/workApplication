/* eslint-disable @typescript-eslint/strict-boolean-expressions */
import { HolderOutlined } from '@ant-design/icons'
import { Drawer, Steps, Tabs } from 'antd'
import React, { useState } from 'react'
import useReduceing from 'Hook/useReducing'
import EditTable from './component/tabEdit/EditTable'
import PeopleStatus from './component/PeopleStatus'
import TabHistory from './component/tabHistory/Index'

const Index = (): JSX.Element => {
    const { company, shiftEdit, state } = useReduceing()
    const [status, setStatus] = useState({
        drawerOpen: false,
        currentTabs: 0
    })
    return (
        <>
            {
                !status.drawerOpen && (
                    <div
                        onClick={() => setStatus((prev) => ({ ...prev, drawerOpen: true }))}
                        className={window.styles.peopleListBtn}
                    >
                        <HolderOutlined />
                    </div>
                )
            }
            <Drawer
                title="目前上線"
                placement="right"
                closable
                mask
                maskClosable
                onClose={() => setStatus((prev) => ({ ...prev, drawerOpen: false }))}
                open={status.drawerOpen}
            >
                {
                    shiftEdit?.OnlineUser?.map((item, index) => {
                        return (
                            <PeopleStatus name={item.UserName} color={item.Color} key={index} currentStatus='online'/>
                        )
                    })
                }
            </Drawer>
            <div className={window.styles.shiftProcessBar}>
                <h3>{company.banch.find((item) => item.Id === state.banchId)?.BanchName || ''}</h3>
                <Steps current={shiftEdit.Status - 1}>
                    <Steps.Step title="尚未開放編輯" />
                    <Steps.Step title="開放編輯" subTitle="" description={`${shiftEdit.StartDay} ～～ ${shiftEdit.EndDay}`} />
                    <Steps.Step title="部門主管確認班表無誤" description="進行中..." />
                    <Steps.Step title="編輯完成" description="" status={shiftEdit.Status === 4 ? 'finish' : 'wait' } />
                </Steps>
            </div>
            <div className={window.styles.shiftEdit}>
                <Tabs
                    items={[
                        {
                            key: '0',
                            label: '當前編輯',
                            children: <EditTable />
                        },
                        {
                            key: '1',
                            label: '歷史班表',
                            children: <TabHistory />
                        }
                    ]}
                    destroyInactiveTabPane
                    onChange={(key) => setStatus((prev) => ({ ...prev, currentTabs: parseInt(key, 10) }))}
                />
            </div>
        </>
    )
}
export default Index
