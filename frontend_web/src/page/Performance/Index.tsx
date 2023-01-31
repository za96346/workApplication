/* eslint-disable @typescript-eslint/restrict-plus-operands */
/* eslint-disable @typescript-eslint/dot-notation */
import React, { useEffect, useMemo, useRef, useState, lazy, Suspense } from 'react'
import { Table, Button, Space, Dropdown, Typography, Spin, Modal } from 'antd'
import useReduceing from 'Hook/useReducing'
import { columns } from './method/columns'
import api from 'api/api'
import { ArrowRightOutlined, CopyOutlined, DeleteOutlined, DownOutlined, EditOutlined, SwitcherOutlined } from '@ant-design/icons'
import { MenuItemType } from 'antd/es/menu/hooks/useItems'
import { performanceType } from 'type'
import SearchBar from './component/SearchBar'
import { v4 } from 'uuid'
import { FullMessage } from 'method/notice'

const EditModal = lazy(async () => await import('./component/modal/Edit'))
const ChangeBanchModal = lazy(async () => await import('./component/modal/ChangeBanch'))
const list = (disabled: boolean): MenuItemType[] => [
    {
        label: '編輯',
        key: 1,
        icon: <EditOutlined />
    },
    {
        label: '更換組別',
        key: 3,
        icon: <SwitcherOutlined />,
        disabled
    },
    {
        label: '複製績效',
        key: 5,
        icon: <CopyOutlined />,
        disabled
    },
    {
        label: '刪除',
        key: 2,
        icon: <DeleteOutlined/>,
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
    const { company, user, state } = useReduceing()
    const [modal, setModal] = useState(modalStateInit)
    const reFetchRef = useRef<() => Promise<void>>(null) // 重新抓取 performance 資料的

    // disabled
    const disabled = (value: performanceType): boolean => user.selfData.Permession === 2 ||
        (user.selfData.Permession === 1 &&
        value.UserId === user.selfData.UserId)

    // 績效的table 資料
    const performance = useMemo(() => {
        return company.performance?.map((item) => ({
            ...item,
            Goal: (
                item.Goal?.split('\n')?.map((i) =>
                    <p className='m-0' key={v4()}>{i}</p>
                )
            ),
            Directions: (
                item.Directions?.split('\n')?.map((i) =>
                    <p className='m-0' key={v4()}>{i}</p>
                )
            ),
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

    // modal 關閉
    const onClose = (): void => {
        setModal(modalStateInit)
    }

    // 當編輯完 儲存時
    const onEditSave = async (v: performanceType): Promise<void> => {
        let res = null
        if (!v.BanchId && user.selfData.Permession === 100) {
            FullMessage.error('請選擇部門')
            return
        }
        if (modal.type === '4') {
            res = await api.createPerformance(v)
        } else {
            res = await api.updatePerformance(v)
        }
        if (res?.status) {
            await reFetchRef.current()
            onClose()
        }
    }
    useEffect(() => {
        if (modal.type === '2') {
            Modal.confirm({
                okText: '確認',
                cancelText: '取消',
                title: '警告 是否刪除',
                content: <>
                        姓名： {modal.value.UserName}<br/>
                        年度： {modal.value.Year}<br/>
                        月份： {modal.value.Month}<br/>
                </>,
                onOk: async () => {
                    const res = await api.deletePerformance(modal.value.PerformanceId)
                    if (res.status) {
                        await reFetchRef.current()
                        onClose()
                    }
                },
                onCancel: onClose
            })
        } else if (modal.type === '5') {
            const toYear = modal.value.Month === 12 ? modal.value.Year + 1 : modal.value.Year
            const toMonth = modal.value.Month === 12 ? 1 : modal.value.Month + 1
            Modal.confirm({
                okText: '確認',
                cancelText: '取消',
                title: '警告 是否複製',
                content: <div className='row'>
                    <div className='col-md-4'>
                        姓名： {modal.value.UserName}<br/>
                        年度： {modal.value.Year}<br/>
                        月份： {modal.value.Month}<br/>
                    </div>
                    <ArrowRightOutlined className='col-md-2' />
                    <div className='col-md-4'>
                        姓名： {modal.value.UserName}<br/>
                        年度： {toYear}<br/>
                        月份： {toMonth}<br/>
                    </div>
                </div>,
                onOk: async () => {
                    const res = await api.copyPerformance(modal.value.PerformanceId)
                    if (res.status) {
                        await reFetchRef.current()
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
                        <EditModal
                            type={modal.type}
                            onSave={onEditSave}
                            onClose={onClose}
                            value={modal.value}
                            open={modal.open}
                            banchId={state.banchId}
                        />
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
                <SearchBar reFetchRef={reFetchRef} />
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
