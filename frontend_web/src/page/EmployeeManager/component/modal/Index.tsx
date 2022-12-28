import { DatePicker, Form, Input, Modal, Descriptions } from 'antd'
import React from 'react'
import StatusSelector from 'Share/StatusSelector'
import { UserType } from 'type'
import PermessionSelector from 'Share/PermessionSelector'

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
    console.log(data)
    return (
        <Modal
            destroyOnClose
            open={open}
            onCancel={onClose}
            onOk={onClose}
            okText="儲存"
            cancelText="取消"
        >
            <Descriptions>
                <Descriptions.Item label="姓名" span={3}>{data?.UserName || ''}</Descriptions.Item>
            </Descriptions>
            <Form className='row mt-5'>
                <Form.Item name="EmployeeNumber" label="員工編號" initialValue={data?.EmployeeNumber || ''} className='col-md-6'>
                    <Input />
                </Form.Item>
                <Form.Item name="WorkState" label="狀態" className='col-md-6'>
                    <StatusSelector defaultValue={data?.WorkState || 'on'}/>
                </Form.Item>
                <Form.Item name="Permession" label="權限" className='col-md-6'>
                    <PermessionSelector defaultValue={data?.Permession || 2} />
                </Form.Item>
                <Form.Item name="OnWorkDay" label="到職日" className='col-md-6'>
                    <DatePicker />
                </Form.Item>
            </Form>
        </Modal>
    )
}

export default ModalEdit
