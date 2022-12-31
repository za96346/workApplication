import { DatePicker, Form, Input, Modal, Descriptions } from 'antd'
import React, { useEffect, useRef } from 'react'
import StatusSelector from 'Share/StatusSelector'
import { UserType } from 'type'
import PermessionSelector from 'Share/PermessionSelector'
import BanchSelector from 'Share/BanchSelector'
import api from 'api/api'
import dayjs from 'dayjs'

interface props {
    open: boolean
    onClose: () => void
    data: UserType
}
const ModalEdit = ({
    open,
    onClose,
    data
}: props): JSX.Element => {
    const formRef = useRef({})
    const onSave = async (): Promise<void> => {
        const res = await api.updateUser({ ...data, ...formRef.current })
        if (res.status) {
            onClose()
        }
    }
    useEffect(() => {
        formRef.current = data
    }, [data])
    return (
        <Modal
            destroyOnClose
            open={open}
            onCancel={onClose}
            onOk={onSave}
            okText="儲存"
            cancelText="取消"
            title='編輯員工資料'
        >
            <Descriptions>
                <Descriptions.Item label="姓名" span={3}>{data?.UserName || ''}</Descriptions.Item>
            </Descriptions>
            <Form onValuesChange={(v, allV) => { formRef.current = allV }} className='row mt-5'>
                <Form.Item name="EmployeeNumber" label="員工編號" initialValue={data?.EmployeeNumber || ''} className='col-md-6'>
                    <Input placeholder='請輸入員工編號' />
                </Form.Item>
                <Form.Item name="WorkState" label="狀態" initialValue={data?.WorkState || 'on'} className='col-md-6'>
                    <StatusSelector />
                </Form.Item>
                <Form.Item name="Permession" initialValue={data?.Permession || 2} label="權限" className='col-md-6'>
                    <PermessionSelector />
                </Form.Item>
                <Form.Item name="OnWorkDay" label="到職日" initialValue={dayjs(data?.OnWorkDay)} className='col-md-6'>
                    <DatePicker />
                </Form.Item>
                <Form.Item name="Banch" label="部門" initialValue={data?.Banch !== -1 && data?.Banch} className='col-md-6'>
                    <BanchSelector />
                </Form.Item>
            </Form>
        </Modal>
    )
}

export default ModalEdit
