import { MinusOutlined, PlusOutlined } from '@ant-design/icons'
import { Button, Collapse, List, Modal, Skeleton } from 'antd'
import React, { useEffect, useState } from 'react'
import api from '../../api/api'
import Btn from '../../component/Btn'
import useReduceing from '../../Hook/useReducing'
import dateHandle from '../../method/dateHandle'
import statics from '../../statics'
import { BanchRuleListType, BanchRuleType, ResType } from '../../type'
import EditForm from './EditForm'

const statusInit = {
    currentListIdx: -1,
    currentTabIdx: '-1',
    currentDeleteListIdx: -1,
    openModal: false
}

const data = (arr: BanchRuleType[]): BanchRuleListType[] => arr.map((item) => {
    return {
        id: `${item.RuleId}`,
        title: `${item.WeekDay}`,
        time: <>
            <span>上班： {item.OnShiftTime}</span><br/>
            <span>下班： {item.OffShiftTime}</span><br/>
            <span>最少員工數：{item.MinPeople}</span><br/>
            <span>最多員工數：{item.MaxPeople}</span>
        </>,
        weekType: item.WeekType
    }
})
interface props {
    banchId: number
}

const BanchRule = ({ banchId }: props): JSX.Element => {
    const [status, setStatus] = useState({ ...statusInit })
    const { loading, company } = useReduceing()
    const onFinish = async (v: any, types: number): Promise<void> => {
        console.log(v)
        let res: ResType<BanchRuleType>
        if (types === 0) {
            res = await api.createBanchRule(
                {
                    ...v,
                    OnShiftTime: dateHandle.dateFormatToTime(v.OnShiftTime._d),
                    OffShiftTime: dateHandle.dateFormatToTime(v.OffShiftTime._d),
                    BanchId: banchId
                }
            )
        } else if (types === 1) {
            res = await api.updateBanchRule({
                ...v,
                OnShiftTime: dateHandle.dateFormatToTime(v.OnShiftTime._d),
                OffShiftTime: dateHandle.dateFormatToTime(v.OffShiftTime._d),
                RuleId: status.currentListIdx
            })
        }
        if (res.status) {
            await api.getBanchRule(banchId)
            setStatus({ ...statusInit })
        }
    }
    const onDelete = async (): Promise<any> => {
        const res = await api.deleteBanchRule(status.currentDeleteListIdx)
        if (res.status) {
            await api.getBanchRule(banchId)
        }
        setStatus(statusInit)
    }
    const onEdit = (id: string): any => {
        setStatus((prev) => ({ ...prev, currentListIdx: parseInt(id) }))
    }
    useEffect(() => {
        api.getBanchRule(banchId)
    }, [banchId])

    useEffect(() => {
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
                        是否刪除此規則，刪除後無法復原
                    </div>
                    <Button onClick={onDelete} block style={{ marginTop: '20px' }}>確認</Button>
                    </>
                }
                {
                    status.currentListIdx !== -1 && (
                        <EditForm
                            currentId={status.currentListIdx}
                            types={1}
                            btnText='確認修改'
                            onFinish={async (v) => await onFinish(v, 1)}
                        />
                    )
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
                        <EditForm types={1} onFinish={async (v) => await onFinish(v, 0)} />
                    </Collapse.Panel>
                </Collapse>
            </div>
            <List
                style={{ marginTop: '20px' }}
                split
                loading={loading.onFetchBanchRule}
                itemLayout="horizontal"
                dataSource={data(company.banchRule)}
                renderItem={(item: BanchRuleListType) => (
                    loading.onFetchBanchRule
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
                                        position: 'relative'
                                    }
                                    : {
                                        opacity: 0.1,
                                        position: 'relative'
                                    }
                            }
                        >
                            <List.Item.Meta
                                avatar={<span style={{ fontSize: '1rem', color: 'blue' }} >{statics.weekType[item.weekType]}</span>}
                                title={statics.weekDay[item.title]}
                                description={item.time}
                                style={{
                                    marginBottom: '30px'
                                }}
                            />
                            <div style={{
                                position: 'absolute',
                                right: '0px',
                                bottom: '5px'
                            }}>
                                <Btn.Edit onClick={() => onEdit(item.id)} />
                                <Btn.Delete onClick={async () => setStatus((prev) => ({ ...prev, currentDeleteListIdx: parseInt(item.id) }))} />
                            </div>
                        </List.Item>
                )}
            />
        </>
    )
}
export default BanchRule
