import { Button, Form, Input, Modal, Table } from 'antd'
import rule from 'method/rule'
import React, { useEffect, useState, useRef } from 'react'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'

import { BanchType } from '../../type'
import { columns, dataSource } from './method/tableData'

const BanchManager: React.FC = () => {
    const { company } = useReduceing()
    const formRef = useRef('')
    const [modal, setModal] = useState<{ open: boolean, value: BanchType, type: number }>({
        open: false,
        value: null,
        type: 1 // 1 編輯 // 2 新增
    })
    const modalTitle = modal.type === 1 ? '編輯部門' : '新增部門'
    const onClose = async (): Promise<void> => {
        setModal((prev) => ({ ...prev, open: false, value: null }))
        await api.getBanch()
    }
    const onEdit = (v: BanchType): void => {
        setModal((prev) => ({ ...prev, value: v, open: true, type: 1 }))
    }
    const onCreate = (): void => {
        setModal((prev) => ({ ...prev, value: null, open: true, type: 2 }))
    }
    const onDelete = async (v: BanchType): Promise<void> => {
        Modal.confirm({
            okText: '確認',
            cancelText: '取消',
            title: "刪除部門",
            content: `使否要刪除 ${v.BanchName}`,
            onOk: async () => {
                await api.deleteBanch(v.Id)
                await onClose()
            }
        })
    }
    const onSave = async (): Promise<void> => {
        if (modal.type === 2) {
            await api.createBanch(formRef.current)
        } else {
            await api.UpdateBanch(formRef.current, modal.value.Id)
        }
        await onClose()
    }
    useEffect(() => {
        void api.getUserAll({
            name: '',
            workState: 'on'
        })
        void api.getBanch()
    }, [])
    return (
        <>
            <Modal
                open={modal.open}
                onOk={onSave}
                onCancel={onClose}
                cancelText="取消"
                okText="儲存"
                title={modalTitle}
                destroyOnClose
            >
                <Form onValuesChange={(v, allV) => { formRef.current = allV?.BanchName }}>
                    <Form.Item
                        rules={rule.banch()}
                        name="BanchName"
                        label={modalTitle}
                        initialValue={modal.value?.BanchName || ''}
                    >
                        <Input placeholder='請輸入部門名稱' />
                    </Form.Item>
                </Form>
            </Modal>
            <div className={window.styles.banchManagerBlock}>
                <Button onClick={onCreate} className='mb-3'>
                    新增部門
                </Button>
                <Table
                    dataSource={dataSource(
                        company.banch,
                        onEdit,
                        company.employee,
                        onDelete
                    )}
                    columns={columns}
                />
            </div>
        </>

    )
}

export default BanchManager
