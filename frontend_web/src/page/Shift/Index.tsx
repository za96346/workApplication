/* eslint-disable @typescript-eslint/strict-boolean-expressions */
import { HolderOutlined } from '@ant-design/icons'
import { Drawer, Tabs } from 'antd'
import React, { useState } from 'react'
import useReduceing from 'Hook/useReducing'
import EditTable from './component/tabEdit/EditTable'
import PeopleStatus from './component/PeopleStatus'
import TabHistory from './component/tabHistory/Index'

const Index = (): JSX.Element => {
    const { shiftEdit } = useReduceing()
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
