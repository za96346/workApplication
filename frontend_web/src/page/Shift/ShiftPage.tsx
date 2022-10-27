/* eslint-disable @typescript-eslint/strict-boolean-expressions */
import { HolderOutlined } from '@ant-design/icons'
import { Drawer, Spin, Steps, Tabs } from 'antd'
import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'
import EditTable from './EditTable'
import PeopleStatus from './PeopleStatus'
const ShiftPage = (): JSX.Element => {
    const { banchId } = useParams()
    const { company } = useReduceing()
    const convertBanchId = parseInt(banchId)
    const [status, setStatus] = useState({
        drawerOpen: false,
        currentTabs: 0
    })

    useEffect(() => {
        api.getUserAll()
    }, [])
    if (!banchId) {
        return (
            <Spin />
        )
    }
    return (
        <>
            {
                !status.drawerOpen && (
                    <div
                        onClick={() => setStatus((prev) => ({ ...prev, drawerOpen: true }))}
                        className={styles.peopleListBtn}
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
                    new Array(20).fill('').map((item, index) => {
                        return (
                            <PeopleStatus key={index} currentStatus='online'/>
                        )
                    })
                }
            </Drawer>
            <div className={styles.shiftProcessBar}>
                <h3>{company.banch.find((item) => item.Id === convertBanchId)?.BanchName || ''}</h3>
                <Steps current={2}>
                    <Steps.Step title="開放編輯" description="2022-11-01 ~ 2022-11-09" />
                    <Steps.Step title="主管審核" subTitle="2022-11-10" description="部門主管確認班表無誤" />
                    <Steps.Step title="確認發布" description="進行中..." />
                </Steps>
            </div>
            <div className={styles.shiftEdit}>
                <Tabs onChange={(key) => setStatus((prev) => ({ ...prev, currentTabs: parseInt(key, 10) }))}>
                    <Tabs.TabPane tab={'當前編輯'} key={0}>
                        <EditTable
                            banchId={convertBanchId}
                            currentTabs={status.currentTabs}
                        />
                    </Tabs.TabPane>
                    <Tabs.TabPane tab={'歷史班表'} key={1}>
                        <EditTable
                            banchId={convertBanchId}
                            currentTabs={status.currentTabs}
                        />
                    </Tabs.TabPane>
                </Tabs>
            </div>
        </>
    )
}
export default ShiftPage
