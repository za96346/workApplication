import { DatePicker, Form, Input, Modal } from 'antd'
import React, { useRef } from 'react'
import StatusSelector from 'Share/StatusSelector'
import PermessionSelector from 'Share/PermessionSelector'
import BanchSelector from 'Share/BanchSelector'
import api from 'api/api'

interface props {
    open: boolean
    onClose: () => void
}
const ModalCreate = ({
    open,
    onClose
}: props): JSX.Element => {
    const formRef = useRef({})
    const onSave = async (): Promise<void> => {
        const res = await api.createUser({ ...formRef.current })
        if (res.status) {
            onClose()
        }
    }
    return (
        <Modal
            destroyOnClose
            open={open}
            onCancel={onClose}
            onOk={onSave}
            okText="儲存"
            cancelText="取消"
            title='新增員工資料'
        >
            <Form onValuesChange={(v, allV) => { formRef.current = allV }} className='row mt-5'>
                <Form.Item name="Account" label="帳號" className='col-md-6'>
                    <Input placeholder='請輸入帳號' />
                </Form.Item>
                <Form.Item name="Password" label="密碼" className='col-md-6'>
                    <Input.Password placeholder='密碼' />
                </Form.Item>
                <Form.Item name="UserName" label="姓名" className='col-md-6'>
                    <Input placeholder='請輸入姓名' />
                </Form.Item>
                <Form.Item name="EmployeeNumber" label="員工編號" className='col-md-6'>
                    <Input placeholder='請輸入員工編號' />
                </Form.Item>
                <Form.Item name="WorkState" label="狀態" initialValue={'on'} className='col-md-6'>
                    <StatusSelector />
                </Form.Item>
                <Form.Item name="Permession" initialValue={2} label="權限" className='col-md-6'>
                    <PermessionSelector />
                </Form.Item>
                <Form.Item name="OnWorkDay" label="到職日" className='col-md-6'>
                    <DatePicker />
                </Form.Item>
                <Form.Item name="Banch" label="部門" className='col-md-6'>
                    <BanchSelector />
                </Form.Item>
            </Form>
        </Modal>
    )
}

export default ModalCreate
