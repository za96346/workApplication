import { EditOutlined, PlusOutlined } from '@ant-design/icons'
import { Card, Input, List, Form, Button } from 'antd'
import React, { useEffect, useState } from 'react'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'
import rule from '../../method/rule'
import { BanchType, ResType } from '../../type'

interface people {
    title: string
    peopleAmount: number
    id: number
}

const data = (arr: BanchType[]): people[] => arr.map((item) => {
    return (
        {
            title: item.BanchName,
            peopleAmount: 0,
            id: item.Id
        }
    )
})

const BanchEdit = (
    { onFinish, label, btnName, initialValue }:
    {
        onFinish: (v: any) => void
        label: string
        btnName: string
        initialValue: string
    }): JSX.Element => {
    return (
        <>
            <Form onFinish={onFinish}>
                <Form.Item rules={rule.banch()} name="BanchName" label={label} initialValue={initialValue}>
                    <Input placeholder='請輸入部門名稱' />
                </Form.Item>
                <Form.Item>
                    <div style={{ display: 'flex', alignItems: 'flex-end', justifyContent: 'flex-end', width: '100%' }}>
                        <Button htmlType='submit'>{btnName}</Button>
                    </div>
                </Form.Item>
            </Form>
        </>
    )
}

const BanchManager: React.FC = () => {
    const { company } = useReduceing()
    const [status, setStatus] = useState({
        plusOnclick: false,
        currentEditIdx: -1
    })
    const onFinish = async (v: any, types: 1 | 2): Promise<void> => {
        let res: ResType<null>
        const { BanchName } = v
        if (types === 1) {
            res = await api.createBanch(BanchName)
        } else if (types === 2) {
            res = await api.UpdateBanch(BanchName, status.currentEditIdx)
        }
        if (res.status) {
            await api.getBanch()
        }
    }
    useEffect(() => {
        api.getBanch()
    }, [])
    return (
        <div className={window.styles.banchManagerBlock}>
            <List
                grid={{
                    gutter: 16,
                    xs: 1,
                    sm: 2,
                    md: 4,
                    lg: 4,
                    xl: 6,
                    xxl: 3
                }}
                dataSource={data(company.banch)}
                renderItem={item => (
                    <List.Item>
                        <Card title={
                            status.currentEditIdx === item.id
                                ? <BanchEdit
                                    initialValue={item.title}
                                    btnName='修改'
                                    label='更改部門'
                                    onFinish={async (v) => await onFinish(v, 2)}
                                />
                                : <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                                    {item.title}
                                    <EditOutlined style={{ color: 'blue' }} onClick={() => setStatus((prev) => ({ ...prev, currentEditIdx: item.id }))} />
                                </div>
                        }>
                            人數：{item.peopleAmount}
                        </Card>
                    </List.Item>
                )}
            />
            {
                status.plusOnclick
                    ? <BanchEdit
                        initialValue={''}
                        btnName='新增'
                        label="新增部門"
                        onFinish={async (v) => await onFinish(v, 1)}
                    />
                    : <PlusOutlined onClick={() => setStatus((prev) => ({ ...prev, plusOnclick: true }))} style={{ fontSize: '2rem' }} />
            }
        </div>
    )
}

export default BanchManager
