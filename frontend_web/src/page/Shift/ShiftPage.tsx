/* eslint-disable @typescript-eslint/strict-boolean-expressions */
import { HolderOutlined } from '@ant-design/icons'
import { Drawer } from 'antd'
import React, { useState } from 'react'
import { useParams } from 'react-router-dom'
import Loading from '../../component/Loading'
import PeopleStatus from './PeopleStatus'
const ShiftPage = (): JSX.Element => {
    const { banch } = useParams()
    const [status, setStatus] = useState({
        drawerOpen: false
    })
    if (!banch) {
        return (
            <Loading />
        )
    }
    return (
        <>
            {
                banch && (banch)
            }
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
        </>
    )
}
export default ShiftPage
