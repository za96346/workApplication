import { DeleteOutlined, EditOutlined, MinusOutlined, PlusOutlined } from '@ant-design/icons'
import { Button, Collapse, List, Modal, Skeleton } from 'antd'
import React, { useEffect, useState } from 'react'
import { useSelector } from 'react-redux'
import api from '../../api/api'
import dateHandle from '../../method/dateHandle'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'
import { statusReducerType } from '../../reduxer/reducer/statusReducer'
import { RootState } from '../../reduxer/store'
import { BanchStyleType, ShiftSettingListType } from '../../type'
import EditForm from './EditForm'

const statusInit = {
    currentListIdx: -1,
    currentTabIdx: '-1',
    currentDeleteListIdx: -1,
    openModal: false
}

const data = (arr: BanchStyleType[]): ShiftSettingListType[] => arr.map((item) => {
    return {
        id: `${item.StyleId}`,
        title: item.TimeRangeName,
        icons: item.Icon,
        time: <>
            <span>上班： {item.OnShiftTime}</span>
            <span style={{ marginLeft: '10px' }}>下班： {item.OffShiftTime}</span>
        </>
    }
})

interface props {
    banchId: number
}
const BanchStyle = ({ banchId }: props): JSX.Element => {
    const [status, setStatus] = useState({ ...statusInit })
    const company: companyReducerType = useSelector((state: RootState) => state.company)
    const loading: statusReducerType = useSelector((state: RootState) => state.status)
    const onDelete = (idx: string): any => {
        setStatus((prev) => ({ ...prev, currentDeleteListIdx: parseInt(idx), openModal: true }))
    }
    const onEdit = (id: string): any => {
        setStatus((prev) => ({ ...prev, currentListIdx: parseInt(id), openModal: true }))
    }
    const onFinish = async (v: any, types: 0 | 1): Promise<void> => {
        console.log(v)
        let res
        if (types === 0) {
            res = await api.createBanchStyle(
                {
                    ...v,
                    OnShiftTime: dateHandle.dateFormatToTime(v.OnShiftTime._d),
                    OffShiftTime: dateHandle.dateFormatToTime(v.OffShiftTime._d),
                    BanchId: banchId
                }
            )
        } else if (types === 1) {
            res = await api.updateBanchStyle(
                {
                    ...v,
                    OnShiftTime: dateHandle.dateFormatToTime(v.OnShiftTime._d),
                    OffShiftTime: dateHandle.dateFormatToTime(v.OffShiftTime._d),
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
    return (
        <>
            <Modal
                onOk={() => setStatus({ ...statusInit })}
                onCancel={() => setStatus({ ...statusInit })}
                open={status.openModal}
                footer={null}
            >
                {
                    status.currentDeleteListIdx !== -1 && <><div>
                        是否刪除此圖標，刪除後無法復原
                    </div>
                    <Button block style={{ marginTop: '20px' }}>確認</Button>
                    </>
                }
                {
                    status.currentListIdx !== -1 && <EditForm types={0} btnText='確認修改' onFinish={async (v) => await onFinish(v, 1)} />
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
            <div style={{ marginTop: '20px' }}>
                <List
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
                                            opacity: 1
                                        }
                                        : {
                                            opacity: 0.1
                                        }
                                }
                            >
                                <List.Item.Meta
                                    avatar={<span style={{ fontSize: '2rem' }} >{item.icons}</span>}
                                    title={item.title}
                                    description={item.time}
                                />

                                <div
                                    onClick={() => onEdit(item.id)}
                                    className={styles.editLabel}
                                    style={{ color: 'blue' }}
                                >
                                    <EditOutlined style={{ marginRight: '10px' }} />
                                    編輯
                                </div>

                                <div
                                    onClick={() => onDelete(item.id)}
                                    className={styles.editLabel}
                                    style={{ marginLeft: '20px', color: 'red' }}
                                >
                                    <DeleteOutlined style={{ marginRight: '10px' }} />
                                        刪除
                                </div>
                            </List.Item>
                    )}
                />
            </div>
        </>
    )
}
export default BanchStyle
