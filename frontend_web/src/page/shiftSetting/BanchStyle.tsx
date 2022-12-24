import { MinusOutlined, PlusOutlined } from '@ant-design/icons'
import { Button, Collapse, List, Modal, Skeleton } from 'antd'
import React, { useEffect, useState } from 'react'
import api from '../../api/api'
import Btn from '../../Share/Btn'
import useReduceing from '../../Hook/useReducing'
import dateHandle from '../../method/dateHandle'
import { BanchStyleType, ResType, ShiftSettingListType } from '../../type'
import EditForm from './EditForm'

const statusInit = {
    currentListIdx: -1,
    currentTabIdx: '-1',
    currentDeleteListIdx: -1,
    openModal: false
}

const data = (arr: BanchStyleType[]): ShiftSettingListType[] => arr?.map((item) => {
    return {
        id: `${item.StyleId}`,
        title: item.TimeRangeName,
        icons: item.Icon,
        time: <>
            <span>上班： {item.OnShiftTime}</span><br/>
            <span>下班： {item.OffShiftTime}</span><br/>
            <span>休息： {item.RestTime}</span>
        </>
    }
})

interface props {
    banchId: number
}
const BanchStyle = ({ banchId }: props): JSX.Element => {
    const [status, setStatus] = useState({ ...statusInit })
    const { loading, company } = useReduceing()

    const onDelete = async (): Promise<any> => {
        const res = await api.deleteBanchStyle(status.currentDeleteListIdx)
        if (res.status) {
            await api.getBanchStyle(banchId)
        }
        setStatus(statusInit)
    }

    const onEdit = (id: string): any => {
        setStatus((prev) => ({ ...prev, currentListIdx: parseInt(id) }))
    }

    const onFinish = async (v: any, types: 0 | 1): Promise<void> => {
        console.log(v, status.currentListIdx)
        let res: ResType<BanchStyleType>
        if (types === 0) {
            res = await api.createBanchStyle(
                {
                    ...v,
                    OnShiftTime: dateHandle.dateFormatToTime(v.OnShiftTime._d),
                    OffShiftTime: dateHandle.dateFormatToTime(v.OffShiftTime._d),
                    RestTime: dateHandle.dateFormatToTime(v.RestTime._d),
                    BanchId: banchId
                }
            )
        } else if (types === 1) {
            res = await api.updateBanchStyle(
                {
                    ...v,
                    OnShiftTime: dateHandle.dateFormatToTime(v.OnShiftTime._d),
                    OffShiftTime: dateHandle.dateFormatToTime(v.OffShiftTime._d),
                    RestTime: dateHandle.dateFormatToTime(v.RestTime._d),
                    StyleId: status.currentListIdx
                }
            )
        }
        if (res.status) {
            await api.getBanchStyle(banchId)
            setStatus({ ...statusInit })
        }
    }
    useEffect(() => {
        api.getBanchStyle(banchId)
    }, [banchId])

    useEffect(() => {
        console.log(status.currentListIdx)
        if (status.currentDeleteListIdx !== -1 || status.currentListIdx !== -1) {
            setStatus((prev) => ({ ...prev, openModal: true }))
        }
    }, [status.currentDeleteListIdx, status.currentListIdx])
    return (
        <>
            <Modal
                forceRender
                onOk={() => setStatus({ ...statusInit })}
                onCancel={() => setStatus({ ...statusInit })}
                open={status.openModal}
                footer={null}
            >
                {
                    status.currentDeleteListIdx !== -1 && <><div>
                        是否刪除此圖標，刪除後無法復原
                    </div>
                    <Button onClick={onDelete} block style={{ marginTop: '20px' }}>確認</Button>
                    </>
                }
                {
                    status.currentListIdx !== -1 && <EditForm currentId={status.currentListIdx} types={0} btnText='確認修改' onFinish={async (v) => await onFinish(v, 1)} />
                }
            </Modal>
            <div>
                <Collapse
                    onChange={(e) => setStatus((prev) => ({ ...prev, currentTabIdx: e[e.length - 1] }))}
                    expandIcon={() => status.currentTabIdx === '1' ? <MinusOutlined /> : <PlusOutlined />}
                    style={{ width: '100%' }}
                    defaultActiveKey={[status.currentTabIdx]}
                >
                    <Collapse.Panel header="新增" key={'1'}>
                        <EditForm types={0} onFinish={async (v) => await onFinish(v, 0)} />
                    </Collapse.Panel>
                </Collapse>
            </div>
            <List
                style={{ marginTop: '20px' }}
                split
                loading={loading.onFetchBanchStyle}
                itemLayout="horizontal"
                dataSource={data(company.banchStyle)}
                renderItem={(item: ShiftSettingListType) => (
                    loading.onFetchBanchStyle
                        ? <Skeleton active />
                        : <List.Item
                            style={
                                (status.currentDeleteListIdx === -1 ||
                                    status.currentDeleteListIdx === parseInt(item.id)
                                ) && (
                                    status.currentListIdx === -1 ||
                                    status.currentListIdx === parseInt(item.id)
                                )
                                    ? {
                                        opacity: 1,
                                        flexWrap: 'wrap',
                                        position: 'relative'
                                    }
                                    : {
                                        opacity: 0.1,
                                        flexWrap: 'wrap',
                                        position: 'relative'
                                    }
                            }
                        >
                            <List.Item.Meta
                                avatar={<span style={{ fontSize: '2rem' }} >{item.icons}</span>}
                                title={item.title}
                                description={item.time}
                                style={{ minWidth: '200px', marginBottom: '30px' }}
                            />

                            <div style={{
                                position: 'absolute',
                                right: '0px',
                                bottom: '5px'
                            }}>
                                <Btn.Edit onClick={() => onEdit(item.id)}/>
                                <Btn.Delete onClick={() => setStatus((prev) => ({ ...prev, currentDeleteListIdx: parseInt(item.id) }))} />
                            </div>

                        </List.Item>
                )}
            />
        </>
    )
}
export default BanchStyle
