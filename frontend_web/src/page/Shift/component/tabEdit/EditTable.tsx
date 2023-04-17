import { Skeleton, Spin, Collapse, Button } from 'antd'
import List from 'rc-virtual-list'
import React, { useEffect, useMemo } from 'react'
import useReduceing from 'Hook/useReducing'
import { useNavigate } from 'react-router-dom'
import { FullScreen, useFullScreenHandle } from 'react-full-screen'
import Row from './Row'
import dateHandle from 'method/dateHandle'
import Head from './Head'
import useShiftEditSocket from '../../Hook/useShiftEdit'
import statics from 'statics'
import api from 'api/api'
import { Dialog } from '@vteam_components/cloud'
import moment from 'moment'

const EditTable = (): JSX.Element => {
    const navigate = useNavigate()
    const fullScreenHandle = useFullScreenHandle()
    const { loading, user, shiftEdit, state } = useReduceing()
    const { sendMessage, close, connectionStatus } = useShiftEditSocket(state.banchId, user?.token || '')

    // 日期
    const dayArray = useMemo(() => {
        const start = shiftEdit?.StartDay
        const end = shiftEdit?.EndDay
        const dayContainer = ['']
        for (let i = 0; i < new Date(end).getDate(); i++) {
            const res = dateHandle.addDays(start, i)
            dayContainer.push(res)
        }
        return [...dayContainer, '總時數']
    }, [shiftEdit?.StartDay, shiftEdit?.EndDay])

    // 寄送確認編輯
    const onClickSubmit = (): void => {
        Dialog.warning({
            text: '請確認 是否 提交班表，提交後無法重新編輯。'
        }).then(() => {
            sendMessage(JSON.stringify({
                Types: statics.shiftSocketEvent.done
            }))
        })
    }

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
            <Collapse
                ghost
                className='mb-4'
                defaultActiveKey={['2']}
                onChange={(v) => {
                    if (v.includes('1')) {
                        api.getShiftHistory({
                            banch: state.banchId,
                            year: 2023,
                            month: 5
                        })
                    }
                }}
            >
                <Collapse.Panel header="編輯歷程" key="1">
                    <List
                        itemKey={''}
                        height={400}
                        itemHeight={30}
                        data={shiftEdit?.history || []}
                        className="list-group"
                    >
                        {
                            (item) => (
                                <div className='list-group-item list-group-item-action flex-column align-items-start' key={item.LogId}>
                                    <div className="d-flex w-100 justify-content-between">
                                        <span>{item?.Msg || ''}</span>
                                        <small className="text-muted">
                                            {moment(item?.LastModify).utcOffset(0).format('YYYY-MM-DD HH:mm').toString()}
                                        </small>
                                    </div>

                                </div>
                            )
                        }
                    </List>
                </Collapse.Panel>
                <Collapse.Panel header="圖標" key="2">
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
                                <Button onClick={onClickSubmit}>確認無誤 ， 提交班表</Button>
                            )
                        }
                    </FullScreen>
            }

        </>
    )
}
export default EditTable
