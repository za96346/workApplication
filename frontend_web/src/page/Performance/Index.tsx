/* eslint-disable @typescript-eslint/dot-notation */
import React, { useEffect, useMemo, useRef, useState, lazy, Suspense } from 'react'
import { Table, Form, Button, Space, Dropdown, Typography, Input, Spin } from 'antd'
import useReduceing from 'Hook/useReducing'
import { columns } from './method/columns'
import api from 'api/api'
import { useParams } from 'react-router-dom'
import { DeleteOutlined, DownOutlined, EditOutlined, SearchOutlined, SwitcherOutlined } from '@ant-design/icons'
import DateSelect from './component/DateSelect'
import dayjs from 'dayjs'
import { MenuItemType } from 'antd/es/menu/hooks/useItems'
import swal from 'sweetalert'

const EditModal = lazy(async () => await import('./component/modal/Edit'))
const ChangeBanchModal = lazy(async () => await import('./component/modal/ChangeBanch'))
const list: MenuItemType[] = [
    {
        label: '編輯',
        key: 1,
        icon: <EditOutlined />
    },
    {
        label: '刪除',
        key: 2,
        icon: <DeleteOutlined/>
    },
    {
        label: '更換組別',
        key: 3,
        icon: <SwitcherOutlined />
    }
]
const modalStateInit = {
    open: false,
    type: '-1',
    value: null
}

const Index = (): JSX.Element => {
    const { company } = useReduceing()
    const { banchId } = useParams()
    const [modal, setModal] = useState(modalStateInit)

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
                        menu={{
                            items: list,
                            selectable: true,
                            onSelect: (v) => { setModal((prev) => ({ ...prev, open: true, type: v.key, value: item })) },
                            selectedKeys: [modal.type]
                        }}
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
    }, [company.performance, modal.type])
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
    const onClose = (): void => {
        setModal(modalStateInit)
    }
    useEffect(() => {
        onSearch()
    }, [convertBanchId])
    useEffect(() => {
        if (modal.type === '2') {
            swal({
                title: '警告',
                text: '是否刪除'
            }).then(() => { onClose() })
        }
    }, [modal])
    return (
        <>
            <Suspense fallback={<Spin />}>
                {
                    modal.type === '1' && (
                        <EditModal onClose={onClose} value={modal.value} open={modal.open}/>
                    )
                }
                {
                    modal.type === '3' && (
                        <ChangeBanchModal onClose={onClose} value={modal.value} open={modal.open}/>
                    )
                }
            </Suspense>
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
