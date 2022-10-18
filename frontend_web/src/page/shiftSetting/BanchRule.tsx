import { DeleteOutlined, EditOutlined, MinusOutlined, PlusOutlined } from '@ant-design/icons'
import { Collapse, List, Skeleton } from 'antd'
import React, { useEffect, useState } from 'react'
import { useSelector } from 'react-redux'
import api from '../../api/api'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'
import { statusReducerType } from '../../reduxer/reducer/statusReducer'
import { RootState } from '../../reduxer/store'
import { BanchRuleListType, BanchRuleType } from '../../type'
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
            <span>下班： {item.OffShiftTime}</span>
        </>
    }
})
interface props {
    banchId: number
}

const BanchRule = ({ banchId }: props): JSX.Element => {
    const [status, setStatus] = useState({ ...statusInit })
    const company: companyReducerType = useSelector((state: RootState) => state.company)
    const loading: statusReducerType = useSelector((state: RootState) => state.status)
    const onFinish = (v: any, types: number): any => {

    }
    const onEdit = (id: string): any => {
    }
    const onDelete = (id: string): any => {
    }
    useEffect(() => {
        api.getBanchRule(banchId)
    }, [banchId])
    return (
        <>
            <div>
                <Collapse
                    onChange={(e) => setStatus((prev) => ({ ...prev, currentTabIdx: e[e.length - 1] }))}
                    expandIcon={() => status.currentTabIdx === '1' ? <MinusOutlined /> : <PlusOutlined />}
                    style={{ width: '100%' }}
                    defaultActiveKey={[status.currentTabIdx]}
                >
                    <Collapse.Panel header="新增" key={'1'}>
                        <EditForm types={1} onFinish={(v) => onFinish(v, 0)} />
                    </Collapse.Panel>
                </Collapse>
            </div>
            <List
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
                                        opacity: 1
                                    }
                                    : {
                                        opacity: 0.1
                                    }
                            }
                        >
                            <List.Item.Meta
                                avatar={<span style={{ fontSize: '2rem' }} >{item.id}</span>}
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
        </>
    )
}
export default BanchRule
