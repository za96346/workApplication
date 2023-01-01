/* eslint-disable @typescript-eslint/restrict-plus-operands */
/* eslint-disable @typescript-eslint/dot-notation */
import React, { useEffect, useMemo, useRef, useState, lazy, Suspense } from 'react'
import { Table, Form, Button, Space, Dropdown, Typography, Input, Spin, Divider, Modal } from 'antd'
import useReduceing from 'Hook/useReducing'
import { columns } from './method/columns'
import api from 'api/api'
import { useParams } from 'react-router-dom'
import { DeleteOutlined, DownOutlined, EditOutlined, SearchOutlined, SwitcherOutlined } from '@ant-design/icons'
import DateSelect from './component/DateSelect'
import dayjs from 'dayjs'
import { MenuItemType } from 'antd/es/menu/hooks/useItems'
import { performanceType } from 'type'
import StatusSelector from 'Share/StatusSelector'

const EditModal = lazy(async () => await import('./component/modal/Edit'))
const ChangeBanchModal = lazy(async () => await import('./component/modal/ChangeBanch'))
const list = (disabled: boolean): MenuItemType[] => [
    {
        label: '編輯',
        key: 1,
        icon: <EditOutlined />
    },
    {
        label: '刪除',
        key: 2,
        icon: <DeleteOutlined/>,
        disabled
    },
    {
        label: '更換組別',
        key: 3,
        icon: <SwitcherOutlined />,
        disabled
    }
]
const modalStateInit: {
    open: boolean
    type: string
    value: performanceType
} = {
    open: false,
    type: '-1',
    value: null
}

const Index = (): JSX.Element => {
    const { company, user } = useReduceing()
    const { banchId } = useParams()
    const [modal, setModal] = useState(modalStateInit)

    const formData = useRef({
        name: '',
        range: [dayjs(), dayjs()],
        workState: 'on'
    })
    const convertBanchId = parseInt(banchId.replace('c', ''))
    const disabled = (value: performanceType): boolean => user.selfData.Permession === 2 ||
        (user.selfData.Permession === 1 &&
        value.UserId === user.selfData.UserId)
    const performance = useMemo(() => {
        return company.performance.map((item) => ({
            ...item,
            Goal: item.Goal.replace('', ''),
            action: (
                <>
                    <Dropdown
                        menu={{
                            items: list(disabled(item)),
                            selectable: true,
                            onSelect: (v) => {
                                setModal((prev) => ({
                                    ...prev,
                                    open: true,
                                    type: v.key,
                                    value: item
                                }))
                            },
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
    const onSearch = async (): Promise<void> => {
        console.log(formData.current.range)
        await api.getPerformance({
            banchId: convertBanchId,
            name: formData.current.name,
            startYear: formData.current.range[0]?.year() - 1911,
            startMonth: formData.current.range[0]?.month() + 1,
            endYear: formData.current.range[1]?.year() - 1911,
            endMonth: formData.current.range[1]?.month() + 1,
            workState: formData.current.workState || 'on'
        })
    }
    const onClose = (): void => {
        setModal(modalStateInit)
    }
    const onEditSave = async (v: performanceType): Promise<void> => {
        let res = null
        if (modal.type === '4') {
            res = await api.createPerformance(v)
        } else {
            res = await api.updatePerformance(v)
        }
        if (res?.status) {
            await onSearch()
            onClose()
        }
    }
    useEffect(() => {
        onSearch()
    }, [convertBanchId])
    useEffect(() => {
        if (modal.type === '2') {
            Modal.confirm({
                okText: '確認',
                cancelText: '取消',
                title: "警告",
                content: <>
                    是否刪除<br/>
                        姓名： {modal.value.UserName}<br/>
                        年度： {modal.value.Year}<br/>
                        日期： {modal.value.Month}<br/>
                </>,
                onOk: async () => {
                    const res = await api.deletePerformance(modal.value.PerformanceId)
                    if (res.status) {
                        void onSearch()
                        onClose()
                    }
                },
                onCancel: onClose
            })
        }
    }, [modal])
    return (
        <>
            <Suspense fallback={<Spin />}>
                {
                    (modal.type === '1' || modal.type === '4') && (
                        <EditModal type={modal.type} onSave={onEditSave} onClose={onClose} value={modal.value} open={modal.open}/>
                    )
                }
                {
                    modal.type === '3' && (
                        <ChangeBanchModal onSave={onEditSave} onClose={onClose} value={modal.value} open={modal.open}/>
                    )
                }
            </Suspense>
            <div className={window.styles.performanceBlock}>
                <Button
                    onClick={() => {
                        setModal((prev) => ({
                            ...prev,
                            open: true,
                            type: '4',
                            value: null
                        }))
                    }}
                >
                    新增
                </Button>
                <Divider />
                <Form onValuesChange={(v, allV) => { formData.current = allV }} className='row'>
                    <Form.Item label='範圍搜尋' name='range' className='col-md-4'>
                        <DateSelect onChange={(v) => { }} />
                    </Form.Item>
                    <Form.Item label='姓名' name='name' className='col-md-4'>
                        <Input placeholder='輸入姓名'/>
                    </Form.Item>
                    <Form.Item label='狀態' name='workState' className='col-md-4'>
                        <StatusSelector defaultValue={'on'}/>
                    </Form.Item>
                    <div className='d-flex w-100 justify-content-end'>
                        <Button htmlType='submit' onClick={onSearch} icon={<SearchOutlined />}>
                        搜尋/reload
                        </Button>
                    </div>
                </Form>
                <Divider />
                <Button onClick={() => { window.open('/print', '績效評核', 'height=800,width=800') }}>
                    預覽 / 列印
                </Button>
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
