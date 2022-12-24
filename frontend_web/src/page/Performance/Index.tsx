import React, { useEffect, useMemo, useRef } from 'react'
import { Input, Table, Form, Button } from 'antd'
import useReduceing from '../../Hook/useReducing'
import { columns } from './method/columns'
import api from '../../api/api'
import { useParams } from 'react-router-dom'
import { SearchOutlined } from '@ant-design/icons'
import dateHandle from '../../method/dateHandle'

const Index = (): JSX.Element => {
    const { company } = useReduceing()
    const { banchId } = useParams()
    const formData = useRef({
        year: dateHandle.getMingQuar(),
        month: new Date().getMonth()
    })
    const convertBanchId = parseInt(banchId.replace('c', ''))
    const performance = useMemo(() => {
        return company.performance.map((item) => ({ ...item, Goal: item.Goal.replace('', '') }))
    }, [company.performance])
    const onSearch = (): void => {
        void api.getPerformance({
            banchId: convertBanchId,
            year: formData.current.year,
            month: formData.current.month
        })
    }
    useEffect(() => {
        void api.getPerformance({
            banchId: convertBanchId
        })
    }, [convertBanchId])
    return (
        <>
            <div className={window.styles.performanceBlock}>
                <Form onValuesChange={(v, allV) => { formData.current = allV }} className='row'>
                    <Form.Item name='year' className='col-md-3'>
                        <Input placeholder='年度（民國）' />
                    </Form.Item>
                    <Form.Item name='month' className='col-md-3'>
                        <Input placeholder='月份'/>
                    </Form.Item>
                    <div className='d-flex w-100 justify-content-end'>
                        <Button onClick={onSearch} icon={<SearchOutlined />}>
                        搜尋
                        </Button>
                    </div>
                </Form>
            </div>

            <Table
                style={{
                    fontSize: '0.5rem'
                }}
                dataSource={performance}
                columns={columns}
            />
        </>
    )
}
export default Index
