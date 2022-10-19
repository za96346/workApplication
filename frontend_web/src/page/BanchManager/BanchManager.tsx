import { PlusOutlined } from '@ant-design/icons'
import { Card, Input, List, Form, Button } from 'antd'
import React, { useState } from 'react'
import { useSelector } from 'react-redux'
import api from '../../api/api'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'
import { RootState } from '../../reduxer/store'
import { BanchType } from '../../type'

interface people {
    title: string
    peopleAmount: number
}

const data = (arr: BanchType[]): people[] => arr.map((item) => {
    return (
        {
            title: item.BanchName,
            peopleAmount: 0
        }
    )
})

const BanchManager: React.FC = () => {
    const company: companyReducerType = useSelector((state: RootState) => state.company)
    const [status, setStatus] = useState({
        plusOnclick: false
    })
    const onFinish = async (v: any, types: 1 | 2): Promise<void> => {
        console.log(v)
        const { BanchName } = v
        const res = await api.createBanch(BanchName)
        if (res.status) {
            await api.getBanch()
        }
    }
    return (
        <div className={styles.banchManagerBlock}>
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
                        <Card title={item.title}>人數：{item.peopleAmount}</Card>
                    </List.Item>
                )}
            />
            {
                status.plusOnclick
                    ? <Form onFinish={async (v) => await onFinish(v, 1)}>
                        <Form.Item name="BanchName" label="新增部門" initialValue="">
                            <Input placeholder='請輸入部門名稱' />
                        </Form.Item>
                        <Form.Item>
                            <div style={{ display: 'flex', alignItems: 'flex-end', justifyContent: 'flex-end', width: '100%' }}>
                                <Button htmlType='submit'>新增</Button>
                            </div>
                        </Form.Item>
                    </Form>
                    : <PlusOutlined onClick={() => setStatus((prev) => ({ ...prev, plusOnclick: true }))} style={{ fontSize: '2rem' }} />
            }
        </div>
    )
}

export default BanchManager
