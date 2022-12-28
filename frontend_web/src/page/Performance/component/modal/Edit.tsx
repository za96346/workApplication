import React, { useRef } from 'react'
import { Modal, Form, Input, InputNumber } from 'antd'
import { performanceType } from 'type'
import useReduceing from 'Hook/useReducing'
import DescribeValue from '../DescribeValue'

interface props {
    open: boolean
    value: performanceType | null
    onClose: () => void
    onSave: ((v: performanceType) => void)
}
const Edit = ({ open, value, onClose, onSave }: props): JSX.Element => {
    const form = useRef<performanceType>(value)
    const { user } = useReduceing()
    const disabled = user.selfData.Permession === 2 ||
        (user.selfData.Permession === 1 &&
        value.UserId === user.selfData.UserId)
    return (
        <Modal
            title="編輯績效"
            destroyOnClose
            open={open}
            onCancel={onClose}
            onOk={() => { onSave({ ...value, ...form.current }) }}
            okText="儲存"
            cancelText="取消"
        >
            <DescribeValue value={value}/>
            <Form onValuesChange={(v, allV) => { form.current = allV; console.log(allV) }} className="row mt-4">
                <Form.Item name="Goal" initialValue={value?.Goal || ''} className="" label="年度目標">
                    <Input.TextArea autoSize placeholder="輸入年度目標"/>
                </Form.Item>
                <Form.Item name="Attitude" initialValue={value?.Attitude || 0} className="col-4" label="態度">
                    <InputNumber disabled={disabled} min={0} max={10} placeholder="態度"/>
                </Form.Item>
                <Form.Item name="Efficiency" initialValue={value?.Efficiency || 0} className="col-4" label="效率">
                    <InputNumber disabled={disabled} min={0} max={10} placeholder="效率"/>
                </Form.Item>
                <Form.Item name="Professional" initialValue={value?.Professional || 0} className="col-4" label="專業">
                    <InputNumber disabled={disabled} min={0} max={10} placeholder="專業"/>
                </Form.Item>
                <Form.Item name="BeLate" initialValue={value?.BeLate || 0} className="col-6" label="遲到/早退">
                    <InputNumber disabled={disabled} min={0} max={100} placeholder="次數"/>
                </Form.Item>
                <Form.Item name="DayOffNotOnRule" initialValue={value?.DayOffNotOnRule || 0} className="col-6" label="未依規定請假">
                    <InputNumber disabled={disabled} min={0} max={100} placeholder="次數"/>
                </Form.Item>
                <Form.Item name="Directions" initialValue={value?.Directions || ''} className="" label="績效描述">
                    <Input.TextArea disabled={disabled} autoSize placeholder="績效描述"/>
                </Form.Item>
            </Form>
        </Modal>
    )
}
export default Edit
