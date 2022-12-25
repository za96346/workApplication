/* eslint-disable @typescript-eslint/dot-notation */
import React, { useEffect, useMemo, useRef } from 'react'
import { Table, Form, Button, Space, Dropdown, Typography, Input } from 'antd'
import useReduceing from 'Hook/useReducing'
import { columns } from './method/columns'
import api from 'api/api'
import { useParams } from 'react-router-dom'
import { DownOutlined, SearchOutlined } from '@ant-design/icons'
import OverLay from './component/OverLay'
import DateSelect from './component/DateSelect'
import dayjs from 'dayjs'

const Index = (): JSX.Element => {
    const { company } = useReduceing()
    const { banchId } = useParams()

    const formData = useRef({
        name: '',
        range: [dayjs(), dayjs()]
    })
    const convertBanchId = parseInt(banchId.replace('c', ''))
    const performance = useMemo(() => {
        return company.performance.map((item) => ({
            ...item,
            Goal: item.Goal.replace('', ''),
            action: (
                <>
                    <Dropdown
                        // menu={{
                        //     items,
                        //     selectable: true,
                        //     defaultSelectedKeys: ['3']
                        // }}
                        overlay={<OverLay />}
                    >
                        <Typography.Link>
                            <Space>
                                更多....
                                <DownOutlined />
                            </Space>
                        </Typography.Link>
                    </Dropdown>
                </>
            )
        }))
    }, [company.performance])
    const onSearch = (): void => {
        console.log(formData.current.range)
        void api.getPerformance({
            banchId: convertBanchId,
            name: formData.current.name,
            startYear: formData.current.range[0].year() - 1911,
            startMonth: formData.current.range[0].month() + 1,
            endYear: formData.current.range[1].year() - 1911,
            endMonth: formData.current.range[1].month() + 1
        })
    }
    useEffect(() => {
        onSearch()
    }, [convertBanchId])
    return (
        <>
            <div className={window.styles.performanceBlock}>
                <Form onValuesChange={(v, allV) => { formData.current = allV }} className='row'>
                    <Form.Item label='範圍搜尋' name='range' className='col-md-4'>
                        <DateSelect onChange={(v) => { }} />
                    </Form.Item>
                    <Form.Item label='姓名' name='name' className='col-md-4'>
                        <Input placeholder='輸入姓名'/>
                    </Form.Item>
                    <div className='d-flex w-100 justify-content-end'>
                        <Button htmlType='submit' onClick={onSearch} icon={<SearchOutlined />}>
                        搜尋
                        </Button>
                    </div>
                </Form>
            </div>

            <Table
                sticky={{ offsetHeader: -20 }}
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
