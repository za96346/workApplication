import React from 'react'
import { Modal, Form, Input, InputNumber, Descriptions } from 'antd'
import { performanceType } from 'Root/type'

interface props {
    open: boolean
    value: performanceType | null
    onClose: () => void
}
const Edit = ({ open, value, onClose }: props): JSX.Element => {
    return (
        <Modal
            title="編輯績效"
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
            <Form className="row mt-4">
                <Form.Item className="" label="年度目標">
                    <Input.TextArea autoSize defaultValue={value?.Goal || ''} placeholder="輸入年度目標"/>
                </Form.Item>
                <Form.Item className="col-4" label="態度">
                    <InputNumber defaultValue={value?.Attitude || 0} min={0} max={10} placeholder="態度"/>
                </Form.Item>
                <Form.Item className="col-4" label="效率">
                    <InputNumber defaultValue={value?.Efficiency || 0} min={0} max={10} placeholder="效率"/>
                </Form.Item>
                <Form.Item className="col-4" label="專業">
                    <InputNumber defaultValue={value?.Professional || 0} min={0} max={10} placeholder="專業"/>
                </Form.Item>
                <Form.Item className="col-6" label="遲到/早退">
                    <InputNumber defaultValue={value?.BeLate || 0} min={0} max={100} placeholder="次數"/>
                </Form.Item>
                <Form.Item className="col-6" label="未依規定請假">
                    <InputNumber defaultValue={value?.DayOffNotOnRule || 0} min={0} max={100} placeholder="次數"/>
                </Form.Item>
                <Form.Item className="" label="績效描述">
                    <Input.TextArea autoSize defaultValue={value?.Directions || ''} placeholder="績效描述"/>
                </Form.Item>
            </Form>
        </Modal>
    )
}
export default Edit
