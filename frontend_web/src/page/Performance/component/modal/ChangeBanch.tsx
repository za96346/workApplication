import { Descriptions, Form, Modal } from 'antd'
import React from 'react'
import BanchSelector from 'Share/BanchSelector'
import { performanceType } from 'Root/type'

interface props {
    open: boolean
    value?: performanceType | null
    onClose: () => void
}

const ChangeBanch = ({ open, value, onClose }: props): JSX.Element => {
    return (
        <Modal
            title="更換部門"
            destroyOnClose
            open={open}
            onCancel={onClose}
            onOk={onClose}
            okText="儲存"
            cancelText="取消"
        >
            <Descriptions>
                <Descriptions.Item span={1} label="姓名">
                    {value?.UserName || ''}
                </Descriptions.Item>
                <Descriptions.Item span={1} label="年度">
                    {value?.Year || ''}
                </Descriptions.Item>
                <Descriptions.Item span={1} label="月份">
                    {value?.Month || ''}
                </Descriptions.Item>
            </Descriptions>
            <Form>
                <Form.Item label="選擇組別">
                    <BanchSelector defaultValue={value.BanchId} />
                </Form.Item>
            </Form>

        </Modal>
    )
}
export default ChangeBanch
