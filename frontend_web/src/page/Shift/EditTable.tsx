import { Result, Skeleton, Spin, Collapse, Button } from 'antd'
import React, { useEffect, useMemo } from 'react'
import useReduceing from 'Hook/useReducing'
import { useNavigate } from 'react-router-dom'
import { FullScreen, useFullScreenHandle } from 'react-full-screen'
import Row from './component.tsx/Row'
import dateHandle from 'method/dateHandle'
import Head from './component.tsx/Head'
import useShiftEditSocket from './Hook/useShiftEdit'

interface EditTableProps {
    currentTabs: number
    banchId: number
}
const EditTable = ({ banchId, currentTabs }: EditTableProps): JSX.Element => {
    const navigate = useNavigate()
    const fullScreenHandle = useFullScreenHandle()
    const { loading, user, shiftEdit } = useReduceing()
    const { sendMessage, close, connectionStatus } = useShiftEditSocket(banchId, user?.token || '')

    // 日期
    const dayArray = useMemo(() => {
        const start = shiftEdit?.StartDay
        const end = shiftEdit?.EndDay
        const dayContainer = ['']
        for (let i = 0; i < new Date(end).getDate(); i++) {
            const res = dateHandle.addDays(start, i)
            dayContainer.push(res)
        }
        return dayContainer
    }, [shiftEdit?.StartDay, shiftEdit?.EndDay])

    useEffect(() => {
        return () => {
            close()
        }
    }, [navigate])
    if (loading.onFetchBanchStyle || loading.onFetchUserAll) {
        return (
            <>
                <Spin />
                <Skeleton active />
            </>
        )
    }
    return (
        <>
            {
                currentTabs === 1 && (
                    <>
                        search bar
                        <Result
                            status="404"
                            title="404"
                            subTitle="找不到資料"
                        // extra={<Button type="primary">Back Home</Button>}
                        />
                    </>
                )
            }
            {
                currentTabs === 0 && (
                    <>

                        <Collapse ghost className='mb-4' defaultActiveKey={['1']} onChange={() => {}}>
                            <Collapse.Panel header="圖標" key="1">
                                <div className={window.styles.shiftSignBlock}>
                                    {
                                        shiftEdit.BanchStyle?.map((item) => {
                                            return (
                                                <div key={item.StyleId}>
                                                    <div>{item.TimeRangeName}</div>
                                                    {item.OnShiftTime} - {item.OffShiftTime}:
                                                    <span>{item.Icon}</span>
                                                </div>
                                            )
                                        })
                                    }
                                </div>
                            </Collapse.Panel>
                        </Collapse>

                        {
                            connectionStatus !== 'Connecting' &&
                            connectionStatus !== 'Open'
                                ? <Spin tip={'進入編輯室中...'} />
                                : <FullScreen handle={fullScreenHandle}>
                                    <Button onClick={() => { fullScreenHandle.enter() }}>全螢幕</Button>
                                    <div>排班日期：{shiftEdit?.StartDay}~{shiftEdit?.EndDay}</div>
                                    <table style={{ cursor: 'pointer' }} className='mb-5 table table-bordered table-hover table-striped table-responsive-md'>
                                        <thead>
                                            <Head dayArray={dayArray}/>
                                        </thead>
                                        <tbody>
                                            <Row sendMessage={sendMessage} dayArray={dayArray} />
                                        </tbody>
                                    </table>
                                    {
                                        shiftEdit?.State?.submitAble && (
                                            <Button>確認無誤 ， 提交班表</Button>
                                        )
                                    }
                                </FullScreen>

                        }

                    </>
                )
            }
        </>
    )
}
export default EditTable
